package core

import (
	"github.com/asibulhasanshanto/go_api/internal/models"
	"github.com/asibulhasanshanto/go_api/internal/store"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log       *zap.Logger
	userStore *store.UserStore
}

func NewAuth(log *zap.Logger, userStore *store.UserStore) *Auth {
	return &Auth{log: log, userStore: userStore}
}

func (a *Auth) FindUserByEmail(email string) (models.User, error) {
	user, err := a.userStore.GetUserByField("email", email)
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
