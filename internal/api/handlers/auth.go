package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asibulhasanshanto/go_api/internal/core"
	"github.com/asibulhasanshanto/go_api/internal/models"
	"github.com/asibulhasanshanto/go_api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHandler struct {
	log   *zap.Logger
	auth  *core.Auth
	token *core.Token
}

func NewAuthHandler(log *zap.Logger, auth *core.Auth, token *core.Token) *AuthHandler {
	return &AuthHandler{log: log, auth: auth, token: token}
}

func (ah *AuthHandler) Signup(ctx *gin.Context) {
	// read email and password from request body
	var signupRequest models.SignUpRequest
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

	// check if user already exists
	_, err := ah.auth.FindUserByEmail(signupRequest.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "user already exists",
		})
		return
	}

	// create user
	user := models.User{
		Email:    signupRequest.Email,
		Password: signupRequest.Password,
		Name:     signupRequest.Name,
		Role:     "user",
	}

	if err := ah.auth.CreateUser(&user); err != nil {
		ah.log.Error("failed to create user", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	// generate and send tokens
	tokens, err := ah.token.GenerateToken(map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})

	if err != nil {
		ah.log.Error("failed to generate token", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	// save refresh token
	if err := ah.token.SaveRefreshToken(tokens[1], user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokens[0],
		"refresh_token": tokens[1],
	})
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&loginRequest); err != nil {
		ah.log.Error("failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// vlidate request
	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := utils.PrepareValidationErrors(errors, loginRequest)
		ah.log.Info("validation error ", zap.Any("errorMessages", errorMessages))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	// find user by email
	user, err := ah.auth.FindUserByEmail(loginRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password is incorrect",
		})
		return
	}

	// check password
	if err := ah.auth.VerifyPassword(user.Password, loginRequest.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password is incorrect",
		})
		return
	}

	// generate tokens
	tokens, err := ah.token.GenerateToken(map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})

	// update refresh token
	if err := ah.token.UpdateRefreshToken(tokens[1], user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokens[0],
		"refresh_token": tokens[1],
	})

}
