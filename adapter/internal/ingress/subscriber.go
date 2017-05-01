package ingress

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/cloudfoundry-incubator/scalable-syslog/adapter/internal/egress"
	v2 "github.com/cloudfoundry-incubator/scalable-syslog/internal/api/loggregator/v2"
	v1 "github.com/cloudfoundry-incubator/scalable-syslog/internal/api/v1"
	"github.com/cloudfoundry-incubator/scalable-syslog/internal/metric"
)

type MetricEmitter interface {
	IncCounter(name string, options ...metric.IncrementOpt)
}

type ClientPool interface {
	Next() (client v2.EgressClient)
}

type SyslogConnector interface {
	Connect(binding *v1.Binding) (cw egress.WriteCloser, err error)
}

// Subscriber streams loggregator egress to the syslog drain.
type Subscriber struct {
	pool      ClientPool
	connector SyslogConnector
	emitter   MetricEmitter
}

// NewSubscriber returns a new Subscriber.
func NewSubscriber(p ClientPool, c SyslogConnector, e MetricEmitter) *Subscriber {
	return &Subscriber{
		pool:      p,
		connector: c,
		emitter:   e,
	}
}

// Start begins to stream logs from a loggregator egress client to the syslog
// egress writer. Start does not block. Start returns a function that can be
// called to stop streaming logs.
func (s *Subscriber) Start(binding *v1.Binding) func() {
	var unsubscribe int32

	go s.connectAndRead(binding, &unsubscribe)

	return func() {
		atomic.AddInt32(&unsubscribe, 1)
	}
}

func (s *Subscriber) connectAndRead(binding *v1.Binding, unsubscribe *int32) {
	for {
		cont := s.attemptConnectAndRead(binding, unsubscribe)
		if !cont {
			return
		}
	}
}

func (s *Subscriber) attemptConnectAndRead(binding *v1.Binding, unsubscribe *int32) bool {
	writer, err := s.connector.Connect(binding)
	if err != nil {
		log.Println("Failed connecting to syslog: %s", err)
		// If connect fails it is likely due to a parse error with the binding
		// URL or other input error. As a result we should not retry
		// connecting.
		return false
	}
	defer writer.Close()

	// TODO: What if we cannot get a client?
	client := s.pool.Next()
	ctx, cancel := context.WithCancel(context.Background())
	receiver, err := client.Receiver(ctx, &v2.EgressRequest{
		ShardId: buildShardId(binding),
		Filter: &v2.Filter{
			SourceId: binding.AppId,
			Message: &v2.Filter_Log{
				Log: &v2.LogFilter{},
			},
		},
	})
	if err != nil {
		return true
	}
	defer cancel()
	defer receiver.CloseSend()

	if err := s.readWriteLoop(receiver, writer, unsubscribe); err != nil {
		log.Printf("Subscriber read/write loop has unexpectedly closed: %s", err)
		return true
	}

	return false
}

func (s *Subscriber) readWriteLoop(r v2.Egress_ReceiverClient, w egress.WriteCloser, unsubscribe *int32) error {
	var count uint64
	lastEmitted := time.Now()

	for {
		if atomic.LoadInt32(unsubscribe) > 0 {
			return nil
		}

		env, err := r.Recv()
		if err != nil {
			return err
		}
		if env.GetLog() == nil {
			continue
		}

		count++
		if count >= 1000 || time.Since(lastEmitted) > 5*time.Second {
			// metric-documentation-v2: (scalablesyslog.adapter.ingress) The number
			// of received log messages.
			s.emitter.IncCounter("ingress",
				metric.WithIncrement(count),
				metric.WithVersion(2, 0),
			)

			lastEmitted = time.Now()
			log.Printf("Ingressed %d envelopes", count)
			count = 0
		}

		// We decided to ignore the error from the writer since in most
		// situations the connector will provide a diode writer and the diode
		// writer never returns an error.
		_ = w.Write(env)
	}
}

func buildShardId(binding *v1.Binding) (key string) {
	return binding.AppId + binding.Hostname + binding.Drain
}
