package apiserver

import "github.com/sirupsen/logrus"

// APIserver structure
type APIserver struct {
	config *Config
	logger *logrus.Logger
}

// New server object
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
	}
}

// Start server
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting the server on ", s.config.BindAddr)
	return nil
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
