package metricemitter

import (
	"fmt"
	"sync/atomic"
	"time"

	v2 "code.cloudfoundry.org/scalable-syslog/internal/api/loggregator/v2"
)

type CounterMetric struct {
	client   *client
	name     string
	sourceID string
	tags     map[string]*v2.Value
	delta    uint64
}

type MetricOption func(taggedMetric)

func NewCounterMetric(name, sourceID string, opts ...MetricOption) *CounterMetric {
	m := &CounterMetric{
		name:     name,
		sourceID: sourceID,
		tags:     make(map[string]*v2.Value),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *CounterMetric) Increment(c uint64) {
	atomic.AddUint64(&m.delta, c)
}

func (m *CounterMetric) GetDelta() uint64 {
	return atomic.LoadUint64(&m.delta)
}

func (m *CounterMetric) SendWith(fn func(*v2.Envelope) error) error {
	d := atomic.SwapUint64(&m.delta, 0)

	if err := fn(m.toEnvelope(d)); err != nil {
		atomic.AddUint64(&m.delta, d)
		return err
	}

	return nil
}

func (m *CounterMetric) toEnvelope(delta uint64) *v2.Envelope {
	return &v2.Envelope{
		SourceId:  m.sourceID,
		Timestamp: time.Now().UnixNano(),
		Message: &v2.Envelope_Counter{
			Counter: &v2.Counter{
				Name: m.name,
				Value: &v2.Counter_Delta{
					Delta: delta,
				},
			},
		},
		Tags: m.tags,
	}
}

func (m *CounterMetric) setTag(k, v string) {
	m.tags[k] = &v2.Value{
		Data: &v2.Value_Text{
			Text: v,
		},
	}
}

func WithVersion(major, minor uint) MetricOption {
	return WithTags(map[string]string{
		"metric_version": fmt.Sprintf("%d.%d", major, minor),
	})
}

type taggedMetric interface {
	setTag(key, value string)
}

func WithTags(tags map[string]string) MetricOption {
	return func(t taggedMetric) {
		for k, v := range tags {
			t.setTag(k, v)
		}
	}
}