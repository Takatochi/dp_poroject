package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"project/pkg/Database/VirtualSql"
	mysqldump "project/pkg/Database/dumpmysql"
	"project/pkg/handler/mapHadler"
	"project/pkg/logger"
)

func (h *Index) GetFileSqL(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer file.Close()
	port, err := h.getPortFromContextPOST(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Errorf("Failed to get port from %s", err.Error())
		return
	}

	if ServerTree.Empty() {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server tree is empty"})
		logger.Errorf("server tree is empty %d", port)
		return
	}

	server, found := ServerTree.Get(port)
	if server == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server is nil"})
		logger.Errorf("server is nil %d", port)
		return
	}
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "server not found"})
		logger.Infof("Server not found port %d", port)
		return
	}

	dir, err := saveFile("upload/", file, header)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("os create file err: %s", err.Error()))
		return
	}

	var bd *VirtualSql.VirtualMySQLDatabase
	err = openVirtualSql(dir, bd, server.(mapHadler.ListServerSql).Config)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("problem with connect : %s", err.Error()))
		return
	}

	ctx.String(http.StatusOK, "File uploaded successfully")
}

func saveFile(dir string, file multipart.File, header *multipart.FileHeader) (string, error) {
	// Create a new file
	dir += header.Filename
	f, err := os.Create(dir)
	if err != nil {
		return "nil", err
	}
	defer f.Close()

	// Copy the uploaded file to the created file
	io.Copy(f, file)

	return dir, nil
}

func openVirtualSql(dir string, db VirtualSql.Database, store *VirtualSql.ConfigVirtual) error {
	//var store *VirtualSql.ConfigVirtual
	//store = &VirtualSql.ConfigVirtual{
	//	DatabaseURL: "",
	//	DriverName:  "mysql",
	//}
	DBStore, err := db.Open(store)
	if err != nil {
		return err
	}
	err = mysqldump.Load(DBStore, dir)
	if err != nil {
		return err
	}
	return nil
}
