package model

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDatabase(database *gorm.DB) {
	db = database
}
