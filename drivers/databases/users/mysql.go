package users

import (
	"context"
	"go-schooling/business/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (repository *mysqlUsersRepository) GetAll(ctx context.Context) ([]users.Domain, error) {
	allUsers := []Users{}
	result := repository.Conn.Preload("Classes").Preload("Images").Find(&allUsers)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	var usrs []users.Domain
	for _, value := range allUsers {
		usrs = append(usrs, *value.toDomain())
	}
	return usrs, nil
}

func (repository *mysqlUsersRepository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	usersById := Users{}
	result := repository.Conn.Preload("Classes").Preload("Images").Where("users.id = ?", id).First(&usersById)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return *usersById.toDomain(), nil
}

func (repository *mysqlUsersRepository) Update(ctx context.Context, userDomain *users.Domain, id int) error {
	usersUpdate := fromDomain(*userDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&usersUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *mysqlUsersRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Preload("Classes").Preload("Images").Where("users.email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return *rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
