package users

import (
	// "encoding/json"
	// "fmt"

	"net/http"
	"strconv"

	// _config "go-schooling/app/config"
	"go-schooling/app/middleware"
	"go-schooling/business/users"
	"go-schooling/controllers/users/request"
	"go-schooling/controllers/users/response"
	base_response "go-schooling/helper/response"

	echo "github.com/labstack/echo/v4"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/facebook"
	// "golang.org/x/oauth2/google"
)

// var (
// 	googleOauthConfig = &oauth2.Config{
// 		ClientID:     _config.GetConfig().GOOGLE_AUTH_CLIENT,
// 		ClientSecret: _config.GetConfig().GOOGLE_AUTH_SECRET,
// 		RedirectURL:  "https://movie-ticketing-test.herokuapp.com/api/v1/auth/google/callback",
// 		Scopes: []string{
// 			"https://www.googleapis.com/auth/userinfo.email",
// 			"https://www.googleapis.com/auth/userinfo.profile",
// 		},
// 		Endpoint: google.Endpoint,
// 	}
// 	googlerandomstate = "random"
// )

// var (
// 	facebookOauthConfig = &oauth2.Config{
// 		ClientID:     _config.GetConfig().FACEBOOK_AUTH_CLIENT,
// 		ClientSecret: _config.GetConfig().FACEBOOK_AUTH_SECRET,
// 		RedirectURL:  "https://movie-ticketing-test.herokuapp.com/api/v1/auth/facebook/callback",
// 		Scopes:       []string{"public_profile","email"},
// 		Endpoint:     facebook.Endpoint,
// 	}
// 	facebookrandomstate = "random"
// )

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.userUsecase.Register(ctx, req.ToDomain(), false)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Users
	if err := c.Bind(&userLogin); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.userUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}
	return base_response.NewSuccessResponse(c, result)
}

func (controller *UserController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.userUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := controller.userUsecase.GetAll(ctx)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(result))
}

func (controller *UserController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	start := c.QueryParam("page")
	startInt, _ := strconv.Atoi(start)
	last := c.QueryParam("per_page")
	lastInt, _ := strconv.Atoi(last)
	articles, count, err := controller.userUsecase.Fetch(ctx, startInt, lastInt)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListPageDomain(articles, count))
}


// //! OAuth2 Google
// func (controller *UserController) OauthLogin(c echo.Context) error {
// 	return c.Render(http.StatusOK, "oauth.html", nil)
// }

// func (controller *UserController) LoginGoogle(c echo.Context) error {
// 	url := googleOauthConfig.AuthCodeURL(googlerandomstate)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// 	return nil
// }

// func (controller *UserController) HandleGoogle(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	if googlerandomstate != c.QueryParam("state") {
// 		return base_response.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", googlerandomstate))
// 	}

// 	token, err := googleOauthConfig.Exchange(ctx, c.QueryParam("code"))
// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	client := googleOauthConfig.Client(ctx, token)
// 	UserInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}
// 	defer UserInfo.Body.Close()

// 	req := request.Users{}
// 	json.NewDecoder(UserInfo.Body).Decode(&req)

// 	controller.userUsecase.Register(ctx, req.ToDomain(), true)
// 	tokenLogin, err := controller.userUsecase.Login(ctx, req.Email, "", true)

// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}
// 	result := struct {
// 		Token string `json:"token"`
// 	}{Token: tokenLogin}

// 	return base_response.NewSuccessResponse(c, result)
// }

// //! OAuth2 Facebook
// func (controller *UserController) LoginFacebook(c echo.Context) error {
// 	url := facebookOauthConfig.AuthCodeURL(facebookrandomstate)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// 	return nil
// }

// func (controller *UserController) HandleFacebook(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	if facebookrandomstate != c.QueryParam("state") {
// 		return base_response.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", facebookrandomstate))
// 	}

// 	token, err := facebookOauthConfig.Exchange(ctx, c.QueryParam("code"))
// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	client := facebookOauthConfig.Client(ctx, token)
// 	UserInfo, err := client.Get("https://graph.facebook.com/me?fields=name,email")
// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}
// 	defer UserInfo.Body.Close()
	
// 	req := request.Users{}
// 	json.NewDecoder(UserInfo.Body).Decode(&req)

// 	controller.userUsecase.Register(ctx, req.ToDomain(), true)

// 	tokenLogin, err := controller.userUsecase.Login(ctx, req.Email, "", true)

// 	if err != nil {
// 		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}
// 	result := struct {
// 		Token string `json:"token"`
// 	}{Token: tokenLogin}

// 	return base_response.NewSuccessResponse(c, result)
// }
