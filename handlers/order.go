package handlers

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/resources"
	"akimbaev/response"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	type params struct {
		Title   string  `json:"title"`
		Content string  `json:"content"`
		Price   float64 `json:"price"`
	}

	var p params

	json.NewDecoder(r.Body).Decode(&p)

	NewOrder := models.Order{
		Title:   p.Title,
		Content: p.Content,
		Price:   p.Price,
		UserId:  userClaims.UserID,
	}

	database.DB.Create(&NewOrder)

	response.Json(w, http.StatusOK, resources.OrderResource(NewOrder))
}
