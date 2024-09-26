package controller

import (
	"akimbaev/helpers"
	"akimbaev/requests"
	"akimbaev/response"
	"akimbaev/service"
	"net/http"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(svc service.AuthService) *AuthController {
	return &AuthController{
		service: svc,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	request := requests.LoginRequest{}

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, http.StatusBadRequest, helpers.INVALIDPAYLOAD)
		return
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	tokenString, err := c.service.Login(request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	request := requests.RegisterRequest{}

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, http.StatusBadRequest, helpers.INVALIDPAYLOAD)
		return
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	user, err := c.service.Register(request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
	}

	response.Json(w, http.StatusOK, user)
}

func (c *AuthController) CheckCode(w http.ResponseWriter, r *http.Request) {
	request := requests.CheckCodeRequest{}

	e := helpers.ReadJson(r, w, &request)
	if e != nil {
		response.Json(w, http.StatusBadRequest, helpers.INVALIDPAYLOAD)
		return
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	tokenString, err := c.service.CheckCode(request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
	}

	response.Json(w, http.StatusOK, map[string]any{
		"message": "auth confirmed",
		"token":   tokenString,
	})
}
