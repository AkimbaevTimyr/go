package controller

import (
	"akimbaev/helpers"
	"akimbaev/response"
	"akimbaev/service"
	"net/http"
	"strconv"
)

type SubscriptionController struct {
	service service.SubscriptionService
}

func NewSubscriptionController(service service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{
		service: service,
	}
}

func (c *SubscriptionController) Purchase(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	userClaims, _ := helpers.ExctractUserFromToken(r)

	err := c.service.Purchase(id, int(userClaims.UserID))

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusCreated, helpers.Envelope{"message": "Subscription purchased"})
}
