package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/app/model"
	"project/pkg/handler/mapHadler"
	"project/pkg/logger"
	"project/pkg/port"
	StoreBD "project/pkg/store"
	"sort"
	"strconv"
)

var serverList []mapHadler.ListServerSql

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

	portGet := ctx.Param("port")
	ports, err := strconv.Atoi(portGet)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err})
		logger.Error(err)
		return
	}

	errs := port.GetQuestionFreePort("localhost", ports)
	if errs != nil {
		logger.Error(errs)
		ctx.JSON(http.StatusTooManyRequests, fmt.Sprintf("Server with port %d already use", ports))
		return
	}

	ListServerSql := make(chan []mapHadler.ListServerSql, 1)
	go func() {
		ctx.JSON(http.StatusOK, fmt.Sprintf("Server with port %d start", ports))
		logger.Infof("Server with port %d already use", ports)
		ListServerSql <- mapHadler.NewServerSql("localhost", "alpha", int32(ports))
		close(ListServerSql)
	}()

	//serverList := <-ListServerSql
	serverList = <-ListServerSql

	sort.Slice(serverList, func(i, j int) bool {
		return serverList[i].Port <= serverList[j].Port
	})
	//for _, data := range *serverList {
	//	if data.Port == int32(ports) {
	//		err = data.Server.Stop()
	//		if err != nil {
	//			ctx.JSON(http.StatusAccepted, fmt.Sprintf("Server with port %d stoped", ports))
	//			return
	//		}
	//	}
	//	logger.Info(data)
	//}

}

func (h *Index) CloseVirtualServer(ctx *gin.Context) {
	portGet := ctx.Param("port")
	ports, err := strconv.Atoi(portGet)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err})
		logger.Error(err)
		return
	}

	idx := sort.Search(len(serverList), func(i int) bool {
		return serverList[i].Port >= int32(ports)
	})

	if idx < len(serverList) && serverList[idx].Port == int32(ports) {

		err := serverList[idx].Server.Stop()
		if err != nil {
			logger.Error(err)
			ctx.JSON(http.StatusBadGateway, fmt.Sprintf("Server have problem with closed this actuality ports ?/|? %d", ports))
		}
		ctx.JSON(http.StatusAccepted, fmt.Sprintf("Server with port %d stoped", ports))
	} else {
		ctx.JSON(http.StatusFound, fmt.Sprintf("Found noting: %d/%d", idx, ports))
	}
	//for _, data := range serverList {
	//	if data.Port == int32(ports) {
	//
	//	}
	//}

	//ctx.JSON(http.StatusOK, "ok")
}
