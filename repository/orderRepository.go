package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests/order"
	"fmt"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetUserOrders(id int, params order.IndexRequest) (*[]models.Post, *helpers.Error)
	CreatePost(id int, request models.Post) (*models.Post, *helpers.Error)
	GetById(id int) (*models.Post, *helpers.Error)
	DeleteById(id int) *helpers.Error
}

type postRepo struct{}

func NewPostRepository() PostRepository {
	return &postRepo{}
}

func (r *postRepo) GetUserOrders(id int, params order.IndexRequest) (*[]models.Post, *helpers.Error) {

	var posts []models.Post

	database.DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return database.Paginate(db, params.Page, params.Sort, params.Count)
	}).Find(&posts)

	return &posts, nil
}

func (r *postRepo) CreatePost(id int, request models.Post) (*models.Post, *helpers.Error) {

	if err := database.DB.Create(&request).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "interval server error"}

	}

	return &request, nil
}

func (r *postRepo) GetById(id int) (*models.Post, *helpers.Error) {
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("post with id %d not found", id)}
	}
	return &post, nil
}

func (r *postRepo) DeleteById(id int) *helpers.Error {
	result := database.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "interval server error"}
	}

	if result.RowsAffected == 0 {
		return &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("post with id %d not found", id)}
	}
	return nil
}
