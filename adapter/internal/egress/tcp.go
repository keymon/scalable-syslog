package egress

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context"

	"code.cloudfoundry.org/go-loggregator/pulseemitter"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	v1 "code.cloudfoundry.org/scalable-syslog/internal/api/v1"
	"github.com/crewjam/rfc5424"
)

// DialFunc represents a method for creating a connection, either TCP or TLS.
type DialFunc func(addr string) (net.Conn, error)

// TCPWriter represents a syslog writer that connects over unencrypted TCP.
// This writer is not meant to be used from multiple goroutines. The same
// goroutine that calls `.Write()` should be the one that calls `.Close()`.
type TCPWriter struct {
	url       *url.URL
	appID     string
	hostname  string
	dialFunc  DialFunc
	ioTimeout time.Duration
	scheme    string
	conn      net.Conn
	closed    bool
	ctx       context.Context

	egressMetric *pulseemitter.CounterMetric
}

// NewTCPWriter creates a new TCP syslog writer.
func NewTCPWriter(
	ctx context.Context,
	binding *v1.Binding,
	dialTimeout time.Duration,
	ioTimeout time.Duration,
	skipCertVerify bool,
	egressMetric *pulseemitter.CounterMetric,
) (WriteCloser, error) {
	drainURL, err := url.Parse(binding.Drain)
	// TODO: remove parsing/error from here
	if err != nil {
		return nil, err
	}

	dialer := &net.Dialer{
		Timeout: dialTimeout,
	}
	df := func(addr string) (net.Conn, error) {
		return dialer.Dial("tcp", addr)
	}

	w := &TCPWriter{
		url:          drainURL,
		appID:        binding.AppId,
		hostname:     binding.Hostname,
		ioTimeout:    ioTimeout,
		dialFunc:     df,
		scheme:       "syslog",
		egressMetric: egressMetric,
		ctx:          ctx,
	}

	return w, nil
}

func (w *TCPWriter) connection() (net.Conn, error) {
	if w.conn == nil {
		return w.connect()
	}
	return w.conn, nil
}

func (w *TCPWriter) connect() (net.Conn, error) {
	for {
		if w.closed {
			return nil, errors.New("attempting connect after close")
		}

		conn, err := w.dialFunc(w.url.Host)
		if err != nil {
			if contextDone(w.ctx) {
				return nil, err
			}

			duration := time.Minute
			log.Printf("failed to connect to %s, retrying in %s: %s", w.url.Host, duration, err)
			time.Sleep(duration)
			continue
		}
		w.conn = conn

		log.Printf("created conn to syslog drain: %s", w.url.Host)

		return conn, nil
	}
}

// Close tears down any active connections to the drain and prevents reconnect.
func (w *TCPWriter) Close() error {
	w.closed = true
	return w.connClose()
}

func (w *TCPWriter) connClose() error {
	if w.conn != nil {
		err := w.conn.Close()
		w.conn = nil

		return err
	}

	return nil
}

// Write writes an envelope to the syslog drain connection.
func (w *TCPWriter) Write(env *loggregator_v2.Envelope) error {
	if env.GetLog() == nil {
		return nil
	}

	msg := rfc5424.Message{
		Priority:  generatePriority(env.GetLog().Type),
		Timestamp: time.Unix(0, env.GetTimestamp()).UTC(),
		Hostname:  w.hostname,
		AppName:   w.appID,
		ProcessID: generateProcessID(
			env.Tags["source_type"].GetText(),
			env.Tags["source_instance"].GetText(),
		),
		Message: appendNewline(removeNulls(env.GetLog().Payload)),
	}

	conn, err := w.connection()
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(w.ioTimeout))
	_, err = msg.WriteTo(conn)
	if err != nil {
		_ = w.connClose()

		return err
	}

	w.egressMetric.Increment(1)

	return nil
}

func removeNulls(msg []byte) []byte {
	return bytes.Replace(msg, []byte{0}, nil, -1)
}

func appendNewline(msg []byte) []byte {
	if !bytes.HasSuffix(msg, []byte("\n")) {
		msg = append(msg, byte('\n'))
	}
	return msg
}

func generatePriority(logType loggregator_v2.Log_Type) rfc5424.Priority {
	switch logType {
	case loggregator_v2.Log_OUT:
		return rfc5424.Info + rfc5424.User
	case loggregator_v2.Log_ERR:
		return rfc5424.Error + rfc5424.User
	default:
		return rfc5424.Priority(-1)
	}
}

func generateProcessID(sourceType, sourceInstance string) string {
	sourceType = strings.ToUpper(sourceType)
	if sourceInstance != "" {
		return fmt.Sprintf("[%s/%s]", sourceType, sourceInstance)
	}

	return fmt.Sprintf("[%s]", sourceType)
}
