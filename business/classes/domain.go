package classes

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
}
