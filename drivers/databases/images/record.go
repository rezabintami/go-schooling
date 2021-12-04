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

func (rec *Images) ToDomain() (res *images.Domain) {
	if rec != nil {
		res = &images.Domain{
			ID:        rec.ID,
			Name:      rec.Name,
			Path:      rec.Path,
			CreatedAt: rec.CreatedAt,
		}
		return res
	}
	return nil
}

func fromDomain(imageDomain images.Domain) *Images {
	return &Images{
		ID:        imageDomain.ID,
		Name:      imageDomain.Name,
		Path:      imageDomain.Path,
		CreatedAt: imageDomain.CreatedAt,
	}
}
