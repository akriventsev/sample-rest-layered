package models

import "github.com/lib/pq"

type Product struct {
	ID          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Description string         `json:"description"`
	Tags        pq.StringArray `gorm:"type:text[]"`
	Quantity    uint64         `json:"quantity"`
	Price       uint64         `json:"price"`
}
