package articles

import (
	"go-schooling/business/articles"
	"go-schooling/drivers/databases/category"
	"go-schooling/drivers/databases/images"
	"time"
)

type Articles struct {
	ID         int `gorm:"primary_key" json:"id"`
	Title      string
	Content    string `gorm:"column:content_data"`
	CategoryID int
	Category   category.Category
	ImageID    int
	Images     images.Images
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func fromDomain(domain *articles.Domain) *Articles {
	return &Articles{
		ID:         domain.ID,
		Title:      domain.Title,
		Content:    domain.Content,
		CategoryID: domain.CategoryID,
		ImageID:    domain.ImageID,
	}
}

func (rec *Articles) toDomain() articles.Domain {
	return articles.Domain{
		ID:           rec.ID,
		Title:        rec.Title,
		Content:      rec.Content,
		CategoryID:   rec.CategoryID,
		CategoryName: rec.Category.Title,
		ImageID:      rec.ImageID,
		ImagePath:    rec.Images.Path,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
	}
}
