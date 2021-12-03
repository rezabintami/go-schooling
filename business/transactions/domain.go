package transactions

import (
	"context"
	"go-schooling/business/payments"
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	OrderID     string
	FraudStatus string
	Name        string
	Amount      float64
	StatusCode  string
	SignKey     string
	Status      string
	PaymentUrl  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	CreateTransactions(ctx context.Context, data *Domain, id int) (payments.DomainResponse, error)
	Update(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (payments.Domain, error)
	Update(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByOrder(ctx context.Context, orderId string) (Domain, error)
}
