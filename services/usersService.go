package services

import(
	"github.com/ntwarijoshua/orostreams/repositories"
	"github.com/NtwariJoshua/orostreams/models"
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

//UsersService a service that handle bussiness logic of all users
type UsersService struct{
	repository *repositories.Repository
}

//Validate validate if all the user requier
func Validate(attributes map[string]interface)error{
	return nil
}

func Create(attributes map[string]interface)(map[string]interface,error){
	hashedPassword,err := bcrypt.GenerateFromPassword(attributes["email"],bcrypt.MinCost)
	if err != nil{
		return attributes,err
	}

	globalHashCode
	user := models.User{
		FirstName: attributes["FirstName"],
		LastName: attributes["LastName"],
		Email: attributes["email"],
		Password: hashedPassword,
		DateOfBirth: attributes["DateOfBirth"],
		PhoneNumber: attributes["PhoneNumber"],


	}
	return user,nil
}

func GetAll(){}



func generateGlobalHashCode(){
	randomFromTime := rand.NewSource(time.Now().UnixNano())
	
}