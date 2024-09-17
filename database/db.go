package database

import (
	"akimbaev/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=newpassword dbname=testgo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB.AutoMigrate(&models.User{}, &models.VerificationCode{})
}
