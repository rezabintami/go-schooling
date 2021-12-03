package transactions

import (
	"context"
	"fmt"
	"go-schooling/business/payments"
	"go-schooling/business/users"
	"go-schooling/helper/statuskey"
	"time"
)

type TransactionUsecase struct {
	transactionRepository    Repository
	userRepository     users.Repository
	paymentsRepository payments.Repository
	contextTimeout     time.Duration
}

func NewTransactionUsecase(tr Repository, timeout time.Duration, us users.Repository, pay payments.Repository) Usecase {
	return &TransactionUsecase{
		transactionRepository:    tr,
		contextTimeout:     timeout,
		userRepository:     us,
		paymentsRepository: pay,
	}
}

func (tu *TransactionUsecase) CreateTransactions(ctx context.Context, topupDomain *Domain, id int) (payments.DomainResponse, error) {
	//!MIDTRANS
	topupDomain.UserID = id

	result, err := tu.transactionRepository.Store(ctx, topupDomain)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	response, err := tu.paymentsRepository.Transactions(ctx, &result)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	return response, nil
}

func (tu *TransactionUsecase) Update(ctx context.Context, topupDomain *Domain) error {
	if topupDomain.Status == "settlement" {
		topupDomain.Status = "paid"
	} else if topupDomain.Status == "deny" || topupDomain.Status == "expire" || topupDomain.Status == "cancel" {
		topupDomain.Status = "canceled"
	}
	if err := statuskey.IsValid(topupDomain.OrderID, topupDomain.StatusCode, fmt.Sprintf("%.2f", topupDomain.Amount), topupDomain.SignKey, tu.paymentsRepository.NotificationValidationKey()); err != nil {
		return err
	}

	err := tu.transactionRepository.Update(ctx, topupDomain)
	if err != nil {
		return err
	}

	if topupDomain.Status == "paid" {
		result, err := tu.transactionRepository.GetByOrder(ctx, topupDomain.OrderID)
		if err != nil {
			return err
		}
		_, err = tu.userRepository.GetByID(ctx, result.UserID)
		if err != nil {
			return err
		}
		//////////////////////! UPDATE USER
	}

	return nil
}

func (tu *TransactionUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := tu.transactionRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
