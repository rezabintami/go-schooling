package category

import (
	"go-schooling/business/category"
	"time"
)

type Category struct {
	ID          int
	Title       string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rec *Category) ToDomain() (res *category.Domain) {
	if rec != nil {
		res = &category.Domain{
			ID:        rec.ID,
			Title:     rec.Title,
			Active:    rec.Active,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}

func fromDomain(classDomain category.Domain) *Category {
	return &Category{
		ID:        classDomain.ID,
		Title:     classDomain.Title,
		Active:    classDomain.Active,
		CreatedAt: classDomain.CreatedAt,
		UpdatedAt: classDomain.UpdatedAt,
	}
}
