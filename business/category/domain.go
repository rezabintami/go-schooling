package category

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	Title       string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id int) error
	Find(ctx context.Context, active bool) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
	FindByID(id int) (Domain, error)
}
