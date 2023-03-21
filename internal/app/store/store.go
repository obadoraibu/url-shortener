package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	dbUrl         string
	db            *sql.DB
	urlRepository *URLRepository
}

func New(u string) *Store {
	return &Store{
		dbUrl: u,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.dbUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) URL() *URLRepository {
	if s.urlRepository != nil {
		return s.urlRepository
	}

	s.urlRepository = &URLRepository{
		store: s,
	}

	return s.urlRepository
}
