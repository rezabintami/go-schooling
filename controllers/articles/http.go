package articles

import (
	"go-schooling/business/articles"
	"go-schooling/controllers/articles/request"
	"go-schooling/controllers/articles/response"
	"net/http"
	"strconv"

	base_response "go-schooling/helper/response"

	"github.com/gosimple/slug"
	echo "github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUsecase articles.Usecase
}

func NewArticleController(au articles.Usecase) *ArticleController {
	return &ArticleController{
		articleUsecase: au,
	}
}

func (controller *ArticleController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Articles{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	req.Title = slug.Make(req.Title)

	err := controller.articleUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *ArticleController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	req := request.Articles{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.articleUsecase.Update(ctx, req.ToDomain(), idInt)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	article, err := controller.articleUsecase.GetByID(ctx, idInt)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessResponse(c, response.FromDomain(article))
}

func (controller *ArticleController) GetByTitle(c echo.Context) error {
	ctx := c.Request().Context()

	title := c.Param("title")

	articles, err := controller.articleUsecase.GetByTitle(ctx, title)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(articles))
}

func (controller *ArticleController) GetByCategory(c echo.Context) error {
	ctx := c.Request().Context()

	var listCategory []string
	values, _ := c.FormParams()
	listCategory = values["category"]
	articles, err := controller.articleUsecase.GetByCategory(ctx, listCategory)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessResponse(c, response.FromListDomain(articles))
}

func (controller *ArticleController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	articles, err := controller.articleUsecase.GetByID(ctx, idInt)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(articles))
}

func (controller *ArticleController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))
	articles, count, err := controller.articleUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListPageDomain(articles, count))
}
