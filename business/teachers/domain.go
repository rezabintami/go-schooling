package teachers

import "time"

type Domain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	NIP       string
	Photo     string
	Roles     string
	Sso       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
}

type Repository interface {
}