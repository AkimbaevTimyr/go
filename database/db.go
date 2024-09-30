package database

import (
	"akimbaev/models"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const projectDirName = "go"

func Init() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	e := godotenv.Load(string(rootPath) + `/.env`)
	if e != nil {
		log.Fatal(e.Error())
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")
	sllmode := os.Getenv("SLLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", host, user, password, database, port, sllmode)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//что-то сделать с этим запросом в бд
	DB.Exec("CREATE TYPE order_status AS ENUM ('rejected', 'approved', 'moderation');")
	DB.Exec("CREATE TYPE report_status AS ENUM ('rejected', 'approved', 'moderation');")
	DB.AutoMigrate(&models.User{}, &models.VerificationCode{}, &models.Plan{}, &models.Subscription{}, &models.Post{})
}
