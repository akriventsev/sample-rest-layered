package gormpg

import "github.com/akriventsev/sample-rest-layered/source/internal/storage/models"

type UsersRepo struct {
	GormRepo[models.User]
}

func (r UsersRepo) FindByLoginPassword(login string, passwword string) (*models.User, error) {
	user := models.User{}
	tx := r.DB.Where("login=? and password=?", login, passwword).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
