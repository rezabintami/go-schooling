package teachers

import (
	"go-schooling/business/teachers"
	"time"
)

type Teachers struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	NIP       string    `json:"nip"`
	Photo     string    `json:"photo"`
	Roles     string    `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Teachers) toDomain() teachers.Domain {
	return teachers.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		NIP:       rec.NIP,
		Photo:     rec.Photo,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain teachers.Domain) *Teachers {
	return &Teachers{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Email:     userDomain.Email,
		NIP:       userDomain.NIP,
		Photo:     userDomain.Photo,
		Roles:     userDomain.Roles,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
