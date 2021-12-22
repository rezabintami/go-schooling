package categoryarticles

import (
	"context"
	"go-schooling/business/categoryarticles"

	"gorm.io/gorm"
)

type mysqlCategoryArticlesRepository struct {
	Conn *gorm.DB
}

func NewMySQLCategoryArticlesRepository(conn *gorm.DB) categoryarticles.Repository {
	return &mysqlCategoryArticlesRepository{
		Conn: conn,
	}
}

func (repository *mysqlCategoryArticlesRepository) Store(ctx context.Context, categoryarticlesDomain *categoryarticles.Domain) error {
	rec := fromDomain(*categoryarticlesDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}