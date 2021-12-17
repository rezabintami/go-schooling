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

func (rec *Category) ToDomain() category.Domain {
	return category.Domain{
		ID:        rec.ID,
		Title:     rec.Title,
		Active:    rec.Active,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
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
