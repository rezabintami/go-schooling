package routes

import (
	_middleware "go-schooling/app/middleware"
	"go-schooling/controllers/classes"
	"go-schooling/controllers/teachers"
	"go-schooling/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	TeacherController teachers.TeacherController
	ClassController   classes.ClassController
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

	//! TEACHERS
	teachers := apiV1.Group("/teachers")
	class := teachers.Group("/class")
	class.GET("", cl.ClassController.GetAll)
	class.POST("", cl.ClassController.Store)
	class.DELETE("/:id", cl.ClassController.Delete)
	
	//! ADMIN
	admin := apiV1.Group("/admin")
	admin.POST("/login", cl.UserController.Login)

	//! ADMIN TEACHERS
	adminTeachers := admin.Group("/teachers")
	adminTeachers.POST("", cl.TeacherController.Store)
	adminTeachers.GET("", cl.TeacherController.GetAll)
	adminTeachers.GET("/:id", cl.TeacherController.GetByID)
	adminTeachers.PUT("/:id", cl.TeacherController.Update)
}
