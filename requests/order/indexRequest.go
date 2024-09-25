package order

type IndexRequest struct {
	Page  int    `json:"page" validate:"required"`
	Count int    `json:"count" validate:"required"`
	Sort  string `json:"sort" validate:"required"`
}
