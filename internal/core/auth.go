package core

import (
	"errors"

	"github.com/asibulhasanshanto/go_api/internal/models"
	"github.com/asibulhasanshanto/go_api/internal/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log       *zap.Logger
	userStore *store.UserStore
	token     *Token
}

func NewAuth(log *zap.Logger, userStore *store.UserStore, token *Token) *Auth {
	return &Auth{log: log, userStore: userStore, token: token}
}

func (a *Auth) FindUserByEmail(email string) (models.User, error) {
	user, err := a.userStore.GetUserByField("email", email)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (a *Auth) FindUserByID(id uint) (models.User, error) {
	user, err := a.userStore.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (a *Auth) CreateUser(user *models.User) error {
	user.Password, _ = a.HashPassword(user.Password)
	user.DeletedAt = nil
	return a.userStore.CreateUser(user)
}

func (a *Auth) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (a *Auth) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *Auth) GetUserIdFromContext(ctx *gin.Context, tokenType string) (int, error) {
	var token string
	var err error

	// Get the token from cookies based on token type
	if tokenType == "access" {
		token, err = ctx.Cookie("access_token")
	} else if tokenType == "refresh" {
		token, err = ctx.Cookie("refresh_token")
	}

	if err != nil {
		a.log.Error("failed to get "+tokenType+" token from cookies", zap.Error(err))
	}

	// Look for req headers if cookie is not set
	if token == "" {
		token = ctx.GetHeader("Authorization")
	}

	if token == "" {
		// Return error if token is not found
		return 0, errors.New(tokenType + " token not found")
	}

	// Validate token with the specified token type
	claims, err := a.token.ValidateToken(token, tokenType)
	if err != nil {
		return 0, err
	}

	// Get user id from claims
	userId := int(claims["user_id"].(float64))

	return userId, nil
}
