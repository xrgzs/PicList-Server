package app

import (
	"fmt"
	"os"
	"piclist-server/internal/app/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		panic(fmt.Sprintf("Failed to initialize config: %v", err))
	}

	// 确保上传目录存在
	if err := os.MkdirAll(config.Conf.Storage.UploadDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create upload directory: %v", err))
	}

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 注册路由
	SetupRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", config.Conf.Server.Host, config.Conf.Server.Port)
	fmt.Printf("Server starting on %s\n", addr)
	fmt.Printf("Upload directory: %s\n", config.Conf.Storage.UploadDir)
	fmt.Printf("Base URL: %s\n", config.Conf.Storage.BaseURL)
	fmt.Printf("Auth enabled: %v\n", config.Conf.Auth.Enabled)

	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
