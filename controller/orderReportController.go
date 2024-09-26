package controller

import (
	"akimbaev/helpers"
	report "akimbaev/requests/reports"
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
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, helpers.Envelope{"message": "report created", "report": resources.ReportResource(report)})
}

func (s *UserReportController) MyReports(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := helpers.ExctractUserFromToken(r)

	q := r.URL.Query()

	params := report.IndexRequest{
		Page:  getQueryInt(q, "page", 1),
		Count: getQueryInt(q, "count", 10),
		Sort:  q.Get("sort"),
	}

	msg, validErr := helpers.ValidateStruct(&params)

	if validErr != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	reports, err := s.service.MyReports(userClaims.UserID, params)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
	}

	response.Json(w, http.StatusOK, resources.ReportsResource(reports))
}

func (c *UserReportController) Show(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	report, err := c.service.Show(id)

	if err != nil {
		response.Json(w, err.HTTPStatus(), err.Details())
		return
	}

	response.Json(w, http.StatusOK, resources.ReportResource(report))

}
