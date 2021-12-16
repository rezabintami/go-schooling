package request

import (
	"go-schooling/business/transactions"
	"strconv"
)

type Transactions struct {
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
}
type MidtransCallback struct {
	StatusCode           string     `json:"status_code"`
	StatusMessage        string     `json:"status_message"`
	PermataVaNumber      string     `json:"permata_va_number"`
	SignKey              string     `json:"signature_key"`
	CardToken            string     `json:"token_id"`
	SavedCardToken       string     `json:"saved_token_id"`
	SavedTokenExpAt      string     `json:"saved_token_id_expired_at"`
	SecureToken          bool       `json:"secure_token"`
	Bank                 string     `json:"bank"`
	BillerCode           string     `json:"biller_code"`
	BillKey              string     `json:"bill_key"`
	XlTunaiOrderID       string     `json:"xl_tunai_order_id"`
	BIIVaNumber          string     `json:"bii_va_number"`
	ReURL                string     `json:"redirect_url"`
	ECI                  string     `json:"eci"`
	ValMessages          []string   `json:"validation_messages"`
	Page                 int        `json:"page"`
	TotalPage            int        `json:"total_page"`
	TotalRecord          int        `json:"total_record"`
	FraudStatus          string     `json:"fraud_status"`
	PaymentType          string     `json:"payment_type"`
	OrderID              string     `json:"order_id"`
	TransactionID        string     `json:"transaction_id"`
	TransactionTime      string     `json:"transaction_time"`
	TransactionStatus    string     `json:"transaction_status"`
	GrossAmount          string     `json:"gross_amount"`
	VANumbers            []VANumber `json:"va_numbers"`
	PaymentCode          string     `json:"payment_code"`
	Store                string     `json:"store"`
	MerchantID           string     `json:"merchant_id"`
	MaskedCard           string     `json:"masked_card"`
	Currency             string     `json:"currency"`
	CardType             string     `json:"card_type"`
	RefundChargebackID   int        `json:"refund_chargeback_id"`
	RefundAmount         string     `json:"refund_amount"`
	RefundKey            string     `json:"refund_key"`
	ChannelResponseCode  string     `json:"channel_response_code"`
	ChannelStatusMessage string     `json:"channel_status_message"`
}

type VANumber struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		OrderID:     req.OrderID,
		Amount:      req.Amount,
	}
}

func (req *MidtransCallback) HandlerToDomain() *transactions.Domain {
	price, _ := strconv.ParseFloat(req.GrossAmount, 32)
	return &transactions.Domain{
		Amount:      price,
		OrderID:     req.OrderID,
		Status:      req.TransactionStatus,
		FraudStatus: req.FraudStatus,
		StatusCode:  req.StatusCode,
		SignKey:     req.SignKey,
	}
}
