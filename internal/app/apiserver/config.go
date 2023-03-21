package apiserver

type Config struct {
	BinAddr     string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:     ":8080",
		LogLevel:    "debug",
		DatabaseUrl: "postgres://postgres:qwerty@localhost:5432/restapi_dev?sslmode=disable",
	}
}
