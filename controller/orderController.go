package controller

import (
	"akimbaev/helpers"
	"akimbaev/requests"
	"akimbaev/requests/order"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
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

	msg, validErr := helpers.ValidateStruct(&params)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	orders, err := c.service.GetOrders(int(userClaims.UserID), params)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.OrdersResource(orders))
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	var request requests.OrderRequest

	e := helpers.ReadJson(r, w, &request)

	if e != nil {
		response.Json(w, http.StatusBadRequest, e.Error())
	}

	msg, validErr := helpers.ValidateStruct(request)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	order, err := c.service.CreateOrder(int(userClaims.UserID), request)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.OrderResource(order))
}

func (c *OrderController) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	orderId, _ := strconv.Atoi(r.FormValue("id"))
	status := r.FormValue("status")

	err := c.service.ChangeStatus(orderId, status)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
	}
	response.Json(w, http.StatusOK, map[string]any{
		"message": "status changed",
	})
}

func (c *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := c.service.Delete(id)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, map[string]string{
		"message": "Order deleted successfully",
	})
}

func getQueryInt(q url.Values, key string, defaultValue int) int {
	if value := q.Get(key); value != "" {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}
	return defaultValue
}
