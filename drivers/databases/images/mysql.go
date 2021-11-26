package images

import (
	"context"
	"go-schooling/business/images"

	"gorm.io/gorm"
)

type mysqlImagesRepository struct {
	Conn *gorm.DB
}

func NewMySQLImagesRepository(conn *gorm.DB) images.Repository {
	return &mysqlImagesRepository{
		Conn: conn,
	}
}

func (repository *mysqlImagesRepository) FindByID(ctx context.Context, id int) (images.Domain, error) {
	image := Images{}
	result := repository.Conn.Where("id = ?", id).First(&image)
	if result.Error != nil {
		return images.Domain{}, result.Error
	}
	return image.toDomain(), nil
}

func (repository *mysqlImagesRepository) Store(ctx context.Context, teacherDomain *images.Domain) error {
	rec := fromDomain(*teacherDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}