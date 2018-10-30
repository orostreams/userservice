package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	Slug string
}
