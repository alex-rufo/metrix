package metrix

var defaultMetrix *Metrix

// SetAsDefault the given metrix instance
func SetAsDefault(m *Metrix) {
	defaultMetrix = m
}

// Counter tracks how many times something happened per second
func Counter(name string, value int64, tags ...Tags) {
	if defaultMetrix != nil {
		defaultMetrix.Counter(name, value, tags...)
	}
}

// Increment is just a Counter of +1
func Increment(name string, tags ...Tags) {
	if defaultMetrix != nil {
		defaultMetrix.Increment(name, tags...)
	}
}

// Histogram tracks the statistical distribution of a set of values
func Histogram(name string, value float64, tags ...Tags) {
	if defaultMetrix != nil {
		defaultMetrix.Histogram(name, value, tags...)
	}
}

// Gauge measures the value of a metric at a particular time
func Gauge(name string, value float64, tags ...Tags) {
	if defaultMetrix != nil {
		defaultMetrix.Gauge(name, value, tags...)
	}
}
