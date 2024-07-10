package statserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// StatServer ...
type StatServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *StatServer {
	return &StatServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *StatServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting Statistics Server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *StatServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *StatServer) configureRouter() {
	s.router.HandleFunc("/get-order-book", s.GetOrderBook())
	s.router.HandleFunc("/save-order-book", s.SaveOrderBook())
	s.router.HandleFunc("/get-order-history", s.GetOrderHistory())
	s.router.HandleFunc("/save-order", s.SaveOrder())
}

func (s *StatServer) GetOrderBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "get_order_book")
	}
}

func (s *StatServer) SaveOrderBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "save_order_book")
	}
}

func (s *StatServer) GetOrderHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "get_order_history")
	}
}

func (s *StatServer) SaveOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "save_order")
	}
}
