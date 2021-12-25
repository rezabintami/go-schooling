package categoryarticles

import (
	"go-schooling/business/categoryarticles"
	"go-schooling/drivers/databases/articles"
	"go-schooling/drivers/databases/category"
)

type CategoryArticles struct {
	ArticleID int 
	Articles   *articles.Articles `gorm:"foreignKey:ArticleID;references:ID"`
	CategoryID int
	Category   *category.Category `gorm:"foreignKey:CategoryID;references:ID"`

}

func fromDomain(categoryarticlesDomain categoryarticles.Domain) *CategoryArticles {
	return &CategoryArticles{
		ArticleID: categoryarticlesDomain.ArticleID,	
		CategoryID: categoryarticlesDomain.CategoryID,
	}
}

func (rec *CategoryArticles) ToDomain() *categoryarticles.Domain {
	return &categoryarticles.Domain{
		ArticleID: rec.ArticleID,
		CategoryID: rec.CategoryID,
	}
}