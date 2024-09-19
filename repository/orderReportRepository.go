package repository

type OrderReportRepository interface {
}

type orderReportRepository struct{}

func NewOrderReportRepository() OrderReportRepository {
	return &orderReportRepository{}
}
