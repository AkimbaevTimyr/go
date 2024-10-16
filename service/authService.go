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
	Login(request requests.LoginRequest) (string, *helpers.Error)
	Register(request requests.RegisterRequest) (*models.User, *helpers.Error)
	CheckCode(requests.CheckCodeRequest) (string, *helpers.Error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(request requests.LoginRequest) (string, *helpers.Error) {
	User := models.User{}
	result := database.DB.First(&User, "email = ?", request.Email)

	if result.Error != nil {
		return "", &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("user with email %s not found", request.Email)}
	}

	sub, _ := GetUserSubscription(int(User.ID))
	e := CacheSubStatusAndUserId(sub, int(User.ID))

	if e != nil {
		return "", e
	}

	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(request.Password))

	if err != nil {
		return "", &helpers.Error{Code: helpers.UNAUTHORIZED, Message: "invalid password"}
	}

	tokenString, err := helpers.CreateToken(User)

	if err != nil {
		return "", &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	return tokenString, nil
}

func (s *authService) Register(request requests.RegisterRequest) (*models.User, *helpers.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "error while hashing password"}
	}

	NewUser := models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(hashedPassword),
	}

	r := helpers.ValidateEmail(NewUser.Email)
	if !r {
		return nil, &helpers.Error{Code: helpers.INVALIDPAYLOAD, Message: "invalid email address"}
	}

	res := database.DB.Create(&NewUser).Error

	if res != nil {
		return nil, &helpers.Error{Code: helpers.STATUSCONFLICT, Message: res.Error()}
	}

	code := generateCode(NewUser)

	go SendSms(code)

	return &NewUser, nil
}

func (s *authService) CheckCode(request requests.CheckCodeRequest) (string, *helpers.Error) {
	var code models.VerificationCode

	if err := database.DB.Where("email = ?", request.Email).Where("code = ?", request.Code).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &helpers.Error{Code: helpers.UNAUTHORIZED, Message: "incorrect code"}
		}
	}
	User := models.User{}

	database.DB.Where("email = ?", request.Email).First(&User)

	User.EmailVerifiedAt = time.Now()
	database.DB.Save(&User)

	clearCodes(User)

	tokenString, _ := helpers.CreateToken(User)

	return tokenString, nil
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
