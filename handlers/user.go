package handlers

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/requests"
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var request requests.UpdateUserRequest

	json.NewDecoder(r.Body).Decode(&request)

	Id := r.FormValue("id")

	var User models.User

	result := database.DB.First(&User, Id)

	userNotFound(result.Error, w)

	if request.Email != "" {
		User.Email = request.Email
	}

	if request.Name != "" {
		User.Name = request.Name
	}

	if request.Password != "" {
		User.Password = request.Password
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
