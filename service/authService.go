package service

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(request requests.LoginRequest) (string, error)
	Register(request requests.RegisterRequest) (*models.User, error)
	CheckCode(requests.CheckCodeRequest) (string, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(request requests.LoginRequest) (string, error) {
	User := models.User{}

	result := database.DB.First(&User, "email = ?", request.Email)

	if result.Error != nil {
		return "", fmt.Errorf("user with email %s not found", request.Email)
	}

	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(request.Password))

	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	tokenString, err := helpers.CreateToken(User)

	if err != nil {
		return "", fmt.Errorf("ERROR")
	}

	return tokenString, nil
}

func (s *authService) Register(request requests.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("error while hashing password")
	}

	NewUser := models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(hashedPassword),
	}

	database.DB.Create(&NewUser)

	generateCode(NewUser)
	return &NewUser, nil

	//логика по отправке кода юсеру на почту после регистрации
}

func (s *authService) CheckCode(request requests.CheckCodeRequest) (string, error) {
	var code models.VerificationCode

	if err := database.DB.Where("email = ?", request.Email).Where("code = ?", request.Code).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("incorrect code")
		}
	} else {
		User := models.User{}

		database.DB.Where("email = ?", request.Email).First(&User)

		User.EmailVerifiedAt = time.Now()
		database.DB.Save(&User)

		clearCodes(User)

		tokenString, _ := helpers.CreateToken(User)

		return tokenString, nil
	}
	return "", fmt.Errorf("ERROR")
}

func generateCode(user models.User) int {
	min := 100000
	max := 999999

	randomNum := rand.Intn(max-min) + min

	NewCode := models.VerificationCode{
		Code:  randomNum,
		Email: user.Email,
	}

	database.DB.Create(&NewCode)
	return randomNum
}

func clearCodes(user models.User) {
	database.DB.Where("email = ?", user.Email).Delete(&models.VerificationCode{})
}
