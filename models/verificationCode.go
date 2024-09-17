package models

import "gorm.io/gorm"

type VerificationCode struct {
	gorm.Model
	Code  int    `gorm:"not null"`
	Email string `gorm:"not null"`
}
