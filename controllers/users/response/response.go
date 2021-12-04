package response

import (
	"go-schooling/business/users"
	classResp "go-schooling/controllers/classes/response"
	imageResp "go-schooling/controllers/images/response"
)

type Users struct {
	ID               int                `gorm:"primary_key" json:"id"`
	Name             string             `json:"name"`
	Classes          *classResp.Classes `json:"class"`
	Images           *imageResp.Images  `json:"images"`
	Email            string             `json:"email"`
	NISN             string             `json:"nisn"`
	BirthCertificate string             `json:"birth_certificate"`
	FamilyCard       string             `json:"family_card"`
	Photo            string             `json:"photo"`
	Roles            string             `json:"roles"`
}

func FromDomain(userDomain *users.Domain) (res *Users) {
	if userDomain != nil {
		res = &Users{
			ID:               userDomain.ID,
			Name:             userDomain.Name,
			Classes:          classResp.FromDomain(userDomain.Classes),
			Images:           imageResp.FromDomain(userDomain.Images),
			Email:            userDomain.Email,
			NISN:             userDomain.NISN,
			BirthCertificate: userDomain.BirthCertificate,
			FamilyCard:       userDomain.FamilyCard,
			Photo:            userDomain.Photo,
			Roles:            userDomain.Roles,
		}
	}
	return res
}
