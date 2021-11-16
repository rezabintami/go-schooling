package classes

import (
	"context"
	"go-schooling/business/classes"

	"gorm.io/gorm"
)

type mysqlClassesRepository struct {
	Conn *gorm.DB
}

func NewMySQLClassesRepository(conn *gorm.DB) classes.Repository {
	return &mysqlClassesRepository{
		Conn: conn,
	}
}

func (repository *mysqlClassesRepository) GetAll(ctx context.Context) ([]classes.Domain, error) {
	getClasses := []Classes{}
	result := repository.Conn.Find(&getClasses)
	if result.Error != nil {
		return []classes.Domain{}, result.Error
	}
	var allClasses []classes.Domain
	for _, value := range getClasses {
		allClasses = append(allClasses, value.toDomain())
	}
	return allClasses, nil
}

func (repository *mysqlClassesRepository) Store(ctx context.Context, classDomain *classes.Domain) error {
	rec := fromDomain(*classDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlClassesRepository) Delete(ctx context.Context, id int) error {
	classDelete := Classes{}
	result := repository.Conn.Where("id = ?", id).Delete(&classDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}