package response

import "go-schooling/business/users"

type Users struct {
	ID               int    `gorm:"primary_key" json:"id"`
	Name             string `json:"name"`
	Password         string `json:"-"`
	Email            string `json:"email"`
	UUID             string `json:"uuid"`
	BirthCertificate string `json:"birth_certificate"`
	FamilyCard       string `json:"family_card"`
	Photo            string `json:"photo"`
	Roles            string `json:"roles"`
	Sso              bool   `json:"-"`
}

func FromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:               userDomain.ID,
		Name:             userDomain.Name,
		Password:         userDomain.Password,
		Email:            userDomain.Email,
		UUID:             userDomain.UUID,
		BirthCertificate: userDomain.BirthCertificate,
		FamilyCard:       userDomain.FamilyCard,
		Photo:            userDomain.Photo,
		Roles:            userDomain.Roles,
		Sso:              userDomain.Sso,
	}
}
