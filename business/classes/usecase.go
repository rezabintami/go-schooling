package classes

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/helper/logging"
	"time"
)

type ClassUsecase struct {
	classRepository Repository
	jwtAuth         *middleware.ConfigJWT
	contextTimeout  time.Duration
	logger          logging.Logger
}

func NewClassUsecase(cr Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &ClassUsecase{
		classRepository: cr,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
		logger:          logger,
	}
}

func (cu *ClassUsecase) Delete(ctx context.Context, id int) error {
	err := cu.classRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (cu *ClassUsecase) Store(ctx context.Context, teacherDomain *Domain) error {
	err := cu.classRepository.Store(ctx, teacherDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cu *ClassUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := cu.classRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (cu *ClassUsecase) GetByID(ctx context.Context, id int) (*Domain, error) {
	result, err := cu.classRepository.GetByID(ctx, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
