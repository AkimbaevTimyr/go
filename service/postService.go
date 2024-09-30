package service

import (
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
	"akimbaev/requests/order"
)

type PostService interface {
	GetPosts(id int, params order.IndexRequest) (*[]models.Post, *helpers.Error)
	CreatePost(id int, request models.Post) (*models.Post, *helpers.Error)
	Delete(id int) *helpers.Error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) GetPosts(id int, params order.IndexRequest) (*[]models.Post, *helpers.Error) {

	orders, err := s.repo.GetUserOrders(id, params)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *postService) CreatePost(id int, request models.Post) (*models.Post, *helpers.Error) {

	post, err := s.repo.CreatePost(id, request)

	if err != nil {
		return nil, err
	}

	return post, nil

}

func (s *postService) Delete(id int) *helpers.Error {

	err := s.repo.DeleteById(id)

	if err != nil {
		return err
	}

	return nil

}
