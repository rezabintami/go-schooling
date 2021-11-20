package images

import "time"

type Images struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
