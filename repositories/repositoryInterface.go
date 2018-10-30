package repositories

import (
	"github.com/jinzhu/gorm"
)

//Repository an interface that all repositories must satisify
type Repository interface {
	FindAll() ([]*gorm.Model, error)
	FindByID(id int) (*gorm.Model, error)
	Update(instance *gorm.Model) (*gorm.Model, error)
	Delete(instance *gorm.Model) error
	create(instance *gorm.Model) (*gorm.Model, error)
}
