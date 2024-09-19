package resources

import (
	"akimbaev/models"
)

func UserResource(user models.User) map[string]any {

	orders := []map[string]any{}

	for _, order := range user.Orders {
		orders = append(orders, OrderResource(order))
	}

	return map[string]any{
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"email_verified_at": user.EmailVerifiedAt,
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
		"orders":            orders,
	}
}

func UserResourc(user *models.User) map[string]any {

	orders := []map[string]any{}

	for _, order := range user.Orders {
		orders = append(orders, OrderResource(order))
	}

	return map[string]any{
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"email_verified_at": user.EmailVerifiedAt,
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
		"orders":            orders,
	}
}
