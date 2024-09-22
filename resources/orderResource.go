package resources

import (
	"akimbaev/models"
	"time"
)

func OrderResource(order *models.Order) map[string]any {
	return map[string]any{
		"id":         order.ID,
		"title":      order.Title,
		"content":    order.Content,
		"price":      order.Price,
		"status":     order.Status,
		"created_at": order.CreatedAt.Format(time.ANSIC),
		"updated_at": order.UpdatedAt.Format(time.ANSIC),
	}
}

func OrdersResource(orders *[]models.Order) []map[string]any {

	resources := []map[string]any{}

	for _, order := range *orders {
		resources = append(resources, OrderResource(&order))
	}

	return resources
}
