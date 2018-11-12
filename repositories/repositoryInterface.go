package repositories

import (
	"github.com/orostreams/userservice/models"
)

//Repository an interface that all repositories must satisify
type Interface interface {
	FindAll() ([]models.Interface, error)
	FindByID(id int) (models.Interface, error)
	Update(instance models.Interface) (models.Interface, error)
	Delete(instance models.Interface) error
	Create(instance models.Interface) (models.Interface, error)
}
