package categoryarticles

import (
	"context"
	"go-schooling/business/category"
)

type Domain struct {
	ArticleID  int
	CategoryID int
	Category   *category.Domain
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetAllByArticleID(ctx context.Context, id int) ([]Domain, error)
	// GetByArticleID(ctx context.Context, id int) (Domain, error)
	GetByCategoryID(ctx context.Context, id int) (Domain, error)
	DeleteByArticleID(ctx context.Context, id int) error
	DeleteByCategoryID(ctx context.Context, id int) error
}
