package requests

type AddBalanceRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
