package middlewares

import (
	"net/http"
	"piclist-server/internal/app/config"
	"piclist-server/internal/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果未启用鉴权，直接放行
		if !config.Auth.Enabled {
			c.Next()
			return
		}

		// 检查 key 参数
		key := c.Query("key")
		if key != config.Auth.SecretKey {
			c.JSON(http.StatusUnauthorized, models.UploadResponse{
				Success: false,
				Message: "Invalid authentication key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
