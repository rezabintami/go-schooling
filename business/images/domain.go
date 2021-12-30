package images

import (
	"context"
	"database/sql"
	"go-schooling/business/users"
	"mime/multipart"
	"time"
)

type Domain struct {
	ID        int
	UserID    sql.NullInt64
	Users     *users.Domain
	Name      string
	Path      string
	CreatedAt time.Time
}

type Usecase interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	GetPresignedURL(ctx context.Context, name string) (string, error)
	Store(ctx context.Context, data *Domain, file *multipart.FileHeader) (string, error)
}

type Repository interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
