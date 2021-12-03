package response

import (
	"go-schooling/business/payments"
	"go-schooling/business/transactions"
)

type Transactions struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Payment struct {
	Token         string   `json:"token"`
	RedirectURL   string   `json:"redirect_url"`
	StatusCode    string   `json:"status_code"`
	ErrorMessages []string `json:"error"`
}

func FromDomain(transactionDomain transactions.Domain) *Transactions {
	return &Transactions{
		ID:     transactionDomain.ID,
		Name:   transactionDomain.Name,
		UserID: transactionDomain.UserID,
	}
}

func FromPaymentDomain(paymentsDomain payments.DomainResponse) Payment {
	return Payment{
		Token:         paymentsDomain.Token,
		RedirectURL:   paymentsDomain.RedirectURL,
		StatusCode:    paymentsDomain.StatusCode,
		ErrorMessages: paymentsDomain.ErrorMessages,
	}
}
