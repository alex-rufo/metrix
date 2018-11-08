package datadog

// Config to create a DataDog provider
type Config struct {
	Rate  float64 `config:"rate"    default:"1"`
	Token string  `config:"token"`
}
