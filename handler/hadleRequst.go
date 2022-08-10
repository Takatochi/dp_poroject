package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routined interface {
	Routing(post any, maineroot string, handlename string, router *gin.Engine)
}
type Handler struct {
	Index   index
	Contact contact
	router  *gin.Engine
}
type path struct {
	maineroot  string
	handlename string
}
type index struct {
	post any
	path
}
type contact struct {
	post any
	path
}

func InitHandler(router *gin.Engine) *Handler {
	return &Handler{
		router: router,
	}

}

func (s *index) Routing(post any, maineroot string, handlename string, router *gin.Engine) {
	s.post = post
	s.maineroot = maineroot
	s.handlename = handlename
	router.GET(s.handlename, s.ServeHTTP)
}

func (s index) ServeHTTP(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
		"Post": s.post,
		"Rget": get,
	})

}
func (s *index) Whole(a int, b int) bool {
	if a%b == 0 {
		return true
	} else if a%b == 1 {
		return false
	}
	return true
}
func (s *contact) Routing(post any, maineroot string, handlename string, router *gin.Engine) {
	s.post = post
	s.maineroot = maineroot
	s.handlename = handlename
	router.GET(s.handlename, s.ServeHTTP)
}

func (s contact) ServeHTTP(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
		"Post": s.post,
		"Rget": get,
	})

}
