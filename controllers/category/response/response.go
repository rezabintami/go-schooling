package response

import (
	"go-schooling/business/category"
	"time"
)

type Category struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Active    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain category.Domain) Category {
	return Category{
		Id:        domain.ID,
		Title:     domain.Title,
		Active:    domain.Active,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(categoryDomain []category.Domain) *[]Category {
	allCategory := []Category{}
	for _, value := range categoryDomain {
		category := Category{
			Id:        value.ID,
			Title:     value.Title,
			Active:    value.Active,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		allCategory = append(allCategory, category)
	}
	return &allCategory
}