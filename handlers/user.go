package handlers

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/resources"
	"akimbaev/response"
	"encoding/json"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	Id := r.FormValue("id")

	result := database.DB.Delete(&models.User{}, Id)

	if result.Error != nil {
		response.Json(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		response.Json(w, http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	} else {
		response.Json(w, http.StatusOK, map[string]string{
			"message": "User deleted successfully",
		})
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
	database.DB.Preload("Orders").First(&User, User.ID)

	if result.Error != nil {
		userNotFound(result.Error, w)
		return
	}

	response.Json(w, http.StatusOK, resources.UserResource(User))
}

// TODO Зарефакторить реквест
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

	response.Json(w, http.StatusOK, resources.UserResource(User))

}

func userNotFound(err error, w http.ResponseWriter) {
	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
		return
	}
}
