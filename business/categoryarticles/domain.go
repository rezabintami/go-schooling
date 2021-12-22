package categoryarticles

import "context"

type Domain struct {
	ArticlesID string
	CategoryID string
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
}