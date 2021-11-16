package classes

import (
	"go-schooling/business/classes"
	"go-schooling/controllers/classes/request"
	"go-schooling/controllers/classes/response"
	base_response "go-schooling/helper/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ClassController struct {
	classUsecase classes.Usecase
}

func NewClassController(cu classes.Usecase) *ClassController {
	return &ClassController{
		classUsecase: cu,
	}
}


func (controller *ClassController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.classUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *ClassController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := controller.classUsecase.GetAll(ctx)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(result))
}

func (controller *ClassController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Classes{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.classUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}