package statserver

import "github.com/sirupsen/logrus"

// StatServer ...
type StatServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *StatServer {
	return &StatServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *StatServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting Statistics Server")
	return nil
}

func (s *StatServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
