package users

import (
	"context"
	"database/sql"
	"go-schooling/business/classes"
	"go-schooling/business/images"
	"time"
)

type Domain struct {
	ID               int
	Name             string
	Password         string
	ClassID          sql.NullInt64
	Classes          *classes.Domain
	ImageID          sql.NullInt64
	Images           *images.Domain
	Email            string
	NISN             *string
	BirthCertificate *string
	FamilyCard       *string
	Photo            string
	Roles            string
	Status           string
	Graduated        bool
	Sso              bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	Login(ctx context.Context, email, password string, sso bool) (string, error)
	Register(ctx context.Context, data *Domain, sso bool) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}
