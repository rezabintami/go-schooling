package articles

import (
	"go-schooling/drivers/databases/category"
	"time"
)

type Articles struct {
	Id         int
	Title      string
	Content    string `gorm:"column:content_data"`
	CategoryID int
	Category   category.Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
}