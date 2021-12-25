package articles

import (
	"context"
	"go-schooling/business/articles"

	"gorm.io/gorm"
)

type mysqlArticlesRepository struct {
	Conn *gorm.DB
}

func NewMySQLArticlesRepository(conn *gorm.DB) articles.Repository {
	return &mysqlArticlesRepository{
		Conn: conn,
	}
}

func (repository *mysqlArticlesRepository) Fetch(ctx context.Context, page, perpage int) ([]articles.DomainFromArticles, int, error) {
	rec := []Articles{}

	offset := (page - 1) * perpage
	err := repository.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []articles.DomainFromArticles{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []articles.DomainFromArticles{}, 0, err
	}

	var domainArticles []articles.DomainFromArticles
	for _, value := range rec {
		domainArticles = append(domainArticles, value.toDomainFromArticles())
	}
	return domainArticles, int(totalData), nil
}

func (repository *mysqlArticlesRepository) GetByID(ctx context.Context, id int) (articles.DomainFromArticles, error) {
	rec := Articles{}
	err := repository.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return articles.DomainFromArticles{}, err
	}
	return rec.toDomainFromArticles(), nil
}

func (repository *mysqlArticlesRepository) GetByTitle(ctx context.Context, title string) (articles.DomainFromArticles, error) {
	rec := Articles{}
	err := repository.Conn.Where("title = ?", title).First(&rec).Error
	if err != nil {
		return articles.DomainFromArticles{}, err
	}
	return rec.toDomainFromArticles(), nil
}

func (repository *mysqlArticlesRepository) Store(ctx context.Context, articleDomain *articles.Domain) (int, error) {
	rec := fromDomain(articleDomain)

	result := repository.Conn.Preload("CategoryArticles").Preload("Category").Create(&rec)
	if result.Error != nil {
		return 0, result.Error
	}

	return rec.ID, nil
}

func (repository *mysqlArticlesRepository) Update(ctx context.Context, articleDomain *articles.Domain, id int) error {
	articleUpdate := fromDomain(articleDomain)

	result := repository.Conn.Preload("CategoryArticles").Preload("Category").Where("id = ?", id).Updates(&articleUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
