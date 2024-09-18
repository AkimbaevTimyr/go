package requests

type OrderRequest struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Price   float64 `json:"price"`
}
