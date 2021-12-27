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

// func (repository *mysqlCategoryArticlesRepository) GetByArticleID(ctx context.Context, id int) (categoryarticles.Domain, error) {
// 	categoryArticles := CategoryArticles{}
// 	result := repository.Conn.Preload("Category").Where("article_id = ?", id).First(&categoryArticles)
// 	if result.Error != nil {
// 		return categoryarticles.Domain{}, result.Error
// 	}

// 	return *categoryArticles.ToDomain(), nil
// }

func (repository *mysqlCategoryArticlesRepository) GetAllByArticleID(ctx context.Context, id int) ([]categoryarticles.Domain, error) {
	allCategoryArticles := []CategoryArticles{}
	result := repository.Conn.Preload("Category").Where("article_id = ?", id).Find(&allCategoryArticles)
	if result.Error != nil {
		return []categoryarticles.Domain{}, result.Error
	}

	allCategoryArticleDomain := []categoryarticles.Domain{}
	for _, value := range allCategoryArticles {
		allCategoryArticleDomain = append(allCategoryArticleDomain, *value.ToDomain())
	}

	return allCategoryArticleDomain, nil
}

func (repository *mysqlCategoryArticlesRepository) GetAllByCategoryID(ctx context.Context, id int) ([]categoryarticles.Domain, error) {
	allCategoryArticles := []CategoryArticles{}
	result := repository.Conn.Preload("Articles").Where("category_id = ?", id).Find(&allCategoryArticles)
	if result.Error != nil {
		return []categoryarticles.Domain{}, result.Error
	}

	allCategoryArticleDomain := []categoryarticles.Domain{}
	for _, value := range allCategoryArticles {
		allCategoryArticleDomain = append(allCategoryArticleDomain, *value.ToDomain())
	}

	return allCategoryArticleDomain, nil
}

func (repository *mysqlCategoryArticlesRepository) GetByCategoryID(ctx context.Context, id int) (categoryarticles.Domain, error) {
	categoryArticles := CategoryArticles{}
	result := repository.Conn.Preload("Articles").Where("category_id = ?", id).First(&categoryArticles)
	if result.Error != nil {
		return categoryarticles.Domain{}, result.Error
	}

	return *categoryArticles.ToDomain(), nil
}

func (repository *mysqlCategoryArticlesRepository) DeleteByArticleID(ctx context.Context, id int) error {
	categoryArticles := CategoryArticles{}
	result := repository.Conn.Preload("Articles").Preload("Category").Where("article_id = ?", id).Delete(&categoryArticles)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlCategoryArticlesRepository) DeleteByCategoryID(ctx context.Context, id int) error {
	categoryArticles := CategoryArticles{}
	result := repository.Conn.Preload("Articles").Preload("Category").Where("article_id = ?", id).Delete(&categoryArticles)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
