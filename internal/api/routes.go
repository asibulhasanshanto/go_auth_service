package api

import (
	"net/http"

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
		}
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	return root
}