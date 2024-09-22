package controller

import (
	"akimbaev/helpers"
	"akimbaev/requests"
	"akimbaev/requests/order"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

// r - page - count - sort
func (c *OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	q := r.URL.Query()
	params := order.IndexRequest{
		Page:  getQueryInt(q, "page", 1),
		Count: getQueryInt(q, "count", 10),
		Sort:  q.Get("sort"),
	}

	orders, err := c.service.GetOrders(int(userClaims.UserID), params)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, resources.OrdersResource(orders))
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	var request requests.OrderRequest

	json.NewDecoder(r.Body).Decode(&request)

	order, err := c.service.CreateOrder(int(userClaims.UserID), request)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, resources.OrderResource(order))
}

func getQueryInt(q url.Values, key string, defaultValue int) int {
	if value := q.Get(key); value != "" {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}
	return defaultValue
}
