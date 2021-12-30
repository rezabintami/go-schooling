package teachers

import (
	"go-schooling/app/middleware"
	"go-schooling/business/teachers"
	"go-schooling/controllers/teachers/request"
	"go-schooling/controllers/teachers/response"
	base_response "go-schooling/helper/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type TeacherController struct {
	teacherUsecase teachers.Usecase
}

func NewTeacherController(tc teachers.Usecase) *TeacherController {
	return &TeacherController{
		teacherUsecase: tc,
	}
}

func (controller *TeacherController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	teacher, err := controller.teacherUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(&teacher))
}

func (controller *TeacherController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	req := request.Teachers{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.teacherUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	teacher, err := controller.teacherUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessResponse(c, response.FromDomain(&teacher))
}

func (controller *TeacherController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := controller.teacherUsecase.GetAll(ctx)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(result))
}

func (controller *TeacherController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Teachers{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.teacherUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}