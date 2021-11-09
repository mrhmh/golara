package core

import (
	"golara/core/config"
	"golara/core/contracts"
	"golara/core/db"
	"golara/core/httpServer"
)

type App struct {
	serviceProviders []contracts.Provider
}

func (a *App) Make() {
	a.registerServiceProviders()
	a.bootServiceProviders()
}

func (a *App) registerServiceProviders() {
	a.serviceProviders = []contracts.Provider{
		new(config.ConfigProvider),
		new(db.DBProvider),
		new(httpServer.HttpServerProvider),
	}
}

func (a *App) bootServiceProviders() {
	for _, provider := range a.serviceProviders {
		provider.Boot()
	}
}
