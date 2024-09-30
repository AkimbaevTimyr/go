package controller

import (
	"akimbaev/helpers"
	"akimbaev/requests"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
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
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.UserResource(user))
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := c.service.DeleteUser(id)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, helpers.Envelope{"message": "User deleted"})
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	var request requests.UpdateUserRequest

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, e.HTTPStatus(), e.Details())
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	user, err := c.service.UpdateUser(id, request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.UserResource(user))
}

// тестовая версия добавления баланса юсеру
func (c *UserController) AddBalance(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	var request requests.AddBalanceRequest

	helpers.ReadJson(r, w, &request)

	msg, e := helpers.ValidateStruct(request)
	if e != nil {
		response.Json(w, http.StatusBadRequest, msg)
	}

	err := c.service.AddBalance(int(userClaims.UserID), request)
	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, helpers.Envelope{"message": "balance added"})
}
