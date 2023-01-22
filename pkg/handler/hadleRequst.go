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
	"sync"
	"time"
)

var ServerTree *redblacktree.Tree

type ServerConfig struct {
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
	mu      sync.Mutex
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
	err = h.Handler.store.Server().DeleteServerFromDB(num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Index) StartVirtualServer(ctx *gin.Context) {

	var serverConfig = new(ServerConfig)

	if err := ctx.ShouldBindJSON(&serverConfig); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errs := port.GetQuestionFreePort("localhost", serverConfig.Port)
	if errs != nil {
		logger.Error(errs)
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": fmt.Sprintf("Server with port %d already use", serverConfig.Port)})
		return
	}
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		defer wg.Done()
		newTree, err := mapHadler.NewServerSql("localhost", stringFMT.StringTitleJoin(serverConfig.Name), serverConfig.Port)
		errCh <- err

		h.mu.Lock()
		ServerTree = newTree
		h.mu.Unlock()

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Server with port %d start", serverConfig.Port)})
		logger.Infof("Server with port %d already in use", serverConfig.Port)
		close(errCh)
	}()
	select {
	case err := <-errCh:
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "problem with start the server"})
			logger.Errorf("Failed to start the server %s", err.Error())
			return
		}
	case <-time.After(5 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "timeout while starting the server"})
		logger.Error("Timeout while starting the server")
		return
	}

	wg.Wait()

}

func (h *Index) CloseVirtualServer(ctx *gin.Context) {

	port, err := h.getPortFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Errorf("Failed to get port from context %s", err.Error())
		return
	}
	server, found := ServerTree.Get(port)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found"})
		logger.Infof("Server not found port %d", port)
		return
	}

	if server == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server is nil"})
		logger.Errorf("server is nil %d", port)
		return
	}
	// Create a channel to receive the error response
	errCh := make(chan error)

	// Use goroutine to run the server stop function concurrently
	go func() {
		errCh <- server.(mapHadler.ListServerSql).Server.Stop()
	}()

	select {
	case err := <-errCh:
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "problem with closing the server"})
			logger.Errorf("Failed to stop the server %s", err.Error())
			return
		}
	case <-time.After(5 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "timeout while stopping the server"})
		logger.Error("Timeout while stopping the server")
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Server with port %d stopped", port)})
}
