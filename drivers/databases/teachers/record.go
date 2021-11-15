package teachers

import "time"

type Teachers struct {
	ID               int       `gorm:"primary_key" json:"id"`
	Name             string    `json:"name"`
	Password         string    `json:"-"`
	Email            string    `json:"email"`
	NIP             string    `json:"nip"`
	Photo            string    `json:"photo"`
	Roles            string    `json:"roles"`
	Sso              bool      `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}