package database

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))

		sort := q.Get("sort")

		if page <= 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("count"))
		switch {
		case pageSize > 10:
			pageSize = 10
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize).Order("created_at " + sort)
	}
}
