package models

import (
	"github.com/mitchellh/mapstructure"
)

type Order struct {
	ID     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID string `json:"user_id"`
	Items  JSONB  `gorm:"type:jsonb"`
}

type OrderItem struct {
	ProductID string `json:"product_id" mapstructure:"product_id"`
	Price     uint64 `json:"price"      mapstructure:"price"`
	Quantity  uint64 `json:"quantity"   mapstructure:"quantity"`
}

type OrderItems struct {
	Items []OrderItem `gorm:"type:jsonb" mapstructure:"items"`
}

func (o OrderItems) AsJSONB() JSONB {
	j := JSONB{}
	err := mapstructure.Decode(o, &j)
	if err != nil {
		//TODO: add log
	}

	return j
}

func (o Order) ItemsAsSlice() []OrderItem {
	items := OrderItems{}
	err := mapstructure.Decode(o.Items, &items)
	if err != nil {
		//TODO: add log
	}

	return items.Items
}
