package models

import (
	"gorm.io/gorm"
)

type Status string

const (
	REJECTED Status = "rejected"
	APPROVED Status = "approved"
	MODERATE Status = "moderation"
)

var StatusMap = map[string]Status{
	"rejected":   REJECTED,
	"approved":   APPROVED,
	"moderation": MODERATE,
}

type Order struct {
	gorm.Model
	Title   string  `gorm:"not null"`
	Content string  `gorm:"not null"`
	Price   float64 `gorm:"not null"`
	UserId  uint    `gorm:"not null"`
	Status  Status  `gorm:"type:order_status;default:'moderation'"`
	User    User    `gorm:"foreignKey:UserId"`
}
