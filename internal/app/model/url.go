package model

type URL struct {
	Id        int    `json:"id"`
	LongURL   string `json:"long_url"`
	ShortURL  string `json:"short_url"`
	DeleteKey string `json:"delete_key"`
}
