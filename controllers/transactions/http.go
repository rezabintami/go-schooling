package transactions

import (
	"go-schooling/app/middleware"
	"go-schooling/business/transactions"
	"go-schooling/controllers/transactions/request"
	"go-schooling/controllers/transactions/response"
	"go-schooling/helper/guid"
	base_response "go-schooling/helper/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type TransactionsController struct {
	transactionsUsecase transactions.Usecase
}

func NewTransactionsController(tc transactions.Usecase) *TransactionsController {
	return &TransactionsController{
		transactionsUsecase: tc,
	}
}

func (ctrl *TransactionsController) CreateTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	id := middleware.GetUser(c).ID

	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	req.OrderID = guid.GenerateUUID()

	resp, err := ctrl.transactionsUsecase.CreateTransactions(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromPaymentDomain(resp))
}

func (ctrl *TransactionsController) TransactionCallbackHandler(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.MidtransCallback{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	
	err := ctrl.transactionsUsecase.Update(ctx, req.HandlerToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Successfully")
}

func (ctrl *TransactionsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	result, err := ctrl.transactionsUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(result))
}
