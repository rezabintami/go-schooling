package category

import (
	"context"
	"go-schooling/business"
	"go-schooling/helper/logging"
	"time"
)

type CategoryUsecase struct {
	categoryRepository Repository
	contextTimeout     time.Duration
	logger             logging.Logger
}

func NewCategoryUsecase(ur Repository, timeout time.Duration, logger logging.Logger) Usecase {
	return &CategoryUsecase{
		categoryRepository: ur,
		contextTimeout:     timeout,
		logger:             logger,
	}
}

func (cu *CategoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.categoryRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *CategoryUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	if id <= 0 {
		return Domain{}, business.ErrIDNotFound
	}

	resp, err := cu.categoryRepository.FindByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (cu *CategoryUsecase) GetByActive(ctx context.Context, active bool) ([]Domain, error) {
	findActive := false
	if active {
		findActive = true
	}

	resp, err := cu.categoryRepository.Find(ctx, findActive)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (cu *CategoryUsecase) Store(ctx context.Context, categoryDomain *Domain) error {
	err := cu.categoryRepository.Store(ctx, categoryDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CategoryUsecase) Delete(ctx context.Context, id int) error {
	err := cu.categoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
