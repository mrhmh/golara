package routes

import (
	"golara/app/http/controllers/v1"
	"golara/core/contracts"
)

func GetRoutes() []contracts.Route {

	questionControllerV1 := new(v1.QuestionController)

	return []contracts.Route{
		{
			Method:     "GET",
			Path:       "/questions",
			Controller: questionControllerV1.Index,
		},
		{
			Method:     "GET",
			Path:       "/questions/:id",
			Controller: questionControllerV1.Show,
		},
	}
}
