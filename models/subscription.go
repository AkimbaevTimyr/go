package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserId    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserId"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	PlanId    int       `gorm:"not null"`
	Plan      Plan      `gorm:"foreignKey:PlanId"`
	IsActive  bool      `gorm:"default:true"`
}

func (s *Subscription) IsEmpty() bool {
	return s.UserId == 0
}
