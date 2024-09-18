package resources

import "akimbaev/models"

func OrderResource(order models.Order) map[string]any {
	return map[string]any{
		"id":         order.ID,
		"title":      order.Title,
		"content":    order.Content,
		"price":      order.Price,
		"created_at": order.CreatedAt,
		"updated_at": order.UpdatedAt,
	}
}
