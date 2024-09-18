package resources

import (
	"akimbaev/models"
	"time"
)

func ReportResource(report models.OrderReport) map[string]any {
	return map[string]any{
		"id":         report.ID,
		"status":     report.Status,
		"order_id":   report.OrderId,
		"order":      OrderResource(report.Order),
		"created_at": report.CreatedAt.Format(time.ANSIC),
		"updated_at": report.UpdatedAt.Format(time.ANSIC),
	}
}
