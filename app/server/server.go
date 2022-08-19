package server

import (
	"context"
	"html/template"
	"net/http"
	"project/pkg/handler"
	"time"
)

type Server struct {
	httpServer *http.Server
	config     *Config
}

func (s *Server) Run(config *Config, had *handler.Handler) error {

	// init handler
	index := handler.Index{Handler: had}
	//register routing
	configureRouter(had, &index)
	//get router
	router := had.Routing()

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

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
func configureRouter(h *handler.Handler, i *handler.Index) {
	router := h.Routing()
	router.Static("/static", "./static/")
	router.SetFuncMap(template.FuncMap{
		"whole":   handler.Whole,
		"decimal": handler.Decimal,
	})

	router.LoadHTMLGlob("templates/*.html")
	api := router.Group("/")
	{
		api.GET("/", i.Index)
		api.GET("/api", i.Api)
	}

}
