package payments

import "context"

type Domain struct {
	ID       int
	UserID   int
	OrderID  string
	FullName string
	Email    string
	Amount   float64
}

type DomainResponse struct {
	Token         string
	RedirectURL   string
	StatusCode    string
	ErrorMessages []string
}

type Repository interface {
	Transactions(ctx context.Context, data *Domain) (DomainResponse, error)
	NotificationValidationKey() string
}
