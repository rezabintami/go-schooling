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
	Name             string
	Password         string
	ClassID          sql.NullInt64
	Classes          *classes.Classes `gorm:"foreignKey:ClassID;references:ID"`
	Email            string
	NISN             string
	BirthCertificate string
	FamilyCard       string
	Photo            string
	Roles            string
	Status           string
	Graduated        bool
	Sso              bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (rec *Users) ToDomain() *users.Domain {
	return &users.Domain{
		ID:               rec.ID,
		Name:             rec.Name,
		Password:         rec.Password,
		Classes:          rec.Classes.ToDomain(),
		Email:            rec.Email,
		NISN:             convertpointer.ConvertPointerString(&rec.NISN),
		BirthCertificate: convertpointer.ConvertPointerString(&rec.BirthCertificate),
		FamilyCard:       convertpointer.ConvertPointerString(&rec.FamilyCard),
		Photo:            convertpointer.ConvertPointerString(&rec.Photo),
		Roles:            rec.Roles,
		Status:           rec.Status,
		Graduated:        rec.Graduated,
		Sso:              rec.Sso,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:               userDomain.ID,
		Name:             userDomain.Name,
		Password:         userDomain.Password,
		ClassID:          userDomain.ClassID,
		Email:            userDomain.Email,
		NISN:             convertpointer.ConvertNilPointerString(userDomain.NISN),
		BirthCertificate: convertpointer.ConvertNilPointerString(userDomain.BirthCertificate),
		FamilyCard:       convertpointer.ConvertNilPointerString(userDomain.FamilyCard),
		Photo:            convertpointer.ConvertNilPointerString(userDomain.Photo),
		Roles:            userDomain.Roles,
		Status:           userDomain.Status,
		Graduated:        userDomain.Graduated,
		Sso:              userDomain.Sso,
		CreatedAt:        userDomain.CreatedAt,
		UpdatedAt:        userDomain.UpdatedAt,
	}
}

func fromRegisterDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		ClassID:   userDomain.ClassID,
		Email:     userDomain.Email,
		Roles:     userDomain.Roles,
		Status:    userDomain.Status,
		Graduated: userDomain.Graduated,
		Sso:       userDomain.Sso,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
