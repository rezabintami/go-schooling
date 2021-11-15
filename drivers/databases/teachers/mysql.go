package teachers

import (
	"go-schooling/business/teachers"

	"gorm.io/gorm"
)

type mysqlTeachersRepository struct {
	Conn *gorm.DB
}

func NewMySQLTeachersRepository(conn *gorm.DB) teachers.Repository {
	return &mysqlTeachersRepository{
		Conn: conn,
	}
}