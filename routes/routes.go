package routes

import (
	"golara/app/contracts"
	"golara/http/controllers"
)

func GetRoutes() []contracts.Route {

	questionController := new(controllers.QuestionController)

	return []contracts.Route{
		{
			Method:     "GET",
			Path:       "/",
			Controller: controllers.Index,
		},
		{
			Method:     "GET",
			Path:       "/questions",
			Controller: questionController.Index,
		},
		{
			Method:     "GET",
			Path:       "/questions/:id",
			Controller: questionController.Show,
		},
	}
}
