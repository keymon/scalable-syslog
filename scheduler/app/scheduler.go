package app

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/cloudfoundry-incubator/scalable-syslog/scheduler/internal/egress"
	"github.com/cloudfoundry-incubator/scalable-syslog/scheduler/internal/health"
	"github.com/cloudfoundry-incubator/scalable-syslog/scheduler/internal/ingress"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Scheduler represents the scheduler component. It is responsible for polling
// and/or streaming events from the cloud controller about syslog drains and
// updating a pool of adapters to service those drains.
type Scheduler struct {
	apiURL           string
	adapterAddrs     []string
	adapterTLSConfig *tls.Config

	healthAddr string
	client     *http.Client
	interval   time.Duration

	adapterService *egress.DefaultAdapterService
	fetcher        *ingress.BlacklistFilter

	blacklist *ingress.IPRanges
}

// NewScheduler returns a new unstarted scheduler.
func NewScheduler(
	apiURL string,
	adapterAddrs []string,
	adapterTLSConfig *tls.Config,
	opts ...SchedulerOption,
) *Scheduler {
	s := &Scheduler{
		apiURL:           apiURL,
		adapterAddrs:     adapterAddrs,
		adapterTLSConfig: adapterTLSConfig,
		healthAddr:       ":8080",
		client:           http.DefaultClient,
		interval:         15 * time.Second,
		blacklist:        &ingress.IPRanges{},
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

// SchedulerOption represents a function that can configure a scheduler.
type SchedulerOption func(c *Scheduler)

// WithHealthAddr sets the address for the health endpoint to bind to.
func WithHealthAddr(addr string) func(*Scheduler) {
	return func(s *Scheduler) {
		s.healthAddr = addr
	}
}

// WithHTTPClient sets the http.Client to poll the syslog drain binding provider.
func WithHTTPClient(client *http.Client) func(*Scheduler) {
	return func(s *Scheduler) {
		s.client = client
	}
}

// WithPollingInterval sets the interval to poll the syslog drain binding provider.
func WithPollingInterval(interval time.Duration) func(*Scheduler) {
	return func(s *Scheduler) {
		s.interval = interval
	}
}

// WithBlacklist sets the blacklist for the syslog IPs.
func WithBlacklist(r *ingress.IPRanges) func(*Scheduler) {
	return func(s *Scheduler) {
		s.blacklist = r
	}
}

// Start starts polling the syslog drain binding provider and serves the HTTP
// health endpoint.
func (s *Scheduler) Start() string {
	s.setupIngress()
	s.startEgress()
	return s.serveHealth()
}

func (s *Scheduler) setupIngress() {
	s.fetcher = ingress.NewBlacklistFilter(
		s.blacklist,
		ingress.NewVersionFilter(
			ingress.NewBindingFetcher(
				ingress.APIClient{
					Client: s.client,
					Addr:   s.apiURL,
				},
			),
		),
	)
}

func (s *Scheduler) startEgress() {
	creds := credentials.NewTLS(s.adapterTLSConfig)

	pool := egress.NewAdapterPool(s.adapterAddrs, grpc.WithTransportCredentials(creds))
	s.adapterService = egress.NewAdapterService(pool)
	orchestrator := egress.NewOrchestrator(s.fetcher, s.adapterService)
	go orchestrator.Run(s.interval)
}

func (s *Scheduler) serveHealth() string {
	router := http.NewServeMux()
	router.Handle("/health", health.NewHealth(s.fetcher, s.adapterService))

	server := http.Server{
		Addr:         s.healthAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server.Handler = router

	lis, err := net.Listen("tcp", s.healthAddr)
	if err != nil {
		log.Fatalf("Unable to setup Health endpoint (%s): %s", s.healthAddr, err)
	}

	go func() {
		log.Printf("Health endpoint is listening on %s", lis.Addr().String())
		log.Fatalf("Health server closing: %s", server.Serve(lis))
	}()
	return lis.Addr().String()
}
