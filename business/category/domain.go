package category

import "time"

type Domain struct {
	ID          int
	Title       string
	Description string
	Active      bool
	Archive     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}