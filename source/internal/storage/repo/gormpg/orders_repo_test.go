package gormpg_test

import (
	"reflect"
	"testing"

	"github.com/akriventsev/sample-rest-layered/source/internal/storage/models"
	"github.com/mitchellh/mapstructure"
)

func TestOrdersGormRepo_Common(t *testing.T) {
	t.Parallel()

	orderedItems := models.OrderItems{
		Items: []models.OrderItem{
			{ProductID: "1", Price: 1, Quantity: 1},
			{ProductID: "2", Price: 2, Quantity: 1},
			{ProductID: "3", Price: 2, Quantity: 1},
			{ProductID: "4", Price: 2, Quantity: 1},
		},
	}

	joi := orderedItems.AsJSONB()

	orderedItems2 := models.OrderItems{}

	err := mapstructure.Decode(joi, &orderedItems2)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(orderedItems, orderedItems2) {
		t.Fatal("karamba")
	}
}

func TestOrdersGormRepo_Slice(t *testing.T) {
	t.Parallel()

	orderedItems := models.OrderItems{
		Items: []models.OrderItem{
			{ProductID: "1", Price: 1, Quantity: 1},
			{ProductID: "2", Price: 2, Quantity: 1},
			{ProductID: "3", Price: 2, Quantity: 1},
			{ProductID: "4", Price: 2, Quantity: 1},
		},
	}

	joi := orderedItems.AsJSONB()

	order := models.Order{
		ID:    "1",
		Items: joi,
	}

	items := order.ItemsAsSlice()

	if len(items) == 0 {
		t.FailNow()
	}
}
