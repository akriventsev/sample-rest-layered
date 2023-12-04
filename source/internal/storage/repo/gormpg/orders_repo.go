package gormpg

import (
	"fmt"
	"log/slog"

	"github.com/akriventsev/sample-rest-layered/source/internal/storage/models"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm/clause"
)

type OrdersGormRepo struct {
	GormRepo[models.Order]
}

func (r OrdersGormRepo) CreateUserOrder(userID string, items []models.OrderItem) (*models.Order, error) {
	tx := r.DB.Begin()
	defer tx.Rollback()

	orderedItems := make([]models.OrderItem, 0, len(items))

	mitems := map[string]models.OrderItem{}

	for _, item := range items {
		item := item
		mitems[item.ProductID] = item
	}

	ids := []string{}

	for _, v := range items {
		ids = append(ids, v.ProductID)
	}

	ps := []models.Product{}

	tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id in ?", ids).Find(&ps)

	if tx.Error != nil {
		return nil, fmt.Errorf("not found")
	}

	for _, p := range ps {
		p := p
		if p.Quantity < mitems[p.ID].Quantity {
			return nil, fmt.Errorf("insufficient quantity of product %s", mitems[p.ID].ProductID)
		}

		p.Quantity -= mitems[p.ID].Quantity

		tx.Model(&models.Product{}).Where("id = ?", p.ID).Update("quantity", p.Quantity)

		// tx = tx.Save(&p)

		if tx.Error != nil {
			return nil, fmt.Errorf("db error")
		}

		orderedItems = append(orderedItems, models.OrderItem{
			ProductID: mitems[p.ID].ProductID,
			Price:     p.Price,
			Quantity:  mitems[p.ID].Quantity,
		})
	}

	jsonItems := models.JSONB{}

	err := mapstructure.Decode(models.OrderItems{
		Items: orderedItems,
	}, &jsonItems)

	if err != nil {
		slog.Info("error", "error", err)
		return nil, err
	}

	order := models.Order{
		UserID: userID,
		Items:  jsonItems,
	}

	tx = tx.Create(&order)

	if tx.Error != nil {
		return nil, fmt.Errorf("db error")
	}

	tx = tx.Commit()

	if tx.Error != nil {
		return nil, fmt.Errorf("db error")
	}

	return &order, nil
}

func (r OrdersGormRepo) FindUserOrders(userID string) ([]models.Order, error) {
	orders := []models.Order{}
	tx := r.DB.Where("user_id=?", userID).Find(&orders)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return orders, nil
}
