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
	Balance         float64   `gorm:"not null;default:0"`
	EmailVerifiedAt time.Time `gorm:"default:null"`
	Orders          []Order
	Role            string `gorm:"default:null"`
}

type UserClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
