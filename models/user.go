package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `gorm:"unique; not null"`
	Name            string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	EmailVerifiedAt time.Time `gorm:"default:null"`
}
