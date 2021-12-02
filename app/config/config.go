package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool `mapstructure:"DEBUG"`

	// //! Server
	// SERVER_PORT    string `mapstructure:"SERVER_PORT"`
	// SERVER_TIMEOUT int    `mapstructure:"SERVER_TIMEOUT"`

	// //! MYSQL
	// MYSQL_DB_HOST string `mapstructure:"MYSQL_DB_HOST"`
	// MYSQL_DB_PORT string `mapstructure:"MYSQL_DB_PORT"`
	// MYSQL_DB_USER string `mapstructure:"MYSQL_DB_USER"`
	// MYSQL_DB_PASS string `mapstructure:"MYSQL_DB_PASS"`
	// MYSQL_DB_NAME string `mapstructure:"MYSQL_DB_NAME"`

	// //! MONGO DB
	// MONGO_DB_HOST string `mapstructure:"MONGO_DB_HOST"`
	// MONGO_DB_PORT string `mapstructure:"MONGO_DB_PORT"`
	// MONGO_DB_USER string `mapstructure:"MONGO_DB_USER"`
	// MONGO_DB_PASS string `mapstructure:"MONGO_DB_PASS"`
	// MONGO_DB_NAME string `mapstructure:"MONGO_DB_NAME"`

	// //! OUATH2 GOOGLE
	// GOOGLE_AUTH_CLIENT string `mapstructure:"GOOGLE_AUTH_CLIENT"`
	// GOOGLE_AUTH_SECRET string `mapstructure:"GOOGLE_AUTH_SECRET"`

	// //! OAUTH2 FACEBOOK
	// FACEBOOK_AUTH_CLIENT string `mapstructure:"FACEBOOK_AUTH_CLIENT"`
	// FACEBOOK_AUTH_SECRET string `mapstructure:"FACEBOOK_AUTH_SECRET"`

	// //! MIDTRANS
	// MIDTRANS_SERVER_KEY  string `mapstructure:"MIDTRANS_SERVER_KEY"`
	// MIDTRANS_CLIENT_KEY  string `mapstructure:"MIDTRANS_CLIENT_KEY"`
	// MIDTRANS_MERCHANT_ID string `mapstructure:"MIDTRANS_MERCHANT_ID"`

	// //! JWT
	// JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	// JWT_EXPIRED int    `mapstructure:"JWT_EXPIRED"`

	// //! REDIS
	// REDIS_ENDPOINT string `mapstructure:"REDIS_ENDPOINT"`
	// REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`

	// //! GOOGLE STORAGE
	// GOOGLE_STORAGE_BUCKET_NAME  string `mapstructure:"GOOGLE_STORAGE_BUCKET_NAME"`
	// GOOGLE_STORAGE_PRIVATE_KEY  string `mapstructure:"GOOGLE_STORAGE_PRIVATE_KEY"`
	// GOOGLE_STORAGE_IAM_EMAIL    string `mapstructure:"GOOGLE_STORAGE_IAM_EMAIL"`
	// GOOGLE_STORAGE_EXPIRED_TIME int    `mapstructure:"GOOGLE_STORAGE_EXPIRED_TIME"`

	//! Server
	Server struct {
		Port    string `mapstructure:"address"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`

	//! MYSQL
	Mysql struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"mysql"`

	//! MONGO DB
	Mongo struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"mongo"`

	//! OUATH2 GOOGLE
	Google struct {
		ClientID string `mapstructure:"clientID"`
		Secret   string `mapstructure:"secret"`
	} `mapstructure:"google"`

	//! OAUTH2 FACEBOOK
	Facebook struct {
		ClientID string `mapstructure:"clientID"`
		Secret   string `mapstructure:"secret"`
	} `mapstructure:"facebook"`

	//! MIDTRANS
	Midtrans struct {
		ServerKey  string `mapstructure:"server_key"`
		ClientKey  string `mapstructure:"client_key"`
		MerchantID string `mapstructure:"merchant_id"`
	} `mapstructure:"midtrans"`

	//! JWT
	JWT struct {
		Secret  string `mapstructure:"secret"`
		Expired int    `mapstructure:"expired"`
	} `mapstructure:"jwt"`

	//! REDIS
	Redis struct {
		Host string `mapstructure:"host"`
		Pass string `mapstructure:"password"`
	} `mapstructure:"redis"`

	//! GOOGLE STORAGE
	GoogleStorage struct {
		BucketName string `mapstructure:"bucket_name"`
		PrivateKey string `mapstructure:"private_key"`
		Email      string `mapstructure:"iam_email"`
		Expired    int    `mapstructure:"expired_time"`
	} `mapstructure:"google_storage"`
}

func GetConfig() Config {
	var conf Config

	// viper.AddConfigPath(".")
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(os.Getenv("APP_PATH") + "app/config/")
	// viper.SetConfigFile(`app/config/config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error: ", err)
		// conf.MYSQL.MYSQL_DB_HOST = os.Getenv("MYSQL_DB_HOST")
		// conf.MYSQL.MYSQL_DB_PORT = os.Getenv("MYSQL_DB_PORT")
		// conf.MYSQL.MYSQL_DB_USER = os.Getenv("MYSQL_DB_USER")
		// conf.MYSQL.MYSQL_DB_PASS = os.Getenv("MYSQL_DB_PASS")
		// conf.MYSQL.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")

		// conf.FACEBOOK.FACEBOOK_AUTH_CLIENT = os.Getenv("FACEBOOK_AUTH_CLIENT")
		// conf.FACEBOOK.FACEBOOK_AUTH_SECRET = os.Getenv("FACEBOOK_AUTH_SECRET")

		// conf.GOOGLE.GOOGLE_AUTH_CLIENT = os.Getenv("GOOGLE_AUTH_CLIENT")
		// conf.GOOGLE.GOOGLE_AUTH_SECRET = os.Getenv("GOOGLE_AUTH_SECRET")

		// conf.MIDTRANS.MIDTRANS_SERVER_KEY = os.Getenv("MIDTRANS_SERVER_KEY")
		// conf.MIDTRANS.MIDTRANS_CLIENT_KEY = os.Getenv("MIDTRANS_CLIENT_KEY")
		// conf.MIDTRANS.MIDTRANS_MERCHANT_ID = os.Getenv("MIDTRANS_MERCHANT_ID")

		// conf.JWT.JWT_SECRET = os.Getenv("JWT_SECRET")
		// conf.JWT.JWT_EXPIRED, _ = strconv.Atoi(os.Getenv("JWT_EXPIRED"))

		// conf.REDIS.REDIS_ENDPOINT = os.Getenv("REDIS_ENDPOINT")
		// conf.REDIS.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

		// conf.GOOGLESTORAGE.GOOGLE_STORAGE_BUCKET_NAME = os.Getenv("GOOGLE_STORAGE_BUCKET_NAME")
		// conf.GOOGLESTORAGE.GOOGLE_STORAGE_PRIVATE_KEY = os.Getenv("GOOGLE_STORAGE_PRIVATE_KEY")
		// conf.GOOGLESTORAGE.GOOGLE_STORAGE_IAM_EMAIL = os.Getenv("GOOGLE_STORAGE_IAM_EMAIL")
		// conf.GOOGLESTORAGE.GOOGLE_STORAGE_EXPIRED_TIME, _ = strconv.Atoi(os.Getenv("GOOGLE_STORAGE_EXPIRED_TIME"))

		conf.Mysql.Host = os.Getenv("MYSQL_DB_HOST")
		conf.Mysql.Port = os.Getenv("MYSQL_DB_PORT")
		conf.Mysql.User = os.Getenv("MYSQL_DB_USER")
		conf.Mysql.Pass = os.Getenv("MYSQL_DB_PASS")
		conf.Mysql.Name = os.Getenv("MYSQL_DB_NAME")

		conf.Facebook.ClientID = os.Getenv("FACEBOOK_AUTH_CLIENT")
		conf.Facebook.Secret = os.Getenv("FACEBOOK_AUTH_SECRET")

		conf.Google.ClientID = os.Getenv("GOOGLE_AUTH_CLIENT")
		conf.Google.Secret = os.Getenv("GOOGLE_AUTH_SECRET")

		conf.Midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
		conf.Midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
		conf.Midtrans.MerchantID = os.Getenv("MIDTRANS_MERCHANT_ID")

		conf.JWT.Secret = os.Getenv("JWT_SECRET")
		conf.JWT.Expired, _ = strconv.Atoi(os.Getenv("JWT_EXPIRED"))

		conf.Redis.Host = os.Getenv("REDIS_ENDPOINT")
		conf.Redis.Pass = os.Getenv("REDIS_PASSWORD")

		conf.GoogleStorage.BucketName = os.Getenv("GOOGLE_STORAGE_BUCKET_NAME")
		conf.GoogleStorage.PrivateKey = os.Getenv("GOOGLE_STORAGE_PRIVATE_KEY")
		conf.GoogleStorage.Email = os.Getenv("GOOGLE_STORAGE_IAM_EMAIL")
		conf.GoogleStorage.Expired, _ = strconv.Atoi(os.Getenv("GOOGLE_STORAGE_EXPIRED_TIME"))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}
