package request

import "go-schooling/business/articles"

type Articles struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	ListCategory []int  `json:"category"`
}

func (req *Articles) ToDomain() *articles.Domain {
	return &articles.Domain{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.ListCategory,
	}
}
