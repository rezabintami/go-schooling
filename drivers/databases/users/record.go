package users

import (
	"database/sql"
	"go-schooling/business/users"
	"go-schooling/drivers/databases/classes"
	"go-schooling/drivers/databases/images"
	"time"
)

type Users struct {
	ID               int `gorm:"primary_key" json:"id"`
	Name             sql.NullString
	Password         sql.NullString
	ClassID          sql.NullInt64
	Classes          *classes.Classes `gorm:"foreignKey:ClassID;references:ID"`
	ImageID          sql.NullInt64
	Images           *images.Images `gorm:"foreignKey:ImageID;references:ID"`
	Email            sql.NullString
	NISN             sql.NullString
	BirthCertificate sql.NullString
	FamilyCard       sql.NullString
	Photo            sql.NullString
	Roles            sql.NullString
	Status           sql.NullString
	Sso              sql.NullBool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (rec *Users) toDomain() *users.Domain {
	return &users.Domain{
		ID:               rec.ID,
		Name:             rec.Name.String,
		Password:         rec.Password.String,
		Classes:          rec.Classes.ToDomain(),
		Images:           rec.Images.ToDomain(),
		Email:            rec.Email.String,
		NISN:             &rec.NISN.String,
		BirthCertificate: rec.BirthCertificate.String,
		FamilyCard:       rec.FamilyCard.String,
		Photo:            rec.Photo.String,
		Roles:            rec.Roles.String,
		Status:           rec.Status.String,
		Sso:              rec.Sso.Bool,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:               userDomain.ID,
		Name:             sql.NullString{String: userDomain.Name, Valid: true},
		Password:         sql.NullString{String: userDomain.Password, Valid: true},
		ClassID:          userDomain.ClassID,
		ImageID:          userDomain.ImageID,
		Email:            sql.NullString{String: userDomain.Email, Valid: true},
		NISN:             sql.NullString{String: *userDomain.NISN, Valid: true},
		BirthCertificate: sql.NullString{String: userDomain.BirthCertificate, Valid: true},
		FamilyCard:       sql.NullString{String: userDomain.FamilyCard, Valid: true},
		Photo:            sql.NullString{String: userDomain.Photo, Valid: true},
		Roles:            sql.NullString{String: userDomain.Roles, Valid: true},
		Status:           sql.NullString{String: userDomain.Status, Valid: true},
		Sso:              sql.NullBool{Bool: userDomain.Sso, Valid: true},
		CreatedAt:        userDomain.CreatedAt,
		UpdatedAt:        userDomain.UpdatedAt,
	}
}
