package teachers

import (
	"context"
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

func (repository *mysqlTeachersRepository) GetByID(ctx context.Context, id int) (teachers.Domain, error) {
	teacherById := Teachers{}
	result := repository.Conn.Where("id = ?", id).First(&teacherById)
	if result.Error != nil {
		return teachers.Domain{}, result.Error
	}
	return teacherById.toDomain(), nil
}

func (repository *mysqlTeachersRepository) GetAll(ctx context.Context) ([]teachers.Domain, error) {
	getTeachers := []Teachers{}
	result := repository.Conn.Find(&getTeachers)
	if result.Error != nil {
		return []teachers.Domain{}, result.Error
	}
	var allTeachers []teachers.Domain
	for _, value := range getTeachers {
		allTeachers = append(allTeachers, value.toDomain())
	}
	return allTeachers, nil
}

func (repository *mysqlTeachersRepository) Update(ctx context.Context, teacherDomain *teachers.Domain, id int) error {
	teacherUpdate := fromDomain(*teacherDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&teacherUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTeachersRepository) Store(ctx context.Context, teacherDomain *teachers.Domain) error {
	rec := fromDomain(*teacherDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}