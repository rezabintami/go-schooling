package images

import (
	"context"
	"errors"
	"go-schooling/drivers/googlestorage"
	"go-schooling/helper/logging"
	"mime/multipart"
	"time"
)

type ImageUsecase struct {
	imageRepository Repository
	contextTimeout  time.Duration
	googlestorage   googlestorage.Connection
	logger          logging.Logger
}

func NewImageUsecase(ur Repository, googlestorage googlestorage.Connection, timeout time.Duration, logger logging.Logger) Usecase {
	return &ImageUsecase{
		imageRepository: ur,
		googlestorage:   googlestorage,
		contextTimeout:  timeout,
		logger:          logger,
	}
}

func (uc *ImageUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := uc.imageRepository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (tu *ImageUsecase) Store(ctx context.Context, imageDomain *Domain, file *multipart.FileHeader) (string, error) {
	err := tu.imageRepository.Store(ctx, imageDomain)
	if err != nil {
		return "", err
	}

	_, err = tu.googlestorage.Upload(imageDomain.Path, file)
	if err != nil {
		return "", errors.New("Unable to upload file: " + err.Error())
	}

	filePath, err := tu.googlestorage.GetPresignedUrl(imageDomain.Path)
	if err != nil {
		return "", errors.New("Unable to get url: " + err.Error())
	}
	return filePath, nil
}

func (tu *ImageUsecase) GetPresignedURL(ctx context.Context, name string) (string, error) {
	filePath, err := tu.googlestorage.GetPresignedUrl(name)
	if err != nil {
		return "", errors.New("Unable to get url: " + err.Error())
	}
	return filePath, nil
}
