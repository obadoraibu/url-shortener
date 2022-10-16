package apiserver

import "obadoraibu/url-shortener/internal/app/storage"

type Config struct {
	BindAddr      string `yaml:"bind_addr"`
	LogLevel      string `yaml:"log_level"`
	StorageConfig *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8080",
		LogLevel:      "debug",
		StorageConfig: storage.NewConfig(),
	}
}
