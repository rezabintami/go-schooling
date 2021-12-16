package request

import (
	"go-schooling/business/users"
)

type Users struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}
