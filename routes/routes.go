package routes

import (
	"golara/app/contracts"
	"golara/http/controllers"
)

var Routes = []contracts.Route{
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
	{
		Method:     "GET",
		Path:       "/questions/:id",
		Controller: controllers.ShowQuestion,
	},
}
