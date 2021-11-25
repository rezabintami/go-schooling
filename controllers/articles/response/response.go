package response

import (
	"go-schooling/business/articles"
	"time"
)

type Articles struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID int       `json:"category_id"`
	ImageID    int       `json:"image_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ArticlesResponse struct {
	Articles *[]Articles `json:"articles"`
	Page     int         `json:"page"`
}

func FromDomain(articleDomain articles.Domain) *Articles {
	return &Articles{
		ID:         articleDomain.ID,
		Title:      articleDomain.Title,
		Content:    articleDomain.Content,
		CategoryID: articleDomain.CategoryID,
		ImageID:    articleDomain.ImageID,
		CreatedAt:  articleDomain.CreatedAt,
		UpdatedAt:  articleDomain.UpdatedAt,
	}
}

func FromListDomain(articlesDomain []articles.Domain) *[]Articles {
	articles := []Articles{}
	for _, value := range articlesDomain {
		article := Articles{
			ID:         value.ID,
			Title:      value.Title,
			Content:    value.Content,
			CategoryID: value.CategoryID,
			ImageID:    value.ImageID,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
		}
		articles = append(articles, article)
	}
	return &articles
}

func FromListPageDomain(articlesDomain []articles.Domain, Page int) *ArticlesResponse {
	articles := []Articles{}
	for _, value := range articlesDomain {
		article := Articles{
			ID:         value.ID,
			Title:      value.Title,
			Content:    value.Content,
			CategoryID: value.CategoryID,
			ImageID:    value.ImageID,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
		}
		articles = append(articles, article)
	}
	articlesResponse := ArticlesResponse{}
	articlesResponse.Articles = &articles
	articlesResponse.Page = Page
	return &articlesResponse
}
