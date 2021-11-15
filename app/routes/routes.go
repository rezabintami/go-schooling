package routes

import (
	_middleware "go-schooling/app/middleware"
	"go-schooling/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(_middleware.MiddlewareLogging)

	apiV1 := e.Group("/api/v1")

	//! USERS
	apiV1.GET("/users", cl.UserController.GetProfile, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.PUT("/users", cl.UserController.UpdateProfile, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	//! ADMIN
	admin := apiV1.Group("/admin")
	admin.POST("/login", cl.UserController.Login)
}
