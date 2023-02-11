package HttpServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer http.Server
}

func (s *Server) HTTPServer(port string, router *gin.Engine) (*http.Server, error) {
	s.httpServer = http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return &s.httpServer, s.httpServer.ListenAndServe()
}
