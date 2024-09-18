package handlers

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/resources"
	"akimbaev/response"
	"net/http"
	"strconv"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)
	orderId, _ := strconv.Atoi(r.FormValue("id"))

	result := database.DB.First(&models.Order{}, orderId).Error
	if result != nil {
		reportNotFound(w)
		return
	}

	Report := models.OrderReport{
		UserId:  userClaims.UserID,
		OrderId: uint(orderId),
	}

	database.DB.Create(&Report)
	database.DB.Preload("Order").Find(&Report, Report.ID)

	response.Json(w, http.StatusOK, resources.ReportResource(Report))
}

func reportNotFound(w http.ResponseWriter) {
	response.Json(w, http.StatusNotFound, map[string]string{
		"message": "Order not found",
	})
}
