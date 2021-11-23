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
	Archive     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, active string) ([]Domain, error)
	FindByID(id int) (Domain, error)
}