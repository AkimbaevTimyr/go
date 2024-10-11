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
	//parse email & password and set to the login request struct
	request := requests.LoginRequest{}

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, e.HTTPStatus(), e.Details())
		return
	}

	//check struct to correct data
	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	//check user in db & compare password & create token
	tokenString, err := c.service.Login(request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	//send user json data & cookie
	cookie := http.Cookie{Name: "token", Value: tokenString, HttpOnly: true, Secure: false}
	http.SetCookie(w, &cookie)

	response.Json(w, http.StatusCreated, helpers.Envelope{"token": tokenString})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	request := requests.RegisterRequest{}

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, e.HTTPStatus(), e.Details())
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
		response.Json(w, e.HTTPStatus(), e.Details())
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

	response.Json(w, http.StatusOK, helpers.Envelope{"message": "auth confirmed", "token": tokenString})
}
