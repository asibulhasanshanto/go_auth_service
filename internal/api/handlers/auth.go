package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asibulhasanshanto/go_api/internal/models"
	"github.com/asibulhasanshanto/go_api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHandler struct {
	log *zap.Logger
}

func NewAuthHandler(log *zap.Logger) *AuthHandler {
	return &AuthHandler{log: log}
}

func (ah *AuthHandler) Signup(ctx *gin.Context) {
	// read email and password from request body
	var signupRequest models.UserCreateRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&signupRequest); err != nil {
		ah.log.Error("failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// validate request body
	validate := validator.New()
	if err := validate.Struct(signupRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := utils.PrepareValidationErrors(errors, signupRequest)
		ah.log.Info("validation error", zap.Any("errorMessages", errorMessages))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
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
