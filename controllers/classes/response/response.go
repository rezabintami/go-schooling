package response

import (
	"go-schooling/business/classes"
	teachersResp "go-schooling/controllers/teachers/response"
)

type Classes struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	TeachersName string `json:"teacher"`
}

func FromListDomain(classDomain []classes.Domain) *[]Classes {
	classes := []Classes{}
	for _, value := range classDomain {
		class := Classes{
			ID:           value.ID,
			Name:         value.Name,
			TeachersName: teachersResp.FromDomain(value.Teachers).Name,
		}
		classes = append(classes, class)
	}
	return &classes
}

func FromDomain(domain *classes.Domain) (res *Classes) {
	if domain != nil {
		res = &Classes{
			ID:           domain.ID,
			Name:         domain.Name,
			TeachersName: teachersResp.FromDomain(domain.Teachers).Name,
		}
	}
	return res
}
