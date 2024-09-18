package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Title   string
	Content string
	Price   float64
	UserId  uint
	User    User `gorm:"foreignKey:UserId"`
}
