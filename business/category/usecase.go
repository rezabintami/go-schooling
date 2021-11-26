package category

import (
	"context"
	"go-schooling/business"
	"time"
)

type CategoryUsecase struct {
	categoryRepository Repository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(ur Repository, timeout time.Duration) Usecase {
	return &CategoryUsecase{
		categoryRepository: ur,
		contextTimeout:     timeout,
	}
}

func (cu *CategoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.categoryRepository.Find(ctx, "")
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
	findActive := "false"
	if active {
		findActive = "true"
	}
	resp, err := cu.categoryRepository.Find(ctx, findActive)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}