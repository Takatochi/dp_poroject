package server

import (
	"database/sql"
	"net/http"
	"project/pkg/handler"
	"project/pkg/store/sqlBd"
	"time"
)

type Server struct {
	httpServer *http.Server
	config     *Config
}

func (s *Server) Run(config *Config) error {

	s.config = config
	db, err := s.newDB()
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlBd.New(db)
	srv := handler.NewHandler(store)
	s.httpServer = &http.Server{
		Addr:           ":" + config.BindAddr,
		MaxHeaderBytes: 1 << 20,
		Handler:        srv.Routing(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()

}
func (s *Server) newDB() (*sql.DB, error) {
	db, err := sql.Open(s.config.DriverName, s.config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	return db, nil
}
