package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"

	"github.com/fatih/structs"

	"github.com/orostreams/userservice/models"
	"github.com/orostreams/userservice/repositories"
	"golang.org/x/crypto/bcrypt"
)

//UsersService a service that handle bussiness logic of all users
type UsersService struct {
	repository repositories.Interface
}

//NewUserService Constructor method

func NewUserService() UsersService {
	return UsersService{
		repository: repositories.NewUserRepository(),
	}
}

//Validate validate if all the user requier
func (service UsersService) Validate(attributes map[string]interface{}) error {
	return nil
}

//Create service method to create users.
func (service UsersService) Create(attributes map[string]interface{}) (map[string]interface{}, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(attributes["Password"].(string)), bcrypt.MinCost)

	if err != nil {
		return attributes, err
	}

	hashedPassword := string(hashedPasswordBytes)
	lname := attributes["LastName"].(string)
	fname := attributes["FirstName"].(string)
	initials := fmt.Sprintf("%s%s", string(fname[1]), string(lname[1]))
	globalHashCode := generateGlobalHashCode(initials)
	RoleId,_ := strconv.Atoi(attributes["RoleID"].(string))
	user := models.User{
		FirstName:      attributes["FirstName"].(string),
		LastName:       attributes["LastName"].(string),
		Email:          attributes["Email"].(string),
		Password:       hashedPassword,
		DateOfBirth:    attributes["DateOfBirth"].(string),
		PhoneNumber:    attributes["PhoneNumber"].(string),
		GlobalHashCode: globalHashCode,
		RoleID:        RoleId,
	}

	userInterface, err := service.repository.Create(user)
	if err != nil {
		return structs.Map(userInterface.(models.User)), err
	}
	user = userInterface.(models.User)
	return structs.Map(user), nil
}

//GetAll returns all users
func (service UsersService) GetAll() ([]map[string]interface{}, error) {
	var users []map[string]interface{}
	usersStructs, err := service.repository.FindAll()
	if err != nil {
		return users, err
	}
	for _, user := range usersStructs {
		userMap := structs.Map(user)
		users = append(users, userMap)
	}
	return users, nil
}

//GetById return a user queried by id
func (service UsersService) GetByID(id int) (map[string]interface{}, error) {
	user := map[string]interface{}{}
	userInterface, err := service.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	user = structs.Map(userInterface.(models.User))
	return user, nil
}

//Update service method used to update a user
func (service UsersService) Update(id int, attributes map[string]interface{}) (map[string]interface{}, error) {
	if value,ok := attributes["RoleID"]; ok{
		attributes["RoleID"],_ = strconv.Atoi(value.(string))
	}
	user := map[string]interface{}{}
	userInterface, err := service.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	userStruct := userInterface.(models.User)

	existingValues := reflect.ValueOf(userStruct)
	existingValuesMap := make([]interface{}, existingValues.NumField())
	for i := 0; i < existingValues.NumField(); i++ {

		existingValuesMap[i] = existingValues.Field(i).Interface()
		for key, value := range attributes {
			if reflect.TypeOf(value) == reflect.TypeOf(1) {
				reflect.ValueOf(&userStruct).Elem().FieldByName(key).SetInt(int64(value.(int)))
			} else {
				reflect.ValueOf(&userStruct).Elem().FieldByName(key).SetString(value.(string))
			}

		}
	}
	userToBeSave := models.Interface(&userStruct)
	updateUserInterface, err := service.repository.Update(userToBeSave)

	if err != nil {
		return user, err
	}

	user = structs.Map(updateUserInterface)
	return user, nil
}

//Delete service method for deleting a user.
func (service UsersService) Delete(id int) error {
	userInterface, err := service.repository.FindByID(id)
	if err != nil {
		return err
	}
	userStruct := userInterface.(models.User)
	userToBeDelete := models.Interface(userStruct)
	if err = service.repository.Delete(userToBeDelete); err != nil {
		return err
	}

	return nil
}

func generateGlobalHashCode(initials string) string {
	randomStringByTime, err := generateRandomString(50)
	if err != nil {
		panic(err)
	}
	finalHashCode := fmt.Sprintf("%s-%s", initials, randomStringByTime)
	return finalHashCode
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
