package routes

import (
	"golara/app/http/controllers/v1"
	"golara/core/contracts"
)

func GetRoutes() []contracts.Route {

	challengeControllerV1 := v1.NewChallengeController()

	return []contracts.Route{
		{
			Method:     "GET",
			Path:       "/challenges",
			Controller: challengeControllerV1.All,
		},
	}
}
