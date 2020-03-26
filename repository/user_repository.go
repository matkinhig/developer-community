package repository

import "github.com/matkinhig/developer-community/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	// FindAll() ([]models.User, error)
	// FindById(uint32) (models.User, error)
	// FindByUsername(string) (models.User, error)
	// UpdateByUsername(string) (models.User, error)
	// Update(uint32) (models.User, error)
	// DeleteByUsername(string) error
	// Delete(uint32) error
}
