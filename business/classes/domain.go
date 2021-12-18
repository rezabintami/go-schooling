package classes

import (
	"context"
	"database/sql"
	"go-schooling/business/teachers"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	TeacherID sql.NullInt64
	Teachers  *teachers.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (*Domain, error)
}

type Repository interface {
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (*Domain, error)
}
