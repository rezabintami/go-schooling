package articles

import (
	"go-schooling/business/articles"
	"time"
)

type Articles struct {
	ID         int `gorm:"primary_key"`
	Title      string
	Content    string `gorm:"column:content_data"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func fromDomain(domain *articles.Domain) *Articles {
	return &Articles{
		ID:         domain.ID,
		Title:      domain.Title,
		Content:    domain.Content,
	}
}

func (rec *Articles) toDomain() articles.Domain {
	return articles.Domain{
		ID:           rec.ID,
		Title:        rec.Title,
		Content:      rec.Content,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
	}
}
