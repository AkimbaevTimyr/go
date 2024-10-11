package resources

import (
	"akimbaev/helpers"
	"akimbaev/models"
	"time"
)

func Category(model *models.Category) helpers.Envelope {
	return helpers.Envelope{
		"id":          model.ID,
		"name":        model.Name,
		"description": model.Description,
		"is_active":   model.IsActive,
		"created_at":  model.CreatedAt.Format(time.ANSIC),
		"updated_at":  model.UpdatedAt.Format(time.ANSIC),
	}
}
