package app

import (
	"piclist-server/internal/app/config"
	"piclist-server/internal/controllers"
	"piclist-server/internal/middlewares"
	"piclist-server/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// 初始化服务
	fileService := services.NewFileService(config.Conf)
	// 初始化控制器
	uploadController := controllers.NewUploadController(fileService)
	heartbeatController := controllers.NewHeartbeatController()

	r.POST("/upload", middlewares.AuthMiddleware(&config.Conf), uploadController.Upload)
	r.POST("/delete", middlewares.AuthMiddleware(&config.Conf), uploadController.Delete)
	r.GET("/heartbeat", heartbeatController.Heartbeat)
	r.POST("/heartbeat", heartbeatController.Heartbeat)
}
