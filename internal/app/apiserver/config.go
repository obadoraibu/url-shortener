package apiserver

import "github.com/obadoraibu/url-shortener/internal/app/store"

type Config struct {
	BinAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	StoreConfig *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		StoreConfig: store.NewConfig(),
	}
}
