package store

import (
	"fmt"
	"github.com/obadoraibu/url-shortener/internal/app/model"
	"github.com/obadoraibu/url-shortener/internal/app/random"
)

type URLRepository struct {
	store *Store
}

func (r *URLRepository) Create(u *model.URL) (*model.URL, error) {
	if u.ShortURL == "" {
		exists := true
		var result string
		for exists {
			result = random.GenerateRandomString(5)
			err := r.store.db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM urls WHERE short_url = '%s')", result)).Scan(&exists)
			if err != nil {
				return nil, err
			}
		}
		u.ShortURL = result
	}
	fmt.Println(u.LongURL)
	fmt.Println(u.ShortURL)
	if err := r.store.db.QueryRow("INSERT INTO urls (long_url, short_url, delete_key) VALUES ($1, $2, $3) ON CONFLICT (short_url) DO NOTHING RETURNING id",
		u.LongURL, u.ShortURL, u.DeleteKey).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *URLRepository) FindByShortUrl(shortUrl string) (*model.URL, error) {
	u := &model.URL{ShortURL: shortUrl}
	if err := r.store.db.QueryRow("SELECT * FROM urls WHERE short_url = $1",
		shortUrl).Scan(&u.Id, &u.LongURL, &u.ShortURL, &u.DeleteKey); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *URLRepository) Delete(shortUrl, deleteKey string) error {
	if _, err := r.store.db.Exec(fmt.Sprintf("DELETE FROM urls WHERE short_url = '%s' AND delete_key = '%s'", shortUrl, deleteKey)); err != nil {
		return err
	}
	return nil
}
