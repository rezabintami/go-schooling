package category

import (
	"go-schooling/business/category"
	"go-schooling/controllers/category/request"
	"go-schooling/controllers/category/response"
	base_response "go-schooling/helper/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoriesController struct {
	categoryUsecase category.Usecase
}

func NewCategoryController(cu category.Usecase) *CategoriesController {
	return &CategoriesController{
		categoryUsecase: cu,
	}
}

func (controller *CategoriesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := controller.categoryUsecase.GetAll(ctx)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(result))
}

func (controller *CategoriesController) GetByActive(c echo.Context) error {
	ctx := c.Request().Context()
	param := c.QueryParam("active")

	active := false
	if param == "true" {
		active = true
	}

	result, err := controller.categoryUsecase.GetByActive(ctx, active)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(result))
}

func (controller *CategoriesController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	category, err := controller.categoryUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(category))
}

func (controller *CategoriesController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.categoryUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *CategoriesController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Category{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.categoryUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}
