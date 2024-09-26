package resources

import (
	"akimbaev/models"
	"time"
)

func ReportResource(report *models.OrderReport) map[string]any {
	return map[string]any{
		"id":         report.ID,
		"status":     report.Status,
		"order_id":   report.OrderId,
		"created_at": report.CreatedAt.Format(time.ANSIC),
		"updated_at": report.UpdatedAt.Format(time.ANSIC),
		"order":      OrderResource(&report.Order),
	}
}

func ReportsResource(orders *[]models.OrderReport) []map[string]any {

	reports := []map[string]any{}

	for _, order := range *orders {
		reports = append(reports, ReportResource(&order))
	}

	return reports
}
