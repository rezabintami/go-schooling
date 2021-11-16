package request

import "go-schooling/business/classes"

type Classes struct {
	Name string `json:"name"`
}

func (req *Classes) ToDomain() *classes.Domain {
	return &classes.Domain{
		Name: req.Name,
	}
}
