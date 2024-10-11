package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string   `json:"title" gorm:"not null" validate:"required"`
	Description string   `json:"description" gorm:"not null" validate:"required"`
	PlanId      int      `json:"plan_id" gorm:"not null" validate:"required"`
	Plan        Plan     `gorm:"foreignKey:PlanId"`
	CategoryId  int      `json:"category_id" gorm:"default:null" validate:"required"`
	Category    Category `gorm:"foreignKey:CategoryId"`
	Media       string   `json:"media" gorm:"default:null" validate:"required"`
}
