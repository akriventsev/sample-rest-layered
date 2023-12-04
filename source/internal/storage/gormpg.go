package storage

import (
	"github.com/akriventsev/sample-rest-layered/source/internal/storage/models"
	"github.com/akriventsev/sample-rest-layered/source/internal/storage/repo/gormpg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormPgStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	if err != nil {
		return nil, err
	}

	s := Storage{
		Users: &gormpg.UsersRepo{
			GormRepo: gormpg.GormRepo[models.User]{
				DB: db,
			},
		},
		Products: &gormpg.GormRepo[models.Product]{
			DB: db,
		},
		Orders: &gormpg.OrdersGormRepo{
			GormRepo: gormpg.GormRepo[models.Order]{
				DB: db,
			},
		},
	}

	return &s, nil
}
