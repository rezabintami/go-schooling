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

func (repository *mysqlCategoryArticlesRepository) GetByArticleID(ctx context.Context, id int) (categoryarticles.Domain, error) {
	categoryArticles := CategoryArticles{}
	result := repository.Conn.Where("articles_id = ?", id).First(&categoryArticles)
	if result.Error != nil {
		return categoryarticles.Domain{}, result.Error
	}

	return *categoryArticles.ToDomain(), nil
}

func (repository *mysqlCategoryArticlesRepository) GetByCategoryID(ctx context.Context, id int) (categoryarticles.Domain, error) {
	categoryArticles := CategoryArticles{}
	result := repository.Conn.Where("category_id = ?", id).First(&categoryArticles)
	if result.Error != nil {
		return categoryarticles.Domain{}, result.Error
	}

	return *categoryArticles.ToDomain(), nil
}