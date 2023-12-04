package gormpg

import "gorm.io/gorm"

type GormRepo[T any] struct {
	DB *gorm.DB
}

func (g *GormRepo[T]) Create(entity *T) error {
	return g.DB.Create(entity).Error
}

func (g *GormRepo[T]) Update(entity *T) error {
	return g.DB.Updates(entity).Error
}

func (g *GormRepo[T]) FindByID(id string) (*T, error) {
	val := new(T)
	tx := g.DB.Where("id=?", id).First(val)

	return val, tx.Error
}

func (g *GormRepo[T]) FindAll() ([]*T, error) {
	val := []*T{}
	tx := g.DB.Find(val)

	return val, tx.Error
}

func (g *GormRepo[T]) DeleteByID(id string) error {
	return g.DB.Delete(new(T), id).Error
}
