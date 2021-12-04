package response

import (
	"go-schooling/business/images"
)

type Images struct {
	Path string `json:"path"`
}

func FromDomain(domain *images.Domain) (res *Images) {
	if domain != nil {
		res = &Images{
			Path: domain.Path,
		}
	}
	return res
}
