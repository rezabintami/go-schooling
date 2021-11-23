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
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context, active string) ([]Domain, error)
	FindByID(id int) (Domain, error)
}