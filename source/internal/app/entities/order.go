package entities

type OrderItem struct {
	ProductID string `json:"product_id"`
	Price     uint64 `json:"price"`
	Quantity  uint64 `json:"quantity"`
}

type Order struct {
	ID     string      `json:"id"`
	UserID string      `json:"user_id"`
	Items  []OrderItem `json:"items"`
}
