package utils

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/orostreams/userservice/models"
)

var (
	ActiveConnection *DatabaseConnector
)

type DatabaseConnector struct {
	Db *gorm.DB
}

func newMySQLConnector(path string) (*DatabaseConnector, error) {

	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("Could Not Connect To The DB")
	}
	db.LogMode(true)
	return &DatabaseConnector{
		Db: db,
	}, nil
}

func NewDatabaseConnector(dbType string, path string) (*DatabaseConnector, error) {
	switch databaseType := dbType; databaseType {
	case "mysql":
		return newMySQLConnector(path)

	default:
		return &DatabaseConnector{}, errors.New("database type not supported")

	}
}

func RunMigrations() {
	if ActiveConnection.Db != nil {
		ActiveConnection.Db.AutoMigrate(
			&models.User{},
			&models.Role{},
		)
		seeds := seedInitialData()
		for i,_ := range seeds{
			ActiveConnection.Db.Where(models.Role{Slug: seeds[i].Slug}).FirstOrCreate(&seeds[i])
		}
	}

}

func seedInitialData()[]models.Role{
	adminRole := models.Role{
		Name:"Administrator",
		Slug:"Admin",
	}
	artistRole := models.Role{
		Name:"Publisher(Artist)",
		Slug:"Pub",
	}

	clientRole := models.Role{
		Name:"Subscriber(Client)",
		Slug:"Sub",
	}
	return  []models.Role{
		adminRole,
		artistRole,
		clientRole,
	}
}
