package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/pkg/Database/VirtualSql"
	intternal "project/pkg/Server"
	"project/pkg/handler/mapHadler"
	"project/pkg/logger"
)

func (h *Handler) createAnyServer(port, group, dir string, handlers gin.HandlerFunc) error {
	//handlers gin.HandlerFunc
	//port := strconv.Itoa(port.GetPort())

	r := intternal.NewInternalServer(0, port)
	err := r.Run(group, dir, handlers)
	if err != nil {
		return err
	}

	//log.Print(port)
	return nil
}
func functionTrain(cfg *ServerConfig) gin.HandlerFunc {
	var store *VirtualSql.VirtualMySQLDatabase

	return func(ctx *gin.Context) {
		server, found := ServerTree.Get(int(cfg.Port))
		if server == nil {
			logger.Errorf("server is nil %d", cfg.Port)
			return
		}
		if !found {
			logger.Infof("Server not found port %d", cfg.Port)
			return
		}
		bd, err := openVirtualSql(store, server.(mapHadler.ListServerSql).Config)
		if err != nil {
			return
		}
		allTables, err := retrieveAllDataFromAllTables(bd)
		if err != nil {

			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": allTables})
	}

}
