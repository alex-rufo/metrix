package dummy

// Provider structure
type Provider struct{}

// New creates a new stdout provider. All metric will be printed to the standard output
func New() *Provider {
	return &Provider{}
}

// Counter tracks how many times something happened per second
func (p *Provider) Counter(name string, value int64, tags map[string]string) error {
	return nil
}

// Histogram tracks the statistical distribution of a set of values
func (p *Provider) Histogram(name string, value float64, tags map[string]string) error {
	return nil
}

// Gauge measures the value of a metric at a particular time
func (p *Provider) Gauge(name string, value float64, tags map[string]string) error {
	return nil
}
