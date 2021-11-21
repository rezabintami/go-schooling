package articles

import (
	"go-schooling/drivers/databases/category"
	"go-schooling/drivers/databases/images"
	"time"
)

type Articles struct {
	Id         int
	Title      string
	Content    string `gorm:"column:content_data"`
	CategoryID int
	Category   category.Category
	ImageID    int
	Images     images.Images
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
