package users

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/helper/encrypt"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	existedUser, err := uc.userRepository.GetByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := uc.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	return token, nil
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := uc.userRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (uc *UserUsecase) Update(ctx context.Context, userDomain *Domain, id int) error {
	err := uc.userRepository.Update(ctx, userDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) Register(ctx context.Context, userDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	existedUser, err := uc.userRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		userDomain.Password, _ = encrypt.Hash(userDomain.Password)
	}

	userDomain.Sso = sso
	userDomain.Roles = "USER"
	userDomain.Graduated = false
	err = uc.userRepository.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := uc.userRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (uc *UserUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
