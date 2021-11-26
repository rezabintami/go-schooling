package images

import (
	"go-schooling/business/images"
	"time"

	"gorm.io/gorm"
)

type Images struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Images) toDomain() images.Domain {
	return images.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Path:      rec.Path,
		CreatedAt: rec.CreatedAt,
	}
}

func fromDomain(imageDomain images.Domain) *Images {
	return &Images{
		ID:        imageDomain.ID,
		Name:      imageDomain.Name,
		Path:      imageDomain.Path,
		CreatedAt: imageDomain.CreatedAt,
	}
}