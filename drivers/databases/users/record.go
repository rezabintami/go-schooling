package users

import (
	"go-schooling/business/users"
	"go-schooling/drivers/databases/classes"
	"go-schooling/drivers/databases/images"
	"time"
)

type Users struct {
	ID               int `gorm:"primary_key" json:"id"`
	Name             string
	Password         string
	ClassID          int
	Classes          classes.Classes `gorm:"foreignKey:ClassID;references:ID"`
	ImageID          int
	Images           images.Images `gorm:"foreignKey:ImageID;references:ID"`
	Email            string
	NISN             string
	BirthCertificate string
	FamilyCard       string
	Photo            string
	Roles            string
	Sso              bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:               rec.ID,
		Name:             rec.Name,
		Password:         rec.Password,
		Email:            rec.Email,
		NISN:             rec.NISN,
		BirthCertificate: rec.BirthCertificate,
		FamilyCard:       rec.FamilyCard,
		Photo:            rec.Photo,
		Roles:            rec.Roles,
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
		Email:            userDomain.Email,
		NISN:             userDomain.NISN,
		BirthCertificate: userDomain.BirthCertificate,
		FamilyCard:       userDomain.FamilyCard,
		Photo:            userDomain.Photo,
		Roles:            userDomain.Roles,
		Sso:              userDomain.Sso,
		CreatedAt:        userDomain.CreatedAt,
		UpdatedAt:        userDomain.UpdatedAt,
	}
}
