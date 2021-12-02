package images

import (
	"context"
	"errors"
	"go-schooling/drivers/googlestorage"
	"mime/multipart"
	"time"
)

type ImageUsecase struct {
	imageRepository Repository
	contextTimeout  time.Duration
	googlestorage   googlestorage.Connection
}

func NewImageUsecase(ur Repository, googlestorage googlestorage.Connection, timeout time.Duration) Usecase {
	return &ImageUsecase{
		imageRepository: ur,
		googlestorage:   googlestorage,
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

func (tu *ImageUsecase) Store(ctx context.Context, imageDomain *Domain, file *multipart.FileHeader) (string, error) {
	err := tu.imageRepository.Store(ctx, imageDomain)
	if err != nil {
		return "", err
	}

	_, err = tu.googlestorage.Upload(imageDomain.Path, imageDomain.Name, file)
	if err != nil {
		return "", errors.New("Unable to upload file: " + err.Error())
	}

	filePath, err := tu.googlestorage.GetPresignedUrl(imageDomain.Name)
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
