package articles

import "time"

type Domain struct {
	Id           int
	Title        string
	Content      string
	CategoryID   int
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
