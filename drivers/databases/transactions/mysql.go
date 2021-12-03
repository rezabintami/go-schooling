package transactions

import (
	"context"
	"go-schooling/business/payments"
	"go-schooling/business/transactions"

	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMySQLTransactionRepository(conn *gorm.DB) transactions.Repository {
	return &mysqlTransactionRepository{
		Conn: conn,
	}
}

func (repository *mysqlTransactionRepository) Store(ctx context.Context, topupDomain *transactions.Domain) (payments.Domain, error) {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return payments.Domain{}, result.Error
	}

	err := repository.Conn.Preload("User").First(&rec, rec.ID).Error
	if err != nil {
		return payments.Domain{}, result.Error
	}

	return rec.toPaymentDomain(), nil
}

func (repository *mysqlTransactionRepository) Update(ctx context.Context, topupDomain *transactions.Domain) error {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Where("order_id = ?", rec.OrderID).Updates(&rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlTransactionRepository) GetByID(ctx context.Context, id int) (transactions.Domain, error) {
	transaction := Transactions{}
	result := repository.Conn.Where("user_id = ?", id).Find(&transaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return transaction.toDomain(), nil
}

func (repository *mysqlTransactionRepository) GetByOrder(ctx context.Context, orderId string) (transactions.Domain, error) {
	transaction := Transactions{}
	result := repository.Conn.Where("order_id = ?", orderId).Find(&transaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return transaction.toDomain(), nil
}
