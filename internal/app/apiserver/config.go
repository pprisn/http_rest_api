package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

//New ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
