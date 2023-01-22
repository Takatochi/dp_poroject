package handler

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/app/model"
	"project/pkg/handler/mapHadler"
	"project/pkg/logger"
	"project/pkg/port"
	StoreBD "project/pkg/store"
	"project/pkg/stringFMT"
	"strconv"
)

var ServerTree *redblacktree.Tree

type serverCFG struct {
	Port int32  `json:"port"`
	Name string `json:"name"`
}

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

	var srvCFG = new(serverCFG)

	if err := ctx.ShouldBindJSON(&srvCFG); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(srvCFG)

	ctx.JSON(http.StatusOK, gin.H{"message": "Server started"})

	errs := port.GetQuestionFreePort("localhost", srvCFG.Port)
	if errs != nil {
		logger.Error(errs)
		ctx.JSON(http.StatusTooManyRequests, gin.H{"message": fmt.Sprintf("Server with port %d already use", srvCFG.Port)})
		return
	}

	//ListServerSql := make(chan []mapHadler.ListServerSql, 1)
	serverTree := make(chan *redblacktree.Tree, 1)
	go func() {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Server with port %d start", srvCFG.Port)})
		logger.Infof("Server with port %d already use", srvCFG.Port)
		serverTree <- mapHadler.NewServerSql("localhost", stringFMT.StringTitleJoin(srvCFG.Name), srvCFG.Port)

		close(serverTree)
	}()

	ServerTree = <-serverTree

}

func (h *Index) CloseVirtualServer(ctx *gin.Context) {
	portGet := ctx.Param("port")
	ports, err := strconv.Atoi(portGet)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err})
		logger.Error(err)
		return
	}
	if server, found := ServerTree.Get(ports); found {
		fmt.Printf("server type: %T ", server)
		if server != nil {
			fmt.Printf("Server found with port %d ", server.(mapHadler.ListServerSql).Port)
			err := server.(mapHadler.ListServerSql).Server.Stop()
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{"message": fmt.Sprintf("Server have problem with closed this actuality ports ?/|? %d", ports)})
				logger.Error(err)
				return
			}
			ctx.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Server with port %d stoped", ports)})
		}
	} else {
		ctx.JSON(http.StatusFound, gin.H{"message": fmt.Sprintf("Found noting: %d", ports)})
		logger.Info("Server not found")
	}

}
