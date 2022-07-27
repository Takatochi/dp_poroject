package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	handlename string
	maineroot  string
	post       any
	router     *gin.Engine
}

func InitHandler(router *gin.Engine) *Handler {
	return &Handler{
		router: router,
	}

}

func (s *Handler) RequestTemplate(post any, maineroot string, handlename string) {

	s.post = post
	s.handlename = handlename
	s.maineroot = maineroot

	s.router.GET(s.handlename, s.ServeHTTP)

}

func (s *Handler) ServeHTTP(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form

	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
		"Post": s.post,
		"Rget": get["id"][0],
	})

}
