package mysql_driver

import (
	"fmt"
	"go-schooling/drivers/databases/articles"
	"go-schooling/drivers/databases/category"
	"go-schooling/drivers/databases/classes"
	"go-schooling/drivers/databases/teachers"
	"go-schooling/drivers/databases/users"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&users.Users{},
		&teachers.Teachers{},
		&classes.Classes{},
		&articles.Articles{},
		&category.Category{},
	)

	return db
}
