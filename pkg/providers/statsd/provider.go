package statsd

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
)

// Config to create a datadog provider
type Config struct {
	Address string  `config:"address" default:"127.0.0.1:8125"`
	Rate    float64 `config:"rate"    default:"1"`
}

// Provider contains the datadog client and its configuration
type Provider struct {
	client *statsd.Client
	rate   float64
}

// New creates a new datadog provider
func New(addr string) (*Provider, error) {
	client, err := statsd.New(addr)
	if err != nil {
		return nil, err
	}

	p := Provider{
		client: client,
		rate:   1,
	}
	return &p, nil
}

// NewFromConfig creates a new datadog provider from a config
func NewFromConfig(c Config) (*Provider, error) {
	p, err := New(c.Address)
	if err != nil {
		return nil, err
	}

	p.WithRate(c.Rate)
	return p, nil
}

// WithRate updates the rate of datadog
func (p *Provider) WithRate(rate float64) {
	p.rate = rate
}

// Counter tracks how many times something happened per second
func (p *Provider) Counter(name string, value int64, tags map[string]string) error {
	return p.client.Count(name, value, p.formatTags(tags), p.rate)
}

// Histogram tracks the statistical distribution of a set of values
func (p *Provider) Histogram(name string, value float64, tags map[string]string) error {
	return p.client.Histogram(name, value, p.formatTags(tags), p.rate)
}

// Gauge measures the value of a metric at a particular time
func (p *Provider) Gauge(name string, value float64, tags map[string]string) error {
	return p.client.Gauge(name, value, p.formatTags(tags), p.rate)
}

// FormatTags from map[string]string to datadog format []string
func (p *Provider) formatTags(tags map[string]string) []string {
	formattedTags := []string{}
	for name, value := range tags {
		formattedTags = append(formattedTags, fmt.Sprintf("%s:%s", name, value))
	}

	return formattedTags
}
