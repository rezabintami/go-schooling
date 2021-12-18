package classes

import (
	"database/sql"
	"go-schooling/business/classes"
	"go-schooling/drivers/databases/teachers"
	"time"
)

type Classes struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	TeacherID sql.NullInt64
	Teachers  *teachers.Teachers `gorm:"foreignKey:TeacherID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Classes) ToDomain() (res *classes.Domain) {
	if rec != nil {
		res = &classes.Domain{
			ID:        rec.ID,
			Name:      rec.Name,
			Teachers:  rec.Teachers.ToDomain(),
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}

func fromDomain(classDomain classes.Domain) *Classes {
	return &Classes{
		ID:        classDomain.ID,
		Name:      classDomain.Name,
		CreatedAt: classDomain.CreatedAt,
		UpdatedAt: classDomain.UpdatedAt,
	}
}
