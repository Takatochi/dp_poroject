package handler

import (
	"html/template"
	"net/http"
	"project/pkg/model"
	"project/pkg/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
	store  store.Store
}

func NewHandler(store store.Store) *Handler {
	hadler := &Handler{
		router: gin.New(),
		store:  store,
	}
	return hadler
}
func (h *Handler) Routing() *gin.Engine {
	h.router.Static("/static", "./static/")
	h.router.SetFuncMap(template.FuncMap{
		"whole":   Whole,
		"decimal": Decimal,
	})
	u := &model.User{
		Email:             "agavor",
		EncryptedPassword: "dsfds",
	}

	h.store.User().Create(u)

	h.router.LoadHTMLGlob("templates/*.html")

	h.router.GET("/", h.index)
	return h.router
}

// func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.router.ServeHTTP(w, r)
// }

func (h *Handler) index(ctx *gin.Context) {

	ctx.Request.ParseForm()
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, "index", gin.H{
		"Post": 2,
		"Rget": get,
	})

}

// func (s *contact) NewRequest(post any, maineroot string, handlename string, router *gin.Engine) {
// 	s.post = post
// 	s.maineroot = maineroot

// 	router.GET(handlename, s.serveHTTP)
// }

// func (s *contact) serveHTTP(ctx *gin.Context) {

// 	ctx.Request.ParseForm()
// 	get := ctx.Request.Form
// 	ctx.HTML(http.StatusOK, s.maineroot, gin.H{
// 		"Post": s.post,
// 		"Rget": get,
// 	})

// }
