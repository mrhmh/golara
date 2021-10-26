package routes

import (
	"golara/app/contracts"
	"golara/http/controllers"
)

func GetRoutes() []contracts.Route {

	routes := []contracts.Route{
		{
			Method:     "GET",
			Path:       "/",
			Controller: controllers.Index,
		},
		{
			Method:     "GET",
			Path:       "/questions",
			Controller: controllers.IndexQuestion,
		},
	}

	return routes
}
