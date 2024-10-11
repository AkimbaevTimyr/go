package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null" validate:"required,min=10,max=50"`
	Description string `json:"description" gorm:"not null" validate:"required,min=10,max=50"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}
