package database

import (
	"fmt"
	"golara/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDatabase() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config().Database.Mysql.UserName, config.Config().Database.Mysql.Password, config.Config().Database.Mysql.Host, config.Config().Database.Mysql.Port, config.Config().Database.Mysql.Database)
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db = db1
}

func DB() *gorm.DB {
	return db
}
