package storage

import "github.com/akriventsev/sample-rest-layered/source/internal/storage/models"

type BaseRepo[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	FindByID(id string) (*T, error)
	FindAll() ([]*T, error)
	DeleteByID(id string) error
}

type OrdersRepo interface {
	BaseRepo[models.Order]
	CreateUserOrder(userID string, items []models.OrderItem) (*models.Order, error)
	FindUserOrders(userID string) ([]models.Order, error)
}

type UsersRepo interface {
	BaseRepo[models.User]
	FindByLoginPassword(login string, passwword string) (*models.User, error)
}

type Storage struct {
	Users    UsersRepo
	Products BaseRepo[models.Product]
	Orders   OrdersRepo
}
