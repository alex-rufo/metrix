package metrix

import (
	"fmt"
)

// Provider is the interface any provider needs to implement to be able to work with metrix
type Provider interface {
	Counter(name string, value int64, tags map[string]string) error
	Histogram(name string, value float64, tags map[string]string) error
	Gauge(name string, value float64, tags map[string]string) error
}

type Metrix struct {
	namespace   string
	defaultTags Tags
	providers   []Provider
}

func New() *Metrix {
	return &Metrix{defaultTags: Tags{}}
}

// WithNamespace sets the namespace to be used
func (m *Metrix) WithNamespace(namespace string) {
	m.namespace = namespace
}

// AddProvider to send metrixs to
func (m *Metrix) AddProvider(provider Provider) {
	m.providers = append(m.providers, provider)
}

// AddDefaultTags add some tags that will be present in all metrics.
// A default tag can be overrided for a specific call if it has the same name.
func (m *Metrix) AddDefaultTags(tags Tags) {
	m.defaultTags.Add(tags)
}

// Counter tracks how many times something happened per second
func (m *Metrix) Counter(name string, value int64, tags ...Tags) {
	formattedTags := m.formatTags(tags)
	formattedName := m.formatName(name)

	for _, provider := range m.providers {
		go func(p Provider) {
			p.Counter(formattedName, value, formattedTags)
		}(provider)
	}
}

// Increment is just a Counter of +1
func (m *Metrix) Increment(name string, tags ...Tags) {
	m.Counter(name, 1, tags...)
}

// Histogram tracks the statistical distribution of a set of values
func (m *Metrix) Histogram(name string, value float64, tags ...Tags) {
	formattedTags := m.formatTags(tags)
	formattedName := m.formatName(name)

	for _, provider := range m.providers {
		go func(p Provider) {
			p.Histogram(formattedName, value, formattedTags)
		}(provider)
	}
}

// Gauge measures the value of a metric at a particular time
func (m *Metrix) Gauge(name string, value float64, tags ...Tags) {
	formattedTags := m.formatTags(tags)
	formattedName := m.formatName(name)

	for _, provider := range m.providers {
		go func(p Provider) {
			p.Gauge(formattedName, value, formattedTags)
		}(provider)
	}
}

func (m *Metrix) formatName(name string) string {
	if m.namespace == "" {
		return name
	}

	return fmt.Sprintf("%s.%s", m.namespace, name)
}

func (m *Metrix) formatTags(tags []Tags) map[string]string {
	allTags := Tags{}
	allTags.Add(m.defaultTags)
	for _, t := range tags {
		allTags.Add(t)
	}

	return allTags.toStringMap()
}
