package repositories

import (
	"golara/app/helpers"
	"golara/app/models/challenge"
	"golara/app/models/scopes"
	"golara/core/facades"
)

type challengeRepository struct{}

func (r challengeRepository) All(filters map[string]interface{}) []challenge.Challenge {

	var entities []challenge.Challenge

	model := new(challenge.Challenge)

	query := facades.DB().Model(entities)

	for key, val := range filters {
		if helpers.InArray(key, model.Filterable()) {
			query.Where(key+" = ?", val)
		}
	}

	query.Scopes(scopes.Paginate(filters))
	query.Find(&entities)

	return entities
}

func NewChallengeRepository() *challengeRepository {
	return &challengeRepository{}
}
