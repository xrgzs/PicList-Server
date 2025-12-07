package controllers

import (
	"net/http"
	"xrpic/internal/models"

	"github.com/gin-gonic/gin"
)

type HeartbeatController struct{}

func NewHeartbeatController() *HeartbeatController {
	return &HeartbeatController{}
}

// Heartbeat 健康检查处理器
func (hc *HeartbeatController) Heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, models.HeartbeatResponse{
		Success: true,
		Result:  "alive",
	})
}
