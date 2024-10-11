package resources

import (
	"akimbaev/models"
	"time"
)

func PostResource(model *models.Post) map[string]any {
	return map[string]any{
		"id":          model.ID,
		"title":       model.Title,
		"description": model.Description,
		"media":       model.Media,
		"planId":      model.PlanId,
		"categoryId":  model.CategoryId,
		"category":    Category(&model.Category),
		"created_at":  model.CreatedAt.Format(time.ANSIC),
		"updated_at":  model.UpdatedAt.Format(time.ANSIC),
	}
}

func PostsResource(posts *[]models.Post) []map[string]any {

	resources := []map[string]any{}

	for _, post := range *posts {
		resources = append(resources, PostResource(&post))
	}

	return resources
}
