package controller

import (
	"akimbaev/requests"
	"akimbaev/response"
	"akimbaev/service"
	"encoding/json"
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

// TODO сделать глобальную обработку ошибок
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	request := requests.LoginRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	tokenString, err := c.service.Login(request)

	if err != nil {
		response.Json(w, http.StatusNotFound, err.Error())
	}

	response.Json(w, http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	request := requests.RegisterRequest{}

	json.NewDecoder(r.Body).Decode(&request)

	user, err := c.service.Register(request)

	if err != nil {
		response.Json(w, http.StatusNotFound, err.Error())
	}

	response.Json(w, http.StatusOK, user)
}

func (c *AuthController) CheckCode(w http.ResponseWriter, r *http.Request) {
	request := requests.CheckCodeRequest{}

	json.NewDecoder(r.Body).Decode(&request)

	tokenString, err := c.service.CheckCode(request)

	if err != nil {
		response.Json(w, http.StatusNotFound, err.Error())
	}

	response.Json(w, http.StatusOK, map[string]any{
		"message": "auth confirmed",
		"token":   tokenString,
	})
}
