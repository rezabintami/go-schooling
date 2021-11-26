package images

import (
	"context"
	"time"
)

type ImageUsecase struct {
	imageRepository Repository
	contextTimeout  time.Duration
}

func NewImageUsecase(ur Repository, timeout time.Duration) Usecase {
	return &ImageUsecase{
		imageRepository: ur,
		contextTimeout:  timeout,
	}
}

func (uc *ImageUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := uc.imageRepository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (tu *ImageUsecase) Store(ctx context.Context, imageDomain *Domain) error {
	err := tu.imageRepository.Store(ctx, imageDomain)
	if err != nil {
		return err
	}

	return nil
}