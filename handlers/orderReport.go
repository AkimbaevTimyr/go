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

	tx := database.DB.Begin()

	var order models.Order

	result := tx.First(&models.Order{}, orderId).Error
	if result != nil {
		reportNotFound(w)
		return
	}

	//создание отчета
	Report := models.OrderReport{
		UserId:  userClaims.UserID,
		OrderId: uint(orderId),
	}

	if err := tx.Create(&Report).Error; err != nil {
		tx.Rollback()
		response.Json(w, http.StatusInternalServerError, err.Error())
		return
	}

	tx.First(&order, orderId)
	tx.Preload("User").First(&order, order.ID)

	order.User.Balance -= order.Price
	if err := tx.Save(&order.User).Error; err != nil {
		tx.Rollback()
		response.Json(w, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Preload("Order").Find(&Report, Report.ID)

	tx.Commit()

	response.Json(w, http.StatusOK, resources.ReportResource(Report))
}

func reportNotFound(w http.ResponseWriter) {
	response.Json(w, http.StatusNotFound, map[string]string{
		"message": "Order not found",
	})
}
