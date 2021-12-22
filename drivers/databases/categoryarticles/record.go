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