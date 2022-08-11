package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	Config     *Config
}

func (s *Server) Run(config *Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + config.BindAddr,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()

}
