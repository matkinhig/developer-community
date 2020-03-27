package repository

import (
	"github.com/matkinhig/developer-community/helper"
	"github.com/matkinhig/developer-community/models"
)

type UserRepository interface {
	Save(models.User) (models.User, error) //ham nay thay the ham singin
	FindAll(helper.Pagination) ([]models.User, error)
	FindByUserLogin(string) (models.User, error)
	UpdateByUserLogin(string, models.User) error
}
