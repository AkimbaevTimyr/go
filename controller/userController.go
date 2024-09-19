package controller

import (
	"akimbaev/requests"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserController struct {
	service service.UserService
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{
		service: svc,
	}
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	user, err := c.service.GetUser(id)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, resources.UserResource(user))
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := c.service.DeleteUser(id)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, map[string]any{
		"message": "User deleted",
	})
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	var request requests.UpdateUserRequest
	json.NewDecoder(r.Body).Decode(&request)

	user, err := c.service.UpdateUser(id, request)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, resources.UserResource(user))
}
