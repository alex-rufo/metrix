package main

import (
	"log"

	"github.com/alex-rufo/metrix/pkg/metrix"
)

func main() {
	config := metrix.Config{
		namespace: "with_config",
		defaultTags: metrix.Tags{
			"tag_1": "key_1",
			"tag_2": 2,
		},
		providers: []metrix.Provider{
			metrix.DataDog{address: "127.0.0.1:8125"},
			metrix.Stdout{},
		},
	}

	// Configuring metrix
	m, err := metrix.NewFromConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Shutdown()

	metrix.SetAsDefault(m)

	// Adding metrics
	metrix.Increment("increment", metrix.Tags{"tag_3": false})
	metrix.Counter("counter", 10, map[string]interface{}{"tag_4": 1.23})
	metrix.Gauge("gauge", 2.4)
	metrix.Histogram("histogram", 1.7)
}
