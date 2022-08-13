package server

import (
	"database/sql"
	"html/template"
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

	// init bd
	store := sqlBd.New(db)
	// init handler
	srv := handler.NewHandler(store)
	//register routing
	configureRouter(srv)
	//get router
	router := srv.Routing()

	// init server add addr and router
	s.httpServer = &http.Server{
		Addr:           ":" + config.BindAddr,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()

}

// init db open
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

func configureRouter(h *handler.Handler) {

	h.Routing().Static("/static", "./static/")
	h.Routing().SetFuncMap(template.FuncMap{
		"whole":   handler.Whole,
		"decimal": handler.Decimal,
	})

	h.Routing().LoadHTMLGlob("templates/*.html")

	h.Routing().GET("/", h.Index.ServeHTTP)

}
