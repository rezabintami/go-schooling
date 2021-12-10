package routes

import (
	_middleware "go-schooling/app/middleware"
	"go-schooling/controllers/articles"
	"go-schooling/controllers/classes"
	"go-schooling/controllers/images"
	"go-schooling/controllers/teachers"
	"go-schooling/controllers/transactions"
	"go-schooling/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware         middleware.JWTConfig
	UserController        users.UserController
	TeacherController     teachers.TeacherController
	ClassController       classes.ClassController
	ArticleController     articles.ArticleController
	ImageController       images.ImageController
	TransactionController transactions.TransactionsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(_middleware.MiddlewareLogging)

	apiV1 := e.Group("/api/v1")

	//! USERS
	user := apiV1.Group("/user")
	user.GET("/", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	user.PUT("/", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	transaction := user.Group("/transactions")
	transaction.POST("/payment", cl.TransactionController.CreateTransaction, middleware.JWTWithConfig(cl.JWTMiddleware))
	transaction.POST("/payment/callback", cl.TransactionController.TransactionCallbackHandler)
	transaction.GET("/", cl.TransactionController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! CATEGORY
	category := apiV1.Group("/category")
	category.GET("", cl.ClassController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.POST("", cl.ClassController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	category.DELETE("/:id", cl.ClassController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! ARTICLES USER
	article := apiV1.Group("/article")
	article.GET("/", cl.ArticleController.Fetch)

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	//! IMAGES
	upload := apiV1.Group("/images")
	upload.POST("/upload", cl.ImageController.Store)
	upload.GET("/:id", cl.ImageController.GetByID)

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
	adminTeachers := admin.Group("/teachers", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("SUPERUSER"))
	adminTeachers.POST("", cl.TeacherController.Store)
	adminTeachers.GET("", cl.TeacherController.GetAll)
	adminTeachers.GET("/:id", cl.TeacherController.GetByID)
	adminTeachers.PUT("/:id", cl.TeacherController.Update)

	//! ADMIN ARTICLES
	adminArticles := admin.Group("/articles", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("SUPERUSER"))
	adminArticles.POST("", cl.ArticleController.Store)
	adminArticles.GET("", cl.ArticleController.Fetch)
	adminArticles.GET("/:id", cl.ArticleController.GetByID)
	adminArticles.GET("/:title", cl.ArticleController.GetByTitle)
	adminArticles.PUT("/:id", cl.ArticleController.Update)

}
