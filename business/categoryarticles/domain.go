package categoryarticles

import "context"

type Domain struct {
	ArticlesID string
	CategoryID string
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByArticleID(ctx context.Context, id int) (Domain, error)
	GetByCategoryID(ctx context.Context, id int) (Domain, error)
}
