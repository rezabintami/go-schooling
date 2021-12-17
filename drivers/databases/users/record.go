package users

import (
	"database/sql"
	"go-schooling/business/users"
	"go-schooling/drivers/databases/classes"
	"go-schooling/helper/convertpointer"
	"time"
)

type Users struct {
	ID               int `gorm:"primary_key" json:"id"`
	Name             sql.NullString
	Password         sql.NullString
	ClassID          sql.NullInt64
	Classes          *classes.Classes `gorm:"foreignKey:ClassID;references:ID"`
	Email            sql.NullString
	NISN             sql.NullString
	BirthCertificate sql.NullString
	FamilyCard       sql.NullString
	Photo            sql.NullString
	Roles            sql.NullString
	Status           sql.NullString
	Graduated        sql.NullBool
	Sso              sql.NullBool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (rec *Users) ToDomain() *users.Domain {
	return &users.Domain{
		ID:               rec.ID,
		Name:             rec.Name.String,
		Password:         rec.Password.String,
		Classes:          rec.Classes.ToDomain(),
		Email:            rec.Email.String,
		NISN:             convertpointer.ConvertPointerString(&rec.NISN.String),
		BirthCertificate: convertpointer.ConvertPointerString(&rec.BirthCertificate.String),
		FamilyCard:       convertpointer.ConvertPointerString(&rec.FamilyCard.String),
		Photo:            convertpointer.ConvertPointerString(&rec.Photo.String),
		Roles:            rec.Roles.String,
		Status:           rec.Status.String,
		Graduated:        rec.Graduated.Bool,
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
		Email:            sql.NullString{String: userDomain.Email, Valid: true},
		NISN:             sql.NullString{String: *userDomain.NISN, Valid: true},
		BirthCertificate: sql.NullString{String: *userDomain.BirthCertificate, Valid: true},
		FamilyCard:       sql.NullString{String: *userDomain.FamilyCard, Valid: true},
		Photo:            sql.NullString{String: *userDomain.Photo, Valid: true},
		Roles:            sql.NullString{String: userDomain.Roles, Valid: true},
		Status:           sql.NullString{String: userDomain.Status, Valid: true},
		Graduated:        sql.NullBool{Bool: userDomain.Graduated, Valid: true},
		Sso:              sql.NullBool{Bool: userDomain.Sso, Valid: true},
		CreatedAt:        userDomain.CreatedAt,
		UpdatedAt:        userDomain.UpdatedAt,
	}
}
