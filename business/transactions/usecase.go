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

func (tu *TransactionUsecase) CreateTransactions(ctx context.Context, transactionDomain *Domain, id int) (payments.DomainResponse, error) {
	transactionDomain.UserID = id
	transactionDomain.Status = "pending"
	result, err := tu.transactionRepository.Store(ctx, transactionDomain)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	response, err := tu.paymentsRepository.Transactions(ctx, &result)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	transactionDomain.PaymentUrl = response.RedirectURL
	err = tu.transactionRepository.Update(ctx, transactionDomain)
	if err != nil {
		return  payments.DomainResponse{}, err
	}

	return response, nil
}

func (tu *TransactionUsecase) Update(ctx context.Context, transactionDomain *Domain) error {
	if transactionDomain.Status == "settlement" {
		transactionDomain.Status = "paid"
	} else if transactionDomain.Status == "deny" || transactionDomain.Status == "expire" || transactionDomain.Status == "cancel" {
		transactionDomain.Status = "canceled"
	}
	if err := statuskey.IsValid(transactionDomain.OrderID, transactionDomain.StatusCode, fmt.Sprintf("%.2f", transactionDomain.Amount), transactionDomain.SignKey, tu.paymentsRepository.NotificationValidationKey()); err != nil {
		return err
	}

	err := tu.transactionRepository.Update(ctx, transactionDomain)
	if err != nil {
		return err
	}

	if transactionDomain.Status == "paid" {
		result, err := tu.transactionRepository.GetByOrder(ctx, transactionDomain.OrderID)
		if err != nil {
			return err
		}

		err = tu.userRepository.Update(ctx, &users.Domain{Status: "Active"}, result.UserID)
		if err != nil {
			return err
		}
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
