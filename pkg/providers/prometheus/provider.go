package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Config to create a prometheus provider
type Config struct {
	Address string  `config:"address" default:"127.0.0.1:8125"`
	Rate    float64 `config:"rate"    default:"1"`
}

type Provider struct {
}

// New creates a new prometheus provider
func New() *Provider {
	return &Provider{}
}

// NewFromConfig creates a new prometheus provider from a config
func NewFromConfig(c Config) (*Provider, error) {
	p, err := New(c.Address)
	if err != nil {
		return nil, err
	}

	p.WithRate(c.Rate)
	return p, nil
}

// Counter tracks how many times something happened per second
func (p *Provider) Counter(name string, value int64, tags map[string]string) error {
	opts := prometheus.CounterOpts{
		Name:        name,
		ConstLabels: tags,
	}
	prometheus.NewCounter(opts).Add(float64(value))
	return nil
}

// Histogram tracks the statistical distribution of a set of values
func (p *Provider) Histogram(name string, value float64, tags map[string]string) error {
	opts := prometheus.HistogramOpts{
		Name:        name,
		ConstLabels: tags,
	}
	prometheus.NewHistogram(opts).Observe(value)
	return nil
}

// Gauge measures the value of a metric at a particular time
func (p *Provider) Gauge(name string, value float64, tags map[string]string) error {
	opts := prometheus.GaugeOpts{
		Name:        name,
		ConstLabels: tags,
	}
	prometheus.NewGauge(opts).Add(value)
	return nil
}
