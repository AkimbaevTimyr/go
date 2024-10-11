package controller

import (
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/resources"
	"akimbaev/response"
	"akimbaev/service"
	"net/http"
)

type Category struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) *Category {
	return &Category{
		service: service,
	}
}

// что-то сделать с названиями ошибок
func (c *Category) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	e := helpers.ReadJson(r, w, &category)

	if e != nil {
		response.Json(w, e.HTTPStatus(), e.Details())
		return
	}

	msg, err := helpers.ValidateStruct(category)

	if err != nil {
		response.Json(w, http.StatusBadRequest, msg)
		return
	}

	model, er := c.service.Create(category)

	if er != nil {
		response.Json(w, er.HTTPStatus(), er.Details())
	}

	response.Json(w, http.StatusOK, resources.Category(model))

}
