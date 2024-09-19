package controller

import (
	"akimbaev/helpers"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
	"net/http"
	"strconv"
)

type UserReportController struct {
	service service.OrderReportService
}

func NewUserReportController(svc service.OrderReportService) *UserReportController {
	return &UserReportController{
		service: svc,
	}
}

func (s *UserReportController) Connect(w http.ResponseWriter, r *http.Request) {

	userClaims, _ := helpers.ExctractUserFromToken(r)
	orderId, _ := strconv.Atoi(r.FormValue("id"))

	report, err := s.service.Connect(orderId, userClaims.UserID)

	if err != nil {
		response.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	response.Json(w, http.StatusOK, resources.ReportResource(report))
}
