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

	//gorm config
	var gormConfig gorm.Config

	//if app on debug, turn on log
	if facades.Config().App.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db1, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		panic(err)
	}

	facades.SetDB(db1)
}
