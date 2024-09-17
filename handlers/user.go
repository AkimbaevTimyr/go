package handlers

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/response"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	p := params{}
	err := decoder.Decode(&p)

	if err != nil {
		response.Json(w, http.StatusBadRequest, "Invalid JSON")
		fmt.Println("Decoding error:", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

	if err != nil {
		response.Json(w, http.StatusInternalServerError, "Error while hashing password")
	}

	NewUser := models.User{
		Email:    p.Email,
		Name:     p.Name,
		Password: string(hashedPassword),
	}

	database.DB.Create(&NewUser)

	response.Json(w, http.StatusOK, NewUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	Id := r.FormValue("id")

	result := database.DB.Delete(&models.User{}, Id)

	if result.Error != nil {
		response.Json(w, http.StatusInternalServerError, result.Error.Error())
	} else if result.RowsAffected == 0 {
		response.Json(w, http.StatusNotFound, "User not found")
	} else {
		response.Json(w, http.StatusOK, "User deleted successfully")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	Id := r.FormValue("id")

	var User models.User

	if database.DB == nil {
		response.Json(w, http.StatusInternalServerError, "Database connection is nil")
		return
	}

	result := database.DB.First(&User, Id)

	userNotFound(result.Error, w)

	response.Json(w, http.StatusOK, User)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)

	p := params{}

	decoder.Decode(&p)

	Id := r.FormValue("id")

	var User models.User

	result := database.DB.First(&User, Id)

	userNotFound(result.Error, w)

	if p.Email != "" {
		User.Email = p.Email
	}

	if p.Name != "" {
		User.Name = p.Name
	}

	if p.Password != "" {
		User.Password = p.Password
	}

	database.DB.Save(&User)

	response.Json(w, http.StatusOK, User)

}

func userNotFound(err error, w http.ResponseWriter) {
	if err != nil {
		response.Json(w, http.StatusNotFound, "User not found")
		return
	}
}
