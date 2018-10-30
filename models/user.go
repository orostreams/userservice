package models

import (
	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
)

//User a mapping of the user table
type User struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string `gorm:"type:varchar(100);unique_index"`
	Password       string
	DateOfBirth    string
	PhoneNumber    string
	GlobalHashCode string `gorm:"type:varchar(100);unique_index"`
	Role           Role
	RoleID         int
}

//ToMap converts the user struct to a map
func (user *User) ToMap() map[string]interface{} {
	return structs.Map(user)
}
