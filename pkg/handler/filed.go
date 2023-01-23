package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (h *Index) GetFileSqL(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer file.Close()

	err = saveFile("upload/", file, header)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("os create file err: %s", err.Error()))
		return
	}
	ctx.String(http.StatusOK, "File uploaded successfully")
}

func saveFile(dir string, file multipart.File, header *multipart.FileHeader) error {
	// Create a new file
	f, err := os.Create(dir + header.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Copy the uploaded file to the created file
	io.Copy(f, file)

	return nil
}
