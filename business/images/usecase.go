package images

import (
	"context"
	"go-schooling/drivers/googlestorage"
	"time"

	"github.com/JoinVerse/xid"
)

type ImageUsecase struct {
	imageRepository Repository
	contextTimeout  time.Duration
	googlestorage   googlestorage.Connection
}

func NewImageUsecase(ur Repository, googlestorage googlestorage.Connection, timeout time.Duration) Usecase {
	return &ImageUsecase{
		imageRepository: ur,
		googlestorage: googlestorage,
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
	filePath := "articles-" + xid.New().String()
	imageDomain.Path = filePath
	err := tu.imageRepository.Store(ctx, imageDomain)
	if err != nil {
		return err
	}

	tu.googlestorage.Upload(imageDomain.Name, filePath)

	return nil
}
