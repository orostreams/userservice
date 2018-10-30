package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ntwarijoshua/orostreams/models"
)

//UserRepository handles all the database logic for the users.
type UserRepository struct {
	db *gorm.DB
}

// FindAll fetches all users from the database.
func (repo *UserRepository) FindAll() ([]*models.User, error) {
	users := []*models.User{}
	if err := repo.db.Find(users).Error; err != nil {
		return users, err
	}
	return users, nil
}

//FindByID fetches a single user from the database by id.
func (repo *UserRepository) FindByID(id int) (*models.User, error) {
	user := models.User{}
	if err := repo.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

//Update takes a user object and update the user.
func (repo *UserRepository) Update(instance *models.User) (*models.User, error) {
	if err := repo.db.Save(instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//Delete takes a user object and delete the user.
func (repo *UserRepository) Delete(instance *models.User) error {
	if err := repo.db.Delete(instance).Error; err != nil {
		return err
	}
	return nil
}

//Create take a user object and save the user.
func (repo *UserRepository) Create(instance *models.User) (*models.User, error) {
	if err := repo.db.Create(instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}
