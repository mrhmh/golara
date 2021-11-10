package repositories

import (
	"golara/app/helpers"
	"golara/app/models/challenge"
	"golara/app/models/scopes"
	"golara/core/facades"
)

type challengeRepository struct{}

func (r challengeRepository) All(filters map[string][]string) []challenge.Challenge {

	var entities []challenge.Challenge

	model := new(challenge.Challenge)

	query := facades.DB().Model(entities)

	for key, val := range filters {
		if helpers.InArray(key, model.Filterable()) {
			query.Where(key+" = ?", val)
		}
	}

	page := "1"
	perPage := "10"
	if len(filters["page"]) > 0 {
		page = filters["page"][0]
	}

	if len(filters["per_page"]) > 0 {
		perPage = filters["per_page"][0]
	}

	query.Scopes(scopes.Paginate(page, perPage))
	query.Find(&entities)

	return entities
}

func NewChallengeRepository() *challengeRepository {
	return &challengeRepository{}
}
