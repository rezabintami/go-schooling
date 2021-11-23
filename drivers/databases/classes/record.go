package classes

import (
	"go-schooling/business/classes"
	"time"
)

type Classes struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Classes) toDomain() classes.Domain {
	return classes.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(classDomain classes.Domain) *Classes {
	return &Classes{
		ID:        classDomain.ID,
		Name:      classDomain.Name,
		CreatedAt: classDomain.CreatedAt,
		UpdatedAt: classDomain.UpdatedAt,
	}
}
