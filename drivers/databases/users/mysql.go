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
	result := repository.Conn.Preload("Classes").Preload("Classes.Teachers").Find(&allUsers)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	var usrs []users.Domain
	for _, value := range allUsers {
		usrs = append(usrs, *value.ToDomain())
	}
	return usrs, nil
}

func (repository *mysqlUsersRepository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	usersById := Users{}
	result := repository.Conn.Preload("Classes").Preload("Classes.Teachers").Where("users.id = ?", id).First(&usersById)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return *usersById.ToDomain(), nil
}

func (repository *mysqlUsersRepository) Update(ctx context.Context, userDomain *users.Domain, id int) error {
	usersUpdate := fromDomain(*userDomain)

	result := repository.Conn.Preload("Classes").Preload("Classes.Teachers").Where("users.id = ?", id).Updates(&usersUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlUsersRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}

	err := repository.Conn.Preload("Classes").Preload("Classes.Teachers").Where("users.email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	
	return *rec.ToDomain(), nil
}

func (repository *mysqlUsersRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	rec := fromRegisterDomain(*userDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlUsersRepository) Fetch(ctx context.Context, page, perpage int) ([]users.Domain, int, error) {
	rec := []Users{}

	offset := (page - 1) * perpage
	err := repository.Conn.Preload("Classes").Preload("Classes.Teachers").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var domainUsers []users.Domain
	for _, value := range rec {
		domainUsers = append(domainUsers, *value.ToDomain())
	}

	return domainUsers, int(totalData), nil
}
