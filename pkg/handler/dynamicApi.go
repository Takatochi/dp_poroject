package handler

import (
	"github.com/gin-gonic/gin"
	intternal "project/pkg/Server"
)

func (h *Handler) createGETServer(group, dir string, handlers gin.HandlerFunc) {

	go func() {
		r := intternal.NewInternalServer(0)
		r.Run("8086")
	}()

}

func (h *Handler) createPOSTServer(group, dir string, handlers gin.HandlerFunc) {
	data := h.router.Group(group)
	{
		data.GET(dir, handlers)
	}
}
func (h *Handler) test(ctx *gin.Context) {
	//var req struct {
	//	Method string `json:"method"`
	//	Path   string `json:"path"`
	//}
	//fmt.Print(req.Method)
	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	ctx.JSON(200, gin.H{
		"message": "Dynamic API is running",
	})
}

func handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Dynamic API is running",
	})
}
