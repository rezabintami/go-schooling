package images

import (
	"database/sql"
	"go-schooling/business/images"
	"go-schooling/drivers/databases/articles"
	"go-schooling/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Images struct {
	ID        int `gorm:"primary_key" json:"id"`
	UserID    sql.NullInt64
	Users     *users.Users `gorm:"foreignKey:UserID;references:ID"`
	ArticleID sql.NullInt64
	Articles  *articles.Articles `gorm:"foreignKey:ArticleID;references:ID"`
	Name      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Images) ToDomain() (res *images.Domain) {
	if rec != nil {
		res = &images.Domain{
			ID:        rec.ID,
			Users:     rec.Users.ToDomain(),
			Name:      rec.Name,
			Path:      rec.Path,
			CreatedAt: rec.CreatedAt,
		}
	}
	return res
}

func fromDomain(imageDomain images.Domain) *Images {
	return &Images{
		ID:        imageDomain.ID,
		UserID:    imageDomain.UserID,
		Name:      imageDomain.Name,
		Path:      imageDomain.Path,
		CreatedAt: imageDomain.CreatedAt,
	}
}
