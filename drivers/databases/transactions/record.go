package transactions

import (
	"go-schooling/business/payments"
	"go-schooling/business/transactions"
	"go-schooling/drivers/databases/users"
	"time"
)

type Transactions struct {
	ID         int         `gorm:"primary_key" json:"id"`
	OrderID    string      `json:"order_id"`
	UserID     int         `json:"user_id"`
	User       users.Users `gorm:"foreignKey:UserID;references:ID"`
	Amount     float64     `json:"amount"`
	Status     string      `json:"status"`
	PaymentUrl string      `json:"payment_url"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func fromDomain(topupDomain transactions.Domain) *Transactions {
	return &Transactions{
		ID:         topupDomain.ID,
		UserID:     topupDomain.UserID,
		OrderID:    topupDomain.OrderID,
		Amount:     topupDomain.Amount,
		Status:     topupDomain.Status,
		PaymentUrl: topupDomain.PaymentUrl,
		CreatedAt:  topupDomain.CreatedAt,
		UpdatedAt:  topupDomain.UpdatedAt,
	}
}

func (rec *Transactions) toDomain() transactions.Domain {
	return transactions.Domain{
		ID:         rec.ID,
		UserID:     rec.UserID,
		OrderID:    rec.OrderID,
		Name:       rec.User.Name.String,
		Amount:     rec.Amount,
		Status:     rec.Status,
		PaymentUrl: rec.PaymentUrl,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func (rec *Transactions) toPaymentDomain() payments.Domain {
	return payments.Domain{
		ID:       rec.ID,
		UserID:   rec.UserID,
		OrderID:  rec.OrderID,
		FullName: rec.User.Name.String,
		Email:    rec.User.Email.String,
		Amount:   rec.Amount,
	}
}
