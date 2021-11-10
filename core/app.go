package core

import (
	"golara/core/contracts"
	"golara/core/providers"
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
		new(providers.ConfigProvider),
		new(providers.DBProvider),
		//new(providers.HttpServerProvider),
		new(providers.GrpcServerProvider),
	}
}

func (a *App) bootServiceProviders() {
	for _, provider := range a.serviceProviders {
		provider.Boot()
	}
}
