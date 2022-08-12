package server

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	DriverName  string `toml:"driverName"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
	}
}
