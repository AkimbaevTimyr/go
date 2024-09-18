package models

import "gorm.io/gorm"

type OrderReport struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserId"`
	OrderId uint   `gorm:"not null"`
	Order   Order  `gorm:"foreignKey:OrderId"`
	Status  Status `gorm:"type:report_status;default:'moderation'"`
}
