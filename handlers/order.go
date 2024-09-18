package handlers

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests"
	"akimbaev/resources"
	"akimbaev/response"
	"encoding/json"
	"net/http"
)

// r - page - count - sort
func GetOrders(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)
	var orders []models.Order

	database.DB.Scopes(database.Paginate(r)).Where("user_id = ?", userClaims.UserID).Find(&orders)
	response.Json(w, http.StatusOK, resources.OrdersResource(orders))
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	var request requests.OrderRequest

	json.NewDecoder(r.Body).Decode(&request)

	NewOrder := models.Order{
		Title:   request.Title,
		Content: request.Content,
		Price:   request.Price,
		UserId:  userClaims.UserID,
	}

	database.DB.Create(&NewOrder)

	response.Json(w, http.StatusOK, resources.OrderResource(NewOrder))
}
