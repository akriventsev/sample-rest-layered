package application

import (
	"fmt"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/entities"
	"github.com/akriventsev/sample-rest-layered/source/internal/storage"
	"github.com/akriventsev/sample-rest-layered/source/internal/storage/models"
)

type IApplication interface {
	CreateUser(u entities.User) (*entities.User, error)
	Login(login string, password string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	CreateOrder(order entities.Order) (*entities.Order, error)
}

type Application struct {
	storage *storage.Storage
}

func NewApplication(s *storage.Storage) (*Application, error) {
	if s == nil {
		return nil, fmt.Errorf("cannot start with nil storage")
	}

	return &Application{
		storage: s,
	}, nil
}

func (app Application) CreateUser(u entities.User) (*entities.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	user := models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Married:   u.Married,
		Password:  u.Password,
		Login:     u.Login,
		Birthday:  u.Birthday,
	}

	err := app.storage.Users.Create(&user)

	return &entities.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Married:   user.Married,
		Password:  user.Password,
		Login:     user.Login,
		Birthday:  user.Birthday,
	}, err
}

func (app Application) Login(login string, password string) (*entities.User, error) {
	u, err := app.storage.Users.FindByLoginPassword(login, password)

	if err != nil {
		return nil, err
	}

	user := entities.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Married:   u.Married,
		Birthday:  u.Birthday,
	}

	return &user, nil
}

func (app Application) GetUserByID(id string) (*entities.User, error) {
	u, err := app.storage.Users.FindByID(id)

	if err != nil {
		return nil, err
	}

	user := entities.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Married:   u.Married,
		Birthday:  u.Birthday,
	}

	return &user, nil
}

func (app Application) CreateOrder(order entities.Order) (*entities.Order, error) {
	items := []models.OrderItem{}

	for _, orderItem := range order.Items {
		items = append(items, models.OrderItem{
			ProductID: orderItem.ProductID,
			Price:     orderItem.Price,
			Quantity:  orderItem.Quantity,
		})
	}

	newOrder, err := app.storage.Orders.CreateUserOrder(order.UserID, items)

	if err != nil {
		return nil, err
	}

	created := entities.Order{
		ID:     newOrder.ID,
		UserID: order.UserID,
		Items:  []entities.OrderItem{},
	}

	for _, item := range newOrder.ItemsAsSlice() {
		created.Items = append(created.Items, entities.OrderItem{
			ProductID: item.ProductID,
			Price:     item.Price,
			Quantity:  item.Quantity,
		})
	}

	return &created, nil
}
