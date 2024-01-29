package paginator

import (
	"math"

	"gorm.io/gorm"
)

type (
	Paginator[T any] struct {
		Limit      int    `json:"-"`
		Page       int    `json:"page"`
		Sort       string `json:"-"`
		TotalPages int    `json:"totalPages"`
		Rows       T      `json:"content"`
	} // @name Paginator

)

func (p *Paginator[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paginator[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Paginator[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Page >= p.TotalPages {
		p.Page = p.TotalPages
	}

	return p.Page
}

func (p *Paginator[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func (p *Paginator[T]) Paginate(value interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(p.Limit)))
	p.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}
