package request

import "go-schooling/business/teachers"

type Teachers struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	NIP      string `json:"nip"`
	Photo    string `json:"photo"`
}

func (req *Teachers) ToDomain() *teachers.Domain {
	return &teachers.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		NIP:      req.NIP,
		Photo:    req.Photo,
	}
}
