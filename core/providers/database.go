package providers

import (
	"fmt"
	"golara/core/facades"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBProvider struct{}

func (p *DBProvider) Boot() {

	cfg := facades.Config().Database.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	facades.SetDB(db1)
}
