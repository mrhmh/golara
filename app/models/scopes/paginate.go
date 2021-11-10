package scopes

import (
	"gorm.io/gorm"
	"strconv"
)

func Paginate(strPage, strPageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(strPage)
		perPage, _ := strconv.Atoi(strPageSize)

		if page == 0 {
			page = 1
		}

		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
