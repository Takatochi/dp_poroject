package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/app/model"
	"project/pkg/logger"
	"project/pkg/port"
	"project/pkg/store"
	"strconv"
)

type Handler struct {
	router *gin.Engine
	store  store.Store
}
type Index struct {
	Handler *Handler
}

func NewHandler(store store.Store) *Handler {
	return &Handler{
		router: gin.New(),
		store:  store,
	}

}
func (h *Handler) Routing() *gin.Engine {
	return h.router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// Index curl  -X GET http://localhost:8088/
func (h *Index) Index(ctx *gin.Context) {

	ctx.Request.ParseForm()

	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, "index", gin.H{
		"Rget": get,
	})

}

// New curl -d "user=user1" -X POST http://localhost:8088/New
func (h *Index) New(ctx *gin.Context) {

	store, err := h.Handler.store.Server().Find()
	if err != nil {
		logger.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, store)

}

// Initiation curl  -X POST http://localhost:8088/Serve/init
func (h *Index) Initiation(ctx *gin.Context) {
	message := ctx.PostForm("message")
	ports, err := port.GetFreePort()
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusTooManyRequests, "No free ports available")
		return
	}
	srv := model.Server{
		Name: message,
		Port: int64(ports),
	}
	err = h.Handler.store.Server().AddServer(&srv)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, "Server not found")
		return
	}
	ctx.JSON(http.StatusOK, srv)
}

// DeleteSever curl  -X DELETE http://localhost:8088/Serve/delete/server/:id
func (h *Index) DeleteSever(ctx *gin.Context) {
	userId := ctx.Param("id")
	fmt.Println(userId)
	num, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err})
		logger.Error(err)
		return
	}
	err = h.Handler.store.Server().DeleteServerFromDB(int(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.Status(http.StatusOK)
}
