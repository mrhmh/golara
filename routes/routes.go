package routes

import (
	"golara/app/http/controllers"
	"golara/core/contracts"
)

func GetRoutes() []contracts.Route {

	questionController := new(controllers.QuestionController)

	return []contracts.Route{
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
