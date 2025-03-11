package models

import "time"

type User struct {
	ID                int        `gorm:"column:id;primaryKey;autoIncrement"`
	Email             string     `gorm:"column:email;unique;not null"`
	Password          string     `gorm:"column:password;not null"`
	Name              string     `gorm:"column:name;not null"`
	Role              string     `gorm:"column:role;not null"`
	CreatedAt         *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt         *time.Time `gorm:"column:deleted_at;index"`
	PasswordChangedAt *time.Time `gorm:"column:password_changed_at"`
}

type Token struct {
	ID        int        `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int        `gorm:"column:user_id;not null"`
	Token     string     `gorm:"column:token;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
