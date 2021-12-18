package teachers

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/helper/encrypt"
	"time"
)

type TeacherUsecase struct {
	teacherRepository Repository
	userRepository    interface{}
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewTeacherUsecase(tr Repository, ur interface{}, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &TeacherUsecase{
		teacherRepository: tr,
		userRepository:    ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
	}
}

func (tu *TeacherUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := tu.teacherRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil

}

func (tu *TeacherUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := tu.teacherRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (tu *TeacherUsecase) Update(ctx context.Context, teacherDomain *Domain, id int) error {
	err := tu.teacherRepository.Update(ctx, teacherDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TeacherUsecase) Store(ctx context.Context, teacherDomain *Domain) error {
	err := tu.teacherRepository.Store(ctx, teacherDomain)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TeacherUsecase) Login(ctx context.Context, email, password string) (string, error) {
	existedUser, err := tu.teacherRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) {
		return "", business.ErrEmailPasswordNotFound
	}

	token := tu.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	return token, nil
}
