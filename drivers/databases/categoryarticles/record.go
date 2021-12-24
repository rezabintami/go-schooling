package categoryarticles

import "go-schooling/business/categoryarticles"

type CategoryArticles struct {
	ArticlesID string
	CategoryID string
}

func fromDomain(categoryarticlesDomain categoryarticles.Domain) *CategoryArticles {
	return &CategoryArticles{
		ArticlesID: categoryarticlesDomain.ArticlesID,	
		CategoryID: categoryarticlesDomain.CategoryID,
	}
}

func (rec *CategoryArticles) ToDomain() *categoryarticles.Domain {
	return &categoryarticles.Domain{
		ArticlesID: rec.ArticlesID,
		CategoryID: rec.CategoryID,
	}
}