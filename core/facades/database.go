package facades

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(d *gorm.DB) {
	db = d
}

func DB() *gorm.DB {
	return db
}
