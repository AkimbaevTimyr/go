package service

import (
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
)

type CategoryService interface {
	Create(category models.Category) (*models.Category, *helpers.Error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *categoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) Create(category models.Category) (*models.Category, *helpers.Error) {
	model, err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}
	return model, nil
}
