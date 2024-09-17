package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int
	Email    string
	Name     string
	Password string
}
