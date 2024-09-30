package service

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type SubscriptionService interface {
	Purchase(id, userId int) *helpers.Error
}

type subscriptionService struct {
	repo     repository.SubscriptionRepository
	planRepo repository.PlanRepository
	userRepo repository.UserRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository, planRepo repository.PlanRepository, userRepo repository.UserRepository) SubscriptionService {
	return &subscriptionService{
		repo:     repo,
		planRepo: planRepo,
		userRepo: userRepo,
	}
}

// 1 получаем подписку & юсера
// 2 проверка баланса юсера +
// 3 отнимаем с баланса юсера цену подписки +
// 4 создаем Subscription +
// 5 обновляем поле подписки у юсера в БД +
// 6 создаем запись в редис
func (s *subscriptionService) Purchase(id, userId int) *helpers.Error {
	var subscription models.Subscription
	plan, e := s.planRepo.GetById(id)

	if e != nil {
		return e
	}

	user, err := s.userRepo.GetUserById(userId)

	if err != nil {
		return err
	}

	tx := database.DB.Begin()

	//при покупке почему-то идет ошибка в транзакция ролбакается
	defer func() {
		log.Println("something were wrong, transaction rollback")
		tx.Rollback()
	}()

	if err := tx.Where("user_id = ? AND is_active = ?", userId, true).First(&subscription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("у юсера нету активной подписки")
		}
	}

	if !subscription.IsEmpty() {
		tx.Rollback()
		return &helpers.Error{Code: helpers.PAYMENTREQUIRED, Message: "user have active subscription"}
	}

	if user.Balance < plan.Price {
		tx.Rollback()
		return &helpers.Error{Code: helpers.PAYMENTREQUIRED, Message: "insufficient balance"}
	}

	user.Balance -= plan.Price

	NewSubscription := models.Subscription{
		UserId:    userId,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0),
		PlanId:    id,
		IsActive:  true,
	}

	if err := tx.Create(&NewSubscription).Error; err != nil {
		database.DB.Rollback()
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	if err := tx.Save(&user).Error; err != nil {
		database.DB.Rollback()
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	if err := tx.Commit().Error; err != nil {
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "Ошибка сервера"}
	}

	//проверка успешно ли создалась подписка и только потом идем в редис
	database.RedisInstance.Set(fmt.Sprintf("%v", user.ID), true, 60*time.Minute)

	// логика по отправке смс на email юсера о покупке подписки

	return nil
}
