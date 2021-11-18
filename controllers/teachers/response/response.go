package response

import "go-schooling/business/teachers"

type Teachers struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	NIP      string `json:"nip"`
	Photo    string `json:"photo"`
	Roles    string `json:"roles"`
}

func FromDomain(teacherDomain teachers.Domain) *Teachers {
	return &Teachers{
		ID:       teacherDomain.ID,
		Name:     teacherDomain.Name,
		Password: teacherDomain.Password,
		Email:    teacherDomain.Email,
		NIP:      teacherDomain.NIP,
		Photo:    teacherDomain.Photo,
		Roles:    teacherDomain.Roles,
	}
}

func FromListDomain(teacherDomain []teachers.Domain) *[]Teachers {
	teachers := []Teachers{}
	for _, value := range teacherDomain {
		teacher := Teachers{
			ID:       value.ID,
			Name:     value.Name,
			Password: value.Password,
			Email:    value.Email,
			NIP:      value.NIP,
			Photo:    value.Photo,
			Roles:    value.Roles,
		}
		teachers = append(teachers, teacher)
	}
	return &teachers
}
