package handler

import (
	"net/http"
	"project/pkg/store"

	"github.com/gin-gonic/gin"
)

type Routined interface {
	Routing(post any, maineroot string, handlename string, router *gin.Engine)
}
type Handler struct {
	Index   index
	Contact contact
	router  *gin.Engine
	Store   *store.Store
}
type path struct {
	maineroot string
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
	router.GET(handlename, s.serveHTTP)
}

func (s index) serveHTTP(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
		"Post": s.post,
		"Rget": get,
	})

}

func (s *contact) Routing(post any, maineroot string, handlename string, router *gin.Engine) {
	s.post = post
	s.maineroot = maineroot

	router.GET(handlename, s.serveHTTP)
}

func (s contact) serveHTTP(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
		"Post": s.post,
		"Rget": get,
	})

}
