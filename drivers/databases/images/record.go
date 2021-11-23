package images

import (
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
