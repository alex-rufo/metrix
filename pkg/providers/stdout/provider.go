package stdout

import (
	"fmt"
)

// Config to create a stdout provider
type Config struct{}

// Provider structure
type Provider struct{}

// New creates a new stdout provider. All metric will be printed to the standard output
func New() *Provider {
	return &Provider{}
}

// NewFromConfig creates a new stadout provider from a config
func NewFromConfig(c Config) *Provider {
	return New()
}

// Counter tracks how many times something happened per second
func (p *Provider) Counter(name string, value int64, tags map[string]string) error {
	fmt.Printf("[COUNTER]: name '%s' with value %d and tags %+v\n", name, value, tags)
	return nil
}

// Histogram tracks the statistical distribution of a set of values
func (p *Provider) Histogram(name string, value float64, tags map[string]string) error {
	fmt.Printf("[HISTOGRAM]: name '%s' with value %f and tags %+v\n", name, value, tags)
	return nil
}

// Gauge measures the value of a metric at a particular time
func (p *Provider) Gauge(name string, value float64, tags map[string]string) error {
	fmt.Printf("[GAUGE]: name '%s' with value %f and tags %+v\n", name, value, tags)
	return nil
}
