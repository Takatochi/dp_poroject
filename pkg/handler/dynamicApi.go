package handler

import (
	"github.com/gin-gonic/gin"
	intternal "project/pkg/Server"
)

// var started bool
// var lock sync.Mutex
//var routes = make(map[string]*gin.Engine)

func (h *Handler) createGETServer(group, dir string, handlers gin.HandlerFunc) {
	//var lock sync.Mutex

	//r := h.router
	//data := h.router.Group(group)
	//{
	//	data.GET(dir, handlers)
	//}
	go func() {
		r := intternal.NewInternalServer(0)
		r.Run("8086")
	}()

	//r.GET("/start", func(c *gin.Context) {
	//	var req struct {
	//		Method string `json:"method"`
	//		Path   string `json:"path"`
	//	}
	//	fmt.Print(req.Method)
	//	if err := c.ShouldBindJSON(&req); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}

	//engine := h.router
	//
	//switch req.Method {
	//case "GET":
	//	engine.GET(req.Path, func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": fmt.Sprintf("You have reached the dynamic GET endpoint at '%s'", req.Path),
	//		})
	//	})
	//case "POST":
	//	engine.POST(req.Path, func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": fmt.Sprintf("You have reached the dynamic POST endpoint at '%s'", req.Path),
	//		})
	//	})
	//default:
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Invalid method",
	//	})
	//	return
	//}
	//
	//routes[req.Path] = engine
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"message": fmt.Sprintf("Dynamic API started at '%s'", req.Path),
	//})
	//})

	//r.POST("/stop", func(c *gin.Context) {
	//	var req struct {
	//		Path string `json:"path"`
	//	}
	//
	//	if err := c.ShouldBindJSON(&req); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	if _, exists := routes[req.Path]; !exists {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": fmt.Sprintf("The path '%s' is not currently in use", req.Path),
	//		})
	//		return
	//	}
	//
	//	delete(routes, req.Path)
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": fmt.Sprintf("Dynamic API stopped at '%s'", req.Path),
	//	})
	//})

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
