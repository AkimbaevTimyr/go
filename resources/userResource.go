package resources

import (
	"akimbaev/models"
)

func UserResource(user *models.User) map[string]any {
	return map[string]any{
		"id":                user.ID,
		"email":             user.Email,
		"name":              user.Name,
		"email_verified_at": user.EmailVerifiedAt,
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
	}
}
