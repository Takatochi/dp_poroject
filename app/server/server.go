package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"project/pkg/HttpServer"
	"project/pkg/handler"
	"time"
)

type Server struct {
	httpServer HttpServer.Server
	httpNet    http.Server
}

func (s *Server) HTTPServer(port string, router *gin.Engine) error {
	httpServer, err := s.httpServer.HTTPServer(port, router)

	s.httpNet = httpServer
	return err
}
func (s *Server) Run(config *Config, had *handler.Handler) error {

	// init handler
	index := handler.Index{Handler: had}
	//register routing
	configureRouter(had, &index)
	//get router
	router := had.Routing()

	// init server add addr and router
	err := s.HTTPServer(config.BindAddr, router)

	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpNet.Shutdown(ctx)
}
func configureRouter(h *handler.Handler, i *handler.Index) {
	router := h.Routing()
	router.Static("/static", "./static/")
	router.SetFuncMap(template.FuncMap{
		"whole":   handler.Whole,
		"decimal": handler.Decimal,
	})
	ss := gin.Logger()
	println(ss)
	router.Use(gin.LoggerWithFormatter(LogerDeager))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.LoadHTMLGlob("templates/*.html")

	data := router.Group("/")
	{
		data.GET("/", i.Index)
		data.POST("/New", i.New)
	}
	server := router.Group("/Server")
	{
		server.POST("/init", i.Initiation)
	}

}
func LogerDeager(param gin.LogFormatterParams) string {

	// your custom format
	log.Print(param.Method)
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
