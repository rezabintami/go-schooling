package users

import (
	"go-schooling/business/users"
	"time"
)

type Users struct {
	ID               int       `gorm:"primary_key" json:"id"`
	Name             string    `json:"name"`
	Password         string    `json:"-"`
	Email            string    `json:"email"`
	NISN             string    `json:"nisn"`
	BirthCertificate string    `json:"birth_certificate"`
	FamilyCard       string    `json:"family_card"`
	Photo            string    `json:"photo"`
	Roles            string    `json:"roles"`
	Sso              bool      `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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
