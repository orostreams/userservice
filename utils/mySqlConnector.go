package utils

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ntwarijoshua/orostreams/models"
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
		return &DatabaseConnector{}, err
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
		return nil, errors.New("database type not supported")

	}
}

func RunMigrations() {
	if ActiveConnection != nil {
		ActiveConnection.Db.AutoMigrate(
			&models.User{},
			&models.Role{},
		)
	}

}
