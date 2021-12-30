package teachers

import (
	"go-schooling/business/teachers"
	"time"
)

type Teachers struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Password  string
	Email     string
	NIP       string
	Photo     string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Teachers) ToDomain() (res *teachers.Domain) {
	if rec != nil {
		res = &teachers.Domain{
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
	return res
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
