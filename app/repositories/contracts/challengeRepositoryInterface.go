package contracts

import (
	"golara/app/models/challenge"
)

type ChallengeRepositoryInterface interface {
	All(filters map[string][]string) []challenge.Challenge
}
