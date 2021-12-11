package request

import "go-schooling/business/category"

type Category struct {
	Title  string `json:"title"`
	Active bool   `json:"status"`
}

func (req *Category) ToDomain() *category.Domain {
	return &category.Domain{
		Title: req.Title,
		Active: req.Active,
	}
}
