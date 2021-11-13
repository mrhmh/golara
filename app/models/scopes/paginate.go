package scopes

import (
	"golara/app/helpers"
	"gorm.io/gorm"
)

func Paginate(filters map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page := helpers.ConvertToInt(filters["page"])
		perPage := helpers.ConvertToInt(filters["per_page"])

		switch {
		case page <= 0:
			page = 1
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
