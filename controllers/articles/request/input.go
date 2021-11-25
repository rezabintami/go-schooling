package request

import "go-schooling/business/articles"

type Articles struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
	ImageID    int    `json:"image_id"`
}

func (req *Articles) ToDomain() *articles.Domain {
	return &articles.Domain{
		Title: 	req.Title,
		Content: req.Content,
		CategoryID: req.CategoryID,
		ImageID: req.ImageID,
	}
}
