package routes

import (
	_middleware "go-schooling/app/middleware"
	"go-schooling/controllers/teachers"
	"go-schooling/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	TeacherController teachers.TeacherController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(_middleware.MiddlewareLogging)

	apiV1 := e.Group("/api/v1")

	//! USERS
	apiV1.GET("/users", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.PUT("/users", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	

	//! ADMIN
	admin := apiV1.Group("/admin")
	admin.POST("/login", cl.UserController.Login)

	//! TEACHERS
	teachers := admin.Group("/teachers")
	teachers.POST("", cl.TeacherController.Store)
	teachers.GET("", cl.TeacherController.GetAll)
	teachers.GET("/:id", cl.TeacherController.GetByID)
	teachers.PUT("/:id", cl.TeacherController.Update)
}
