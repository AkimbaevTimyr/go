package repository

import (
	"akimbaev/helpers"
	"akimbaev/models"
)

type SubscriptionRepository interface {
	GetById(id int) (*models.Subscription, *helpers.Error)
}

type subscriptionRepo struct {
}

func NewSubscriptionRepository() SubscriptionRepository {
	return &subscriptionRepo{}
}

func (r *subscriptionRepo) GetById(id int) (*models.Subscription, *helpers.Error) {

	return nil, nil
}
