package images

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Path      string
	CreatedAt time.Time
}

type Usecase interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
