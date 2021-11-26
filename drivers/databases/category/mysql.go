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

func (cr *mysqlCategoryRepository) Find(ctx context.Context, active string) ([]category.Domain, error) {
	rec := []Category{}

	query := cr.Conn.Where("archive = ?", false)

	if active != "" {
		query = query.Where("active = ?", active)
	}

	err := query.Find(&rec).Error
	if err != nil {
		return []category.Domain{}, err
	}

	categoryDomain := []category.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.ToDomain())
	}

	return categoryDomain, nil
}

func (cr *mysqlCategoryRepository) FindByID(id int) (category.Domain, error) {
	rec := Category{}

	if err := cr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return category.Domain{}, err
	}
	return rec.ToDomain(), nil
}