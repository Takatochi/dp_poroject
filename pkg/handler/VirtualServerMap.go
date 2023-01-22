package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Index) getPortFromContext(ctx *gin.Context) (int, error) {
	portGet := ctx.Param("port")
	if portGet == "" {
		return 0, fmt.Errorf("missing port param")
	}
	port, err := strconv.Atoi(portGet)
	if err != nil {
		return 0, fmt.Errorf("invalid port: %v", err)
	}
	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("invalid port range: %d", port)
	}
	return port, nil
}
