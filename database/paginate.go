package database

import (
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, page int, sort string, count int) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	switch {
	case count > 10:
		count = 10
	case count <= 0:
		count = 10
	}

	offset := (page - 1) * count
	return db.Offset(offset).Limit(count).Order("created_at " + sort)
}
