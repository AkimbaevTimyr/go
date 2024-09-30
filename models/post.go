package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null" validate:"required"`
	Description string `json:"description" gorm:"not null" validate:"required"`
	PlanId      int    `json:"plan_id" gorm:"not null" validate:"required"`
	Plan        Plan   `gorm:"foreignkey:PlanId"`
	Media       string `json:"media" gorm:"default:null" validate:"required"`
}
