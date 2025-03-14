package store

import (
	"github.com/asibulhasanshanto/go_api/internal/models"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetUserByField(field string, value string) (*models.User, error) {
	var user models.User
	if err := us.db.Where(field+" = ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := us.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) CreateUser(user *models.User) error {
	if err := us.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
