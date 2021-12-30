package request

import (
	"go-schooling/business/users"
)

type Users struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

type UpdateUsers struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	BirthCertificate string `json:"birth_certificate"`
	FamilyCard       string `json:"family_card"`
	Photo            string `json:"photo"`
}

type UpdateAdminUsers struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	NISN             string `json:"nisn"`
	BirthCertificate string `json:"birth_certificate"`
	FamilyCard       string `json:"family_card"`
	Photo            string `json:"photo"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}

func (req *UpdateUsers) ToUpdateDomain() *users.Domain {
	return &users.Domain{
		Name:             req.Name,
		Email:            req.Email,
		BirthCertificate: &req.BirthCertificate,
		FamilyCard:       &req.FamilyCard,
		Photo:            &req.Photo,
	}
}
