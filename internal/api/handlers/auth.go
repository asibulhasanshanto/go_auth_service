package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	log *zap.Logger
}

func NewAuthHandler(log *zap.Logger) *AuthHandler {
	return &AuthHandler{log: log}
}

func (ah *AuthHandler) Signup(ctx *gin.Context) {
	ah.log.Info("signup handler")
	ctx.JSON(200, gin.H{
		"message": "signup",
	})
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	ah.log.Info("login handler")
	ctx.JSON(200, gin.H{
		"message": "login",
	})
}