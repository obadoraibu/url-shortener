package apiserver

type Config struct {
	BindAddr string `yaml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
