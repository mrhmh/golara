package v1

import (
	"context"
	"golara/app/grpc/pb/v1"
	"golara/app/repositories"
	"golara/app/repositories/contracts"
)

type challengeController struct {
	repository contracts.ChallengeRepositoryInterface
}

func (c *challengeController) All(ctx context.Context, r *v1.IndexRequest) (*v1.ChallengesResponse, error) {

	var filters map[string]interface{}
	entities := c.repository.All(filters)

	var challenges []*v1.ChallengeResponse
	for _, entity := range entities {
		challenges = append(challenges, &v1.ChallengeResponse{
			Id:    uint64(entity.ID),
			Title: entity.Title,
		})
	}

	return &v1.ChallengesResponse{Challenges: challenges}, nil
}

func NewChallengeController() *challengeController {
	return &challengeController{
		repository: repositories.NewChallengeRepository(),
	}
}
