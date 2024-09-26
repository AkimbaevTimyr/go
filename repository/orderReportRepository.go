package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	report "akimbaev/requests/reports"

	"gorm.io/gorm"
)

type OrderReportRepository interface {
	Create(orderId, userId uint) (*models.OrderReport, *helpers.Error)
	GetReportsByUserId(userId uint, params report.IndexRequest) (*[]models.OrderReport, *helpers.Error)
}

type orderReportRepository struct{}

func NewOrderReportRepository() OrderReportRepository {
	return &orderReportRepository{}
}

func (r *orderReportRepository) Create(orderId, userId uint) (*models.OrderReport, *helpers.Error) {

	report := models.OrderReport{
		UserId:  userId,
		OrderId: orderId,
	}

	if err := database.DB.Create(&report).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	return &report, nil
}

func (r *orderReportRepository) GetReportsByUserId(userId uint, params report.IndexRequest) (*[]models.OrderReport, *helpers.Error) {

	var reports []models.OrderReport

	if err := database.DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return database.Paginate(db, params.Page, params.Sort, params.Count)
	}).Preload("Order").Preload("User").Where("user_id = ?", userId).Find(&reports).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	return &reports, nil

}
