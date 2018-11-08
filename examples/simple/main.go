package main

import (
	"log"

	"github.com/alex-rufo/metrix/pkg/metrix"
	"github.com/alex-rufo/metrix/pkg/providers/statsd"
	"github.com/alex-rufo/metrix/pkg/providers/stdout"
)

func main() {
	// Configuring metrix
	m := metrix.New()
	m.WithNamespace("simple")
	m.AddDefaultTags(metrix.Tags{
		"tag_1": "key_1",
		"tag_2": 2,
	})

	// Adding providers
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	m.AddProvider(statsd)
	m.AddProvider(stdout.New())

	metrix.SetAsDefault(m)

	// Adding metrics
	metrix.Increment("increment", metrix.Tags{"tag_3": false})
	metrix.Counter("counter", 10, map[string]interface{}{"tag_4": 1.23})
	metrix.Gauge("gauge", 2.4)
	metrix.Histogram("histogram", 1.7)
}
