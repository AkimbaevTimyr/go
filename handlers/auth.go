package handlers

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/requests"
	"akimbaev/response"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

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

	tokenString, err := createToken(User.Email)

	response.Json(w, http.StatusOK, tokenString)
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
