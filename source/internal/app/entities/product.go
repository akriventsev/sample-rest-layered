package entities

type Product struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Quantity    uint64   `json:"quantity"`
	Price       uint64   `json:"price"`
}
