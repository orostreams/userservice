package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/orostreams/userservice/models"
	"github.com/orostreams/userservice/utils"
)

//UserRepository handles all the database logic for the users.
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() Interface {
	return UserRepository{
		db: utils.ActiveConnection.Db,
	}
}

// FindAll fetches all users from the database.
func (repo UserRepository) FindAll() ([]models.Interface, error) {
	var usersModelSlice []models.User
	if err := repo.db.Find(&usersModelSlice).Error; err != nil {
		return []models.Interface{}, err
	}
	users := make([]models.Interface,len(usersModelSlice))
	for i,v := range usersModelSlice{
		users[i] = models.Interface(v)
	}
	return users, nil
}

//FindByID fetches a single user from the database by id.
func (repo UserRepository) FindByID(id int) (models.Interface, error) {
	user := models.User{}
	if err := repo.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Update takes a user object and update the user.
func (repo UserRepository) Update(instance models.Interface) (models.Interface, error) {
	if err := repo.db.Save(instance).Error; err != nil {
		interfaceValue := models.Interface(instance)
		return interfaceValue, err
	}
	updatedUser := models.Interface(instance)
	return updatedUser, nil
}

//Delete takes a user object and delete the user.
func (repo UserRepository) Delete(instance models.Interface) error {
	if err := repo.db.Delete(instance).Error; err != nil {
		return err
	}
	return nil
}

//Create take a user object and save the user.
func (repo UserRepository) Create(instance models.Interface) (models.Interface, error) {
	instanceAsLocal := instance.(models.User)
	if err := repo.db.Create(&instanceAsLocal).Error; err != nil {
		return instance, err
	}
	return instance, nil
}
