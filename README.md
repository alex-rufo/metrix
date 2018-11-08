# Metrix
This library provides a `metrics` package which can be used to instrument code, expose application metrics, and profile runtime performance in a flexible manner.

## How to use it
There are different ways to use metrix. You can use it like a SDK, an agent (HTTP or grpc) or a Lambda function.

### SDK
```
go get github.com\alex-rufo\metrix
```

```
m := monitor.New()
m.AddProvider(datadog.New("127.0.0.1:500"))
```

### Agent

### Lambda

## Providers
The following providers are already built in for you:
 - Prometheus
 - DataDog
 - RedShift
 - Elasticsearch
 - StdOut

### Add more provider
To add a new provider you just need to implement this interface:
```
type Provider interface {
	Counter(name string, value int64, tags map[string]string) error
	Histogram(name string, value float64, tags map[string]string) error
	Gauge(name string, value float64, tags map[string]string) error
}
```

Once you have it you just need to register it to the monitor.
```
yourProvider := ...

m := monitor.New()
m.AddProvider(yourProvider)
```
