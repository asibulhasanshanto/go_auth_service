package api

import (
	"github.com/asibulhasanshanto/go_api/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, log *zap.Logger) *gin.RouterGroup {
	root := r.Group("/api")

	v1 := root.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.Signup)
			auth.POST("/login", authHandler.Login)
			auth.GET("/refresh-access-token", authHandler.RefreshAccessToken)
		}

	}
	return root
}
