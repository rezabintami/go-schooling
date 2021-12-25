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

func (au *ArticleUsecase) Fetch(ctx context.Context, page, perpage int) ([]DomainFromArticles, int, error) {
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
		return []DomainFromArticles{}, 0, err
	}

	for j, value := range res {
		var ListCategory []string
		category, err := au.categoryArticlesRepository.GetAllByArticleID(ctx, value.ID)
		if err != nil {
			return []DomainFromArticles{}, 0, err
		}
		for _, value := range category {
			ListCategory = append(ListCategory, value.Category.Title)
		}
		res[j].Category = ListCategory
	}

	return res, total, nil
}

func (au *ArticleUsecase) GetByID(ctx context.Context, id int) (DomainFromArticles, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if id <= 0 {
		return DomainFromArticles{}, business.ErrNewsIDResource
	}
	res, err := au.articleRepository.GetByID(ctx, id)
	if err != nil {
		return DomainFromArticles{}, err
	}

	var ListCategory []string
	category, err := au.categoryArticlesRepository.GetAllByArticleID(ctx, res.ID)
	if err != nil {
		return DomainFromArticles{}, err
	}
	for _, value := range category {
		ListCategory = append(ListCategory, value.Category.Title)
	}
	res.Category = ListCategory

	return res, nil
}

func (au *ArticleUsecase) GetByTitle(ctx context.Context, title string) (DomainFromArticles, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	if strings.TrimSpace(title) == "" {
		return DomainFromArticles{}, business.ErrNewsTitleResource
	}
	res, err := au.articleRepository.GetByTitle(ctx, title)
	if err != nil {
		return DomainFromArticles{}, err
	}

	var ListCategory []string
	category, err := au.categoryArticlesRepository.GetAllByArticleID(ctx, res.ID)
	if err != nil {
		return DomainFromArticles{}, err
	}
	for _, value := range category {
		ListCategory = append(ListCategory, value.Category.Title)
	}
	res.Category = ListCategory

	return res, nil
}

func (au *ArticleUsecase) Store(ctx context.Context, articleDomain *Domain) error {
	articleID, err := au.articleRepository.Store(ctx, articleDomain)
	if err != nil {
		return err
	}

	for i := range articleDomain.Category {
		categoryart := categoryarticles.Domain{}
		categoryart.ArticleID = articleID
		categoryart.CategoryID = articleDomain.Category[i]
		au.categoryArticlesRepository.Store(ctx, &categoryart)
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
