package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/orostreams/userservice/models"
	"log"
	"os"
	"testing"
	"gopkg.in/testfixtures.v2"
)

var(
	db *gorm.DB
	err error
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M){
	db,err = gorm.Open("sqlite3","/test-data/test.db")
	if err != nil{
		log.Fatal(err)
	}

	//create migrations
	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		)
	fixtures, err = testfixtures.NewFolder(db.DB(), &testfixtures.SQLite{}, "test-data/fixtures")
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestCreate(t *testing.T)  {
	if err = fixtures.Load(); err != nil{
		log.Fatal(err)
	}
	testUser := models.User{
		FirstName: "John",
		LastName: "Doe",
		Email: "johnDoe@gmail.com",
		Password: "#222iii",
		DateOfBirth: "12/12/2010",
		PhoneNumber: "0786932945",
		GlobalHashCode: "#JOHN-DOE",
		RoleID: 1,
	}

	expectedUserId := uint(1)
	expectedUserEmail := testUser.Email

	userRepo := UserRepository{
		db: db,
	}
	response,err := userRepo.Create(testUser)
	if err != nil{
		log.Fatal(err)
	}
	responseAsModel := response.(models.User)
	if responseAsModel.ID != expectedUserId {
		t.Errorf("Expected user id %d got %d",expectedUserId,responseAsModel.ID)

	}else if responseAsModel.Email != expectedUserEmail{
		t.Errorf("Expected user email %s got %s",expectedUserEmail,responseAsModel.Email)
	}

}
