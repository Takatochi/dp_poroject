package handler

import (
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql/analyzer"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/app/model"
	"project/pkg/MYSQLserver"
	"project/pkg/logger"
	"project/pkg/port"
	StoreBD "project/pkg/store"
	"strconv"
)

type Handler struct {
	router *gin.Engine
	store  StoreBD.Store
	stores StoreBD.ListenStore
}
type Index struct {
	Handler *Handler
}

func NewHandler(store StoreBD.Store) *Handler {
	storeBD := &StoreBD.Listen{Store: store}
	return &Handler{
		router: gin.New(),
		store:  store,
		stores: storeBD,
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
	//stores, err := h.Handler.stores.StoreBD().Server().Find()

	store, err := h.Handler.stores.StoreBD().Server().Find()
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

func (h *Index) StartVirtualServer(ctx *gin.Context) {
	go startVirtualSqlserver("localhost", "alpha", 3310)
	ctx.Status(http.StatusOK)
}
func startVirtualSqlserver(address, dbname string, port int32) {
	config := &sqle.Config{
		VersionPostfix:     "Version",
		IsReadOnly:         false,
		IsServerLocked:     false,
		IncludeRootAccount: false,
	}

	cfg := &server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
		Version:  "Version",
	}

	db := memory.NewDatabase(dbname)
	analyzer := analyzer.NewDefault(analyzer.NewDatabaseProvider(db, information_schema.NewInformationSchemaDatabase()))

	MYs := MYSQLserver.NewMySqliDefault(cfg, analyzer, config)

	if err := MYs.Run(); err != nil {
		logger.Error(err)
	}

}
