package store

import (
	"github.com/asibulhasanshanto/go_api/internal/models"
	"gorm.io/gorm"
)

type TokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) *TokenStore {
	return &TokenStore{
		db: db,
	}
}

func (ts *TokenStore) StoreRefreshToken(token string, userId int) error {
	var tokenObj models.Token
	tokenObj.Token = token
	tokenObj.UserID = userId
	if err := ts.db.Create(&tokenObj).Error; err != nil {
		return err
	}
	return nil
}

func (ts *TokenStore) DeleteRefreshToken(userId int) error {
	if err := ts.db.Where("user_id = ?", userId).Delete(&models.Token{}).Error; err != nil {
		return err
	}
	return nil
}

func (ts *TokenStore) UpdateRefreshToken(token string, userId int) error {
	if err := ts.db.Model(&models.Token{}).Where("user_id = ?", userId).Update("token", token).Error; err != nil {
		return err
	}
	return nil
}
