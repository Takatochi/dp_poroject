package intternal

import (
	"github.com/gin-gonic/gin"
	"project/pkg/HttpServer"
	"project/pkg/logger"
)

type InternalServer struct {
	Id     int
	router *gin.Engine
	server HttpServer.Server
	port   string
}

func NewInternalServer(id int, port string) *InternalServer {
	return &InternalServer{
		Id:     id,
		router: gin.New(),
		port:   port,
	}

}
func (s *InternalServer) Run(group, dir string, handlers gin.HandlerFunc) error {

	srv := new(HttpServer.Server)
	logger.Info(s.port)
	_, err := srv.HTTPServer(s.port, s.RouterVp(group, dir, handlers))
	if err != nil {
		return err
	}

	return nil
}
func (s *InternalServer) RouterVp(group, dir string, handlers gin.HandlerFunc) *gin.Engine {

	s.router.Use(gin.Recovery())

	s.router.Any(dir, handlers)
	//info := s.router.Routes()
	//log.Print(info)
	return s.router
}
