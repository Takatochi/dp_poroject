package intternal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/pkg/HttpServer"
)

type InternalServer struct {
	Id     int
	router *gin.Engine
	server HttpServer.Server
}

func NewInternalServer(id int) *InternalServer {
	return &InternalServer{
		Id:     id,
		router: gin.New(),
	}

}
func (s *InternalServer) Run(port string) {

	info := s.router.Routes()
	log.Print(info)
	s.server.HTTPServer(port, s.router02())

}
func (s *InternalServer) router02() *gin.Engine {

	s.router.Use(gin.Recovery())

	s.router.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return s.router
}
