package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/obadoraibu/url-shortener/internal/app/model"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/mux"
	"github.com/obadoraibu/url-shortener/internal/app/store"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting apiserver")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.HandleMain())
	s.router.HandleFunc("/add", s.HandleUrlAdd()).Methods("POST")
	s.router.HandleFunc("/{short_url}", s.HandleShort()).Methods("GET")
	s.router.HandleFunc("/delete", s.HandleDelete()).Methods("POST")
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *APIServer) Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func (s *APIServer) HandleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, 200, "URL-Shortener")
	}
}

func (s *APIServer) HandleUrlAdd() http.HandlerFunc {
	type request struct {
		LongUrl  string `json:"long_url"`
		ShortUrl string `json:"short_url"`
	}
	type respond struct {
		DeleteKey string `json:"delete_key"`
		ShortUrl  string `json:"short_url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		deleteKey := uuid.NewV4()
		u := &model.URL{
			LongURL:   req.LongUrl,
			ShortURL:  req.ShortUrl,
			DeleteKey: deleteKey.String(),
		}
		u, err := s.store.URL().Create(u)
		if err != nil {
			s.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.respond(w, r, 200, respond{DeleteKey: deleteKey.String(), ShortUrl: u.ShortURL})
	}
}

func (s *APIServer) HandleShort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortUrl := vars["short_url"]
		u, err := s.store.URL().FindByShortUrl(shortUrl)
		if err != nil {
			s.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Redirect(w, r, u.LongURL, http.StatusSeeOther)
	}
}

func (s *APIServer) HandleDelete() http.HandlerFunc {
	type request struct {
		ShortUrl  string `json:"short_url"`
		DeleteKey string `json:"delete_key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.store.URL().Delete(req.ShortUrl, req.DeleteKey); err != nil {
			s.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.respond(w, r, 200, "")
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.DatabaseUrl)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
