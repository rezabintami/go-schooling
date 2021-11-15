package main

import (
	"os"

	_userUsecase "go-schooling/business/users"
	_userController "go-schooling/controllers/users"
	_userRepo "go-schooling/drivers/databases/users"

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

	
	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
