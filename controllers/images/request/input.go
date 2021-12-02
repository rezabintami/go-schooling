package request

import "go-schooling/business/images"

type Images struct {
	Name string
	Path string
}

func (req *Images) ToDomain() *images.Domain {
	return &images.Domain{
		Name: req.Name,
		Path: req.Path,
	}
}
