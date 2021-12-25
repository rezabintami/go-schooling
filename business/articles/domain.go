package articles

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Title     string
	Content   string
	Category  []int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain, id int) error
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByTitle(ctx context.Context, title string) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain, id int) error
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByTitle(ctx context.Context, title string) (Domain, error)
}
