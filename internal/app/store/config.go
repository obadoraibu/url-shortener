package store

type Config struct {
	DatabaseUrl string `toml: database_url`
}

func NewConfig() *Config {
	return &Config{
		DatabaseUrl: "postgres://postgres:qwerty@localhost:5432/restapi_dev?sslmode=disable",
	}
}
