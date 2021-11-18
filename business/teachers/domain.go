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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
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
	GetByEmail(ctx context.Context, email string) (Domain, error)
}
