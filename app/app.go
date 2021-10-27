package app

import (
	"golara/app/config"
	"golara/app/database"
	"golara/app/httpServer"
)

func Boot() {
	config.InitConfig()
	database.InitDatabase()
	httpServer.InitHttpServer()
}
