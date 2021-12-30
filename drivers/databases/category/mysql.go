package category

import (
	"context"
	"go-schooling/business/category"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMySQLCategoryRepository(conn *gorm.DB) category.Repository {
	return &mysqlCategoryRepository{
		Conn: conn,
	}
}

func (repository *mysqlCategoryRepository) Find(ctx context.Context, active bool) ([]category.Domain, error) {
	rec := []Category{}

	query := repository.Conn.Where("active = ?", active)

	err := query.Find(&rec).Error
	if err != nil {
		return []category.Domain{}, err
	}

	categoryDomain := []category.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, *value.ToDomain())
	}

	return categoryDomain, nil
}

func (repository *mysqlCategoryRepository) GetAll(ctx context.Context) ([]category.Domain, error) {
	rec := []Category{}
	result := repository.Conn.Find(&rec)
	if result.Error != nil {
		return []category.Domain{}, result.Error
	}
	categoryDomain := []category.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, *value.ToDomain())
	}

	return categoryDomain, nil
}

func (repository *mysqlCategoryRepository) FindByID(id int) (category.Domain, error) {
	rec := Category{}

	if err := repository.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return category.Domain{}, err
	}
	return *rec.ToDomain(), nil
}

func (repository *mysqlCategoryRepository) Store(ctx context.Context, classDomain *category.Domain) error {
	rec := fromDomain(*classDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlCategoryRepository) Delete(ctx context.Context, id int) error {
	classDelete := Category{}
	result := repository.Conn.Where("id = ?", id).Delete(&classDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
