package images

import (
	"errors"
	"go-schooling/business/images"
	"go-schooling/controllers/images/request"
	"go-schooling/controllers/images/response"
	base_response "go-schooling/helper/response"
	"io/ioutil"
	"net/http"

	"github.com/JoinVerse/xid"
	"github.com/labstack/echo/v4"
)

type ImageController struct {
	imageUsecase images.Usecase
}

func NewImageController(ic images.Usecase) *ImageController {
	return &ImageController{
		imageUsecase: ic,
	}
}

func (controller *ImageController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Images{}

	file, err := c.FormFile("file")
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("can't open file"))
	}

	// defer src.Close()

	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, int64(5000000))
	c.Request().ParseMultipartForm(int64(5000000))
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("file is too large"))
	}

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("file is too large"))
	}

	filetype := http.DetectContentType(fileBytes)
	if filetype != "image/jpeg" && filetype != "image/jpg" &&
		filetype != "image/gif" && filetype != "image/png" {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, errors.New("invalid file type"))
	}

	// fileEndings, err := mime.ExtensionsByType(filetype)
	// if err != nil {
	// 	return base_response.NewErrorResponse(c, http.StatusBadRequest, errors.New("can't read file type"))
	// }
	req.Name = file.Filename
	filePath := "articles-" + xid.New().String()
	req.Path = filePath

	path, err := controller.imageUsecase.Store(ctx, req.ToDomain(), file)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp := response.Images{}
	resp.Message = "Success"
	resp.Path = path
	return base_response.NewSuccessResponse(c, resp)
}
