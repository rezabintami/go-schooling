package main

import (
	"os"

	_userUsecase "go-schooling/business/users"
	_userController "go-schooling/controllers/users"
	_userRepo "go-schooling/drivers/databases/users"

	_teacherUsecase "go-schooling/business/teachers"
	_teacherController "go-schooling/controllers/teachers"
	_teacherRepo "go-schooling/drivers/databases/teachers"

	_classUsecase "go-schooling/business/classes"
	_classController "go-schooling/controllers/classes"
	_classRepo "go-schooling/drivers/databases/classes"

	_articleUsecase "go-schooling/business/articles"
	_articleController "go-schooling/controllers/articles"
	_articleRepo "go-schooling/drivers/databases/articles"

	_categoryUsecase "go-schooling/business/category"
	_categoryRepo "go-schooling/drivers/databases/category"

	_imageUsecase "go-schooling/business/images"
	_imageRepo "go-schooling/drivers/databases/images"

	// _midtrans "ticketing/drivers/thirdparties/midtrans"

	_config "go-schooling/app/config"
	_dbMysqlDriver "go-schooling/drivers/mysql"

	// _dbRedisDriver "ticketing/drivers/redis"

	_middleware "go-schooling/app/middleware"
	_routes "go-schooling/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configApp := _config.GetConfig()
	mysqlConfigDB := _dbMysqlDriver.ConfigDB{
		DB_Username: configApp.MYSQL_DB_USER,
		DB_Password: configApp.MYSQL_DB_PASS,
		DB_Host:     configApp.MYSQL_DB_HOST,
		DB_Port:     configApp.MYSQL_DB_PORT,
		DB_Database: configApp.MYSQL_DB_NAME,
	}
	// mongoConfigDB := _dbMongoDriver.ConfigDB{
	// 	DB_Username: configApp.MONGO_DB_USER,
	// 	DB_Password: configApp.MONGO_DB_PASS,
	// 	DB_Host:     configApp.MONGO_DB_HOST,
	// 	DB_Port:     configApp.MONGO_DB_PORT,
	// 	DB_Database: configApp.MONGO_DB_NAME,
	// }

	mysql_db := mysqlConfigDB.InitialMysqlDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       configApp.JWT_SECRET,
		ExpiresDuration: configApp.JWT_EXPIRED,
	}

	timeoutContext := time.Duration(configApp.JWT_EXPIRED) * time.Second

	e := echo.New()

	userRepo := _userRepo.NewMySQLUserRepository(mysql_db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	teacherRepo := _teacherRepo.NewMySQLTeachersRepository(mysql_db)
	teacherUsecase := _teacherUsecase.NewTeacherUsecase(teacherRepo, userRepo, &configJWT, timeoutContext)
	teacherCtrl := _teacherController.NewTeacherController(teacherUsecase)

	classRepo := _classRepo.NewMySQLClassesRepository(mysql_db)
	classUsecase := _classUsecase.NewClassUsecase(classRepo, &configJWT, timeoutContext)
	classCtrl := _classController.NewClassController(classUsecase)

	categoryRepo := _categoryRepo.NewMySQLCategoryRepository(mysql_db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepo, timeoutContext)

	imageRepo := _imageRepo.NewMySQLImagesRepository(mysql_db)
	imageUsecase := _imageUsecase.NewImageUsecase(imageRepo, timeoutContext)

	articleRepo := _articleRepo.NewMySQLArticlesRepository(mysql_db)
	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, categoryUsecase, imageUsecase, &configJWT, timeoutContext)
	articleCtrl := _articleController.NewArticleController(articleUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
		TeacherController: *teacherCtrl,
		ClassController:   *classCtrl,
		ArticleController: *articleCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
