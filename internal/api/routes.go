package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, log *zap.Logger) *gin.RouterGroup {
	root := r.Group("/api")
	
	v1 := root.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	return root
}