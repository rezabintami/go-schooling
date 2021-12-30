package response

import (
	"go-schooling/business/articles"
	"time"
)

type Articles struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Catagory  []string  `json:"catagory"`
	ImageID   int       `json:"image_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticlesResponse struct {
	Articles *[]Articles `json:"articles"`
	Count    int         `json:"count"`
}

func FromDomain(articleDomain articles.DomainFromArticles) *Articles {
	return &Articles{
		ID:        articleDomain.ID,
		Title:     articleDomain.Title,
		Content:   articleDomain.Content,
		Catagory:  articleDomain.Category,
		CreatedAt: articleDomain.CreatedAt,
		UpdatedAt: articleDomain.UpdatedAt,
	}
}

func FromListDomain(articlesDomain []articles.DomainFromArticles) *[]Articles {
	articles := []Articles{}
	for _, value := range articlesDomain {
		article := Articles{
			ID:        value.ID,
			Title:     value.Title,
			Content:   value.Content,
			Catagory:  value.Category,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		articles = append(articles, article)
	}
	return &articles
}

func FromListPageDomain(articlesDomain []articles.DomainFromArticles, Count int) *ArticlesResponse {
	articles := []Articles{}
	for _, value := range articlesDomain {
		article := Articles{
			ID:        value.ID,
			Title:     value.Title,
			Content:   value.Content,
			Catagory:  value.Category,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		articles = append(articles, article)
	}
	articlesResponse := ArticlesResponse{}
	articlesResponse.Articles = &articles
	articlesResponse.Count = Count
	return &articlesResponse
}
