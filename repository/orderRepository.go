package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests"
	"akimbaev/requests/order"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetUserOrders(id int, params order.IndexRequest) (*[]models.Order, *helpers.Error)
	CreateOrder(id int, request requests.OrderRequest) (*models.Order, *helpers.Error)
	GetById(id int) (*models.Order, *helpers.Error)
	DeleteById(id int) *helpers.Error
}

type orderRepo struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepo{}
}

func (r *orderRepo) GetUserOrders(id int, params order.IndexRequest) (*[]models.Order, *helpers.Error) {

	var orders []models.Order

	database.DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return database.Paginate(db, params.Page, params.Sort, params.Count)
	}).Where("user_id = ?", id).Find(&orders)

	return &orders, nil
}

func (r *orderRepo) CreateOrder(id int, request requests.OrderRequest) (*models.Order, *helpers.Error) {

	NewOrder := models.Order{
		Title:   request.Title,
		Content: request.Content,
		Price:   request.Price,
		UserId:  uint(id),
	}

	if err := database.DB.Create(&NewOrder).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "interval server error"}

	}

	return &NewOrder, nil
}

func (r *orderRepo) GetById(id int) (*models.Order, *helpers.Error) {
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("order with id %d not found", id)}
	}
	return &order, nil
}

func (r *orderRepo) DeleteById(id int) *helpers.Error {
	result := database.DB.Delete(&models.Order{}, id)

	if result.Error != nil {
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "interval server error"}
	}

	if result.RowsAffected == 0 {
		return &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("order with id %d not found", id)}
	}
	return nil
}
