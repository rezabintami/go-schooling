package response

import (
	"go-schooling/business/users"
	classResp "go-schooling/controllers/classes/response"
)

type Users struct {
	ID               int                `gorm:"primary_key" json:"id"`
	Name             string             `json:"name"`
	Classes          *classResp.Classes `json:"class"`
	Email            string             `json:"email"`
	NISN             *string            `json:"nisn"`
	BirthCertificate *string            `json:"birth_certificate"`
	FamilyCard       *string            `json:"family_card"`
	Photo            *string             `json:"photo"`
	Roles            string             `json:"roles"`
}

type UsersPageResponse struct {
	Users *[]Users `json:"users"`
	Total  int      `json:"total"`
}

func FromDomain(userDomain users.Domain) Users {
	return Users{
		ID:               userDomain.ID,
		Name:             userDomain.Name,
		Classes:          classResp.FromDomain(userDomain.Classes),
		Email:            userDomain.Email,
		NISN:             userDomain.NISN,
		BirthCertificate: userDomain.BirthCertificate,
		FamilyCard:       userDomain.FamilyCard,
		Photo:            userDomain.Photo,
		Roles:            userDomain.Roles,
	}
}

func FromListDomain(userDomain []users.Domain) *[]Users {
	allUsers := []Users{}
	for _, value := range userDomain {
		user := Users{
			ID:               value.ID,
			Name:             value.Name,
			Classes:          classResp.FromDomain(value.Classes),
			Email:            value.Email,
			NISN:             value.NISN,
			BirthCertificate: value.BirthCertificate,
			FamilyCard:       value.FamilyCard,
			Photo:            value.Photo,
			Roles:            value.Roles,
		}
		allUsers = append(allUsers, user)
	}
	return &allUsers
}

func FromListPageDomain(userDomain []users.Domain, Count int) *UsersPageResponse {
	allUsers := []Users{}
	for _, value := range userDomain {
		user := Users{
			ID:               value.ID,
			Name:             value.Name,
			Classes:          classResp.FromDomain(value.Classes),
			Email:            value.Email,
			NISN:             value.NISN,
			BirthCertificate: value.BirthCertificate,
			FamilyCard:       value.FamilyCard,
			Photo:            value.Photo,
			Roles:            value.Roles,
		}
		allUsers = append(allUsers, user)
	}
	articlesResponse := UsersPageResponse{}
	articlesResponse.Users = &allUsers
	articlesResponse.Total = Count
	return &articlesResponse
}
