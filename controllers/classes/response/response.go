package response

import "go-schooling/business/classes"

type Classes struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func FromListDomain(classDomain []classes.Domain) *[]Classes {
	classes := []Classes{}
	for _, value := range classDomain {
		class := Classes{
			ID:   value.ID,
			Name: value.Name,
		}
		classes = append(classes, class)
	}
	return &classes
}