package order

type IndexRequest struct {
	Page  int    `json:"page"`
	Count int    `json:"count"`
	Sort  string `json:"sort"`
}
