package teachers

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	NIP       string
	Photo     string
	Roles     string
	Sso       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Store(ctx context.Context, data *Domain) error
}