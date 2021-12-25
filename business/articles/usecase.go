package articles

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/business/categoryarticles"
	"go-schooling/business/images"
	"strings"
	"time"
)

type ArticleUsecase struct {
	articleRepository          Repository
	categoryArticlesRepository categoryarticles.Repository
	imageUsecase               images.Usecase
	contextTimeout             time.Duration
	jwtAuth                    *middleware.ConfigJWT
}

func NewArticleUsecase(ur Repository, ca categoryarticles.Repository, iu images.Usecase, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ArticleUsecase{
		articleRepository:          ur,
		categoryArticlesRepository: ca,
		imageUsecase:               iu,
		jwtAuth:                    jwtauth,
		contextTimeout:             timeout,
	}
}

func (au *ArticleUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := au.articleRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (au *ArticleUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if id <= 0 {
		return Domain{}, business.ErrNewsIDResource
	}
	res, err := au.articleRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (au *ArticleUsecase) GetByTitle(ctx context.Context, title string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if strings.TrimSpace(title) == "" {
		return Domain{}, business.ErrNewsTitleResource
	}
	res, err := au.articleRepository.GetByTitle(ctx, title)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (au *ArticleUsecase) Store(ctx context.Context, articleDomain *Domain) error {
	for i := range articleDomain.Category {
		categoryart := categoryarticles.Domain{}
		categoryart.ArticleID = articleDomain.ID
		categoryart.CategoryID = articleDomain.Category[i]
		au.categoryArticlesRepository.Store(ctx, &categoryart)
	}

	err := au.articleRepository.Store(ctx, articleDomain)
	if err != nil {
		return err
	}

	return nil
}

func (au *ArticleUsecase) Update(ctx context.Context, articleDomain *Domain, id int) error {
	err := au.articleRepository.Update(ctx, articleDomain, id)
	if err != nil {
		return err
	}
	return nil
}
