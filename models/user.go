package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `gorm:"unique; not null"`
	Name            string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	EmailVerifiedAt time.Time `gorm:"default:null"`
	Orders          []Order
}

type UserClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
