package articles

import (
	"context"
	"go-schooling/app/middleware"
	"go-schooling/business"
	"go-schooling/business/category"
	"go-schooling/business/categoryarticles"
	"go-schooling/business/images"
	"go-schooling/helper/logging"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ArticleUsecase struct {
	articleRepository          Repository
	categoryArticlesRepository categoryarticles.Repository
	categoryRepository         category.Repository
	imageUsecase               images.Usecase
	contextTimeout             time.Duration
	jwtAuth                    *middleware.ConfigJWT
	logger                     logging.Logger
}

func NewArticleUsecase(ur Repository, ca categoryarticles.Repository, cr category.Repository, iu images.Usecase, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &ArticleUsecase{
		articleRepository:          ur,
		categoryArticlesRepository: ca,
		categoryRepository:         cr,
		imageUsecase:               iu,
		jwtAuth:                    jwtauth,
		contextTimeout:             timeout,
		logger:                     logger,
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
		return DomainFromArticles{}, business.ErrArticleIDResource
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
		return DomainFromArticles{}, business.ErrArticleTitleResource
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

func (au *ArticleUsecase) GetByCategory(ctx context.Context, category []string) ([]DomainFromArticles, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	var allListArticle []DomainFromArticles
	var ListArticle []DomainFromArticles

	for iter := range category {
		categoryID, _ := strconv.Atoi(category[iter])
		categoryarticles, err := au.categoryArticlesRepository.GetAllByCategoryID(ctx, categoryID)

		if err != nil {
			return []DomainFromArticles{}, err
		}

		for _, value := range categoryarticles {
			var ListCategory []string

			article, err := au.articleRepository.GetByID(ctx, value.ArticleID)
			if err != nil {
				return []DomainFromArticles{}, err
			}

			category, err := au.categoryArticlesRepository.GetAllByArticleID(ctx, article.ID)
			if err != nil {
				return []DomainFromArticles{}, err
			}

			for _, value := range category {
				ListCategory = append(ListCategory, value.Category.Title)
			}

			article.Category = ListCategory
			ListArticle = append(ListArticle, article)
		}
		if iter == 0 {
			allListArticle = append(allListArticle, ListArticle...)
		} else {
			for _, value := range ListArticle {
				var isExist bool
				for _, value2 := range allListArticle {
					if value.ID == value2.ID {
						isExist = true
					}
				}
				if !isExist {
					allListArticle = append(allListArticle, value)
				}
			}
		}
	}
	sort.Slice(allListArticle, func(i, j int) bool {
		return allListArticle[i].ID < allListArticle[j].ID
	})

	return allListArticle, nil
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
	err := au.categoryArticlesRepository.DeleteByArticleID(ctx, id)
	if err != nil {
		return err
	}

	err = au.articleRepository.Update(ctx, articleDomain, id)
	if err != nil {
		return err
	}

	for i := range articleDomain.Category {
		categoryart := categoryarticles.Domain{}
		categoryart.ArticleID = id
		categoryart.CategoryID = articleDomain.Category[i]
		au.categoryArticlesRepository.Store(ctx, &categoryart)
	}

	return nil
}
