package handlers

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests"
	"akimbaev/resources"
	"akimbaev/response"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request := requests.LoginRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	User := models.User{}

	result := database.DB.First(&User, "email = ?", request.Email)

	if result.Error != nil {
		response.Json(w, http.StatusNotFound, "User not found")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(request.Password))

	if err != nil {
		response.Json(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	tokenString, err := helpers.CreateToken(User.Email)

	if err != nil {
		response.Json(w, http.StatusInternalServerError, "ERROR")
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	request := requests.RegisterRequest{}

	json.NewDecoder(r.Body).Decode(&request)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		response.Json(w, http.StatusInternalServerError, "Error while hashing password")
	}

	NewUser := models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(hashedPassword),
	}

	database.DB.Create(&NewUser)

	response.Json(w, http.StatusOK, resources.UserResource(NewUser))

	generateCode(NewUser)

	//логика по отправке кода юсеру на почту после регистрации
}

func CheckCode(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}

	request := params{}

	json.NewDecoder(r.Body).Decode(&request)

	var code models.VerificationCode

	if err := database.DB.Where("email = ?", request.Email).Where("code = ?", request.Code).First(&code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Json(w, http.StatusNotFound, "Incorrect code")
		}
	} else {
		User := models.User{}

		database.DB.Where("email = ?", request.Email).First(&User)

		User.EmailVerifiedAt = time.Now()
		database.DB.Save(&User)

		clearCodes(User)

		tokenString, _ := helpers.CreateToken(User.Email)

		response.Json(w, http.StatusOK, map[string]any{
			"message": "auth confirmed",
			"token":   tokenString,
		})
	}
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
