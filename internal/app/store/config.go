package store

type Config struct {
	DatabaseUrl string `toml: database_url`
}

func NewConfig() *Config {
	return &Config{
		DatabaseUrl: "postgres://postgres:postgrespw@localhost:55002/apiserver_dev?sslmode=disable",
	}
}
