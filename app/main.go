package main

import (
	"fmt"
	"os"

	_userUsecase "go-schooling/business/users"
	_userController "go-schooling/controllers/users"
	_userRepo "go-schooling/drivers/databases/users"
	"go-schooling/drivers/googlestorage"

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
	_imageController "go-schooling/controllers/images"
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
		DB_Username: configApp.Mysql.User,
		DB_Password: configApp.Mysql.Pass,
		DB_Host:     configApp.Mysql.Host,
		DB_Port:     configApp.Mysql.Port,
		DB_Database: configApp.Mysql.Name,
	}
	fmt.Println("User :", configApp.Mysql.User)
	fmt.Println("Pass :", configApp.Mysql.Pass)
	fmt.Println("Host :", configApp.Mysql.Host)
	fmt.Println("Port :", configApp.Mysql.Port)
	fmt.Println("Name :", configApp.Mysql.Name)
	// mongoConfigDB := _dbMongoDriver.ConfigDB{
	// 	DB_Username: configApp.MONGO_DB_USER,
	// 	DB_Password: configApp.MONGO_DB_PASS,
	// 	DB_Host:     configApp.MONGO_DB_HOST,
	// 	DB_Port:     configApp.MONGO_DB_PORT,
	// 	DB_Database: configApp.MONGO_DB_NAME,
	// }

	mysql_db := mysqlConfigDB.InitialMysqlDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       configApp.JWT.Secret,
		ExpiresDuration: configApp.JWT.Expired,
	}

	configGoogleStorage := googlestorage.Connection{
		BucketName: configApp.GoogleStorage.BucketName,
		PrivateKey: configApp.GoogleStorage.PrivateKey,
		IAMEmail:   configApp.GoogleStorage.Email,
		ExpTime:    configApp.GoogleStorage.Expired,
	}

	timeoutContext := time.Duration(configApp.JWT.Expired) * time.Second

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
	imageUsecase := _imageUsecase.NewImageUsecase(imageRepo, configGoogleStorage, timeoutContext)
	imagesCtrl := _imageController.NewImageController(imageUsecase)

	articleRepo := _articleRepo.NewMySQLArticlesRepository(mysql_db)
	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, categoryUsecase, imageUsecase, &configJWT, timeoutContext)
	articleCtrl := _articleController.NewArticleController(articleUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
		TeacherController: *teacherCtrl,
		ClassController:   *classCtrl,
		ArticleController: *articleCtrl,
		ImageController:   *imagesCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
