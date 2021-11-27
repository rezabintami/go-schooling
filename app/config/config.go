package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool `mapstructure:"DEBUG"`

	//! Server
	SERVER_PORT    string `mapstructure:"SERVER_PORT"`
	SERVER_TIMEOUT int    `mapstructure:"SERVER_TIMEOUT"`

	//! MYSQL
	MYSQL_DB_HOST string `mapstructure:"MYSQL_DB_HOST"`
	MYSQL_DB_PORT string `mapstructure:"MYSQL_DB_PORT"`
	MYSQL_DB_USER string `mapstructure:"MYSQL_DB_USER"`
	MYSQL_DB_PASS string `mapstructure:"MYSQL_DB_PASS"`
	MYSQL_DB_NAME string `mapstructure:"MYSQL_DB_NAME"`

	//! MONGO DB
	MONGO_DB_HOST string `mapstructure:"MONGO_DB_HOST"`
	MONGO_DB_PORT string `mapstructure:"MONGO_DB_PORT"`
	MONGO_DB_USER string `mapstructure:"MONGO_DB_USER"`
	MONGO_DB_PASS string `mapstructure:"MONGO_DB_PASS"`
	MONGO_DB_NAME string `mapstructure:"MONGO_DB_NAME"`

	//! OUATH2 GOOGLE
	GOOGLE_AUTH_CLIENT string `mapstructure:"GOOGLE_AUTH_CLIENT"`
	GOOGLE_AUTH_SECRET string `mapstructure:"GOOGLE_AUTH_SECRET"`

	//! OAUTH2 FACEBOOK
	FACEBOOK_AUTH_CLIENT string `mapstructure:"FACEBOOK_AUTH_CLIENT"`
	FACEBOOK_AUTH_SECRET string `mapstructure:"FACEBOOK_AUTH_SECRET"`

	//! MIDTRANS
	MIDTRANS_SERVER_KEY  string `mapstructure:"MIDTRANS_SERVER_KEY"`
	MIDTRANS_CLIENT_KEY  string `mapstructure:"MIDTRANS_CLIENT_KEY"`
	MIDTRANS_MERCHANT_ID string `mapstructure:"MIDTRANS_MERCHANT_ID"`

	//! JWT
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	JWT_EXPIRED int    `mapstructure:"JWT_EXPIRED"`

	//! REDIS
	REDIS_ENDPOINT string `mapstructure:"REDIS_ENDPOINT"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`

	//! GOOGLE STORAGE
	GOOGLE_STORAGE_BUCKET_NAME  string `mapstructure:"GOOGLE_STORAGE_BUCKET_NAME"`
	GOOGLE_STORAGE_PRIVATE_KEY  string `mapstructure:"GOOGLE_STORAGE_PRIVATE_KEY"`
	GOOGLE_STORAGE_IAM_EMAIL    string `mapstructure:"GOOGLE_STORAGE_IAM_EMAIL"`
	GOOGLE_STORAGE_EXPIRED_TIME int    `mapstructure:"GOOGLE_STORAGE_EXPIRED_TIME"`
}

func GetConfig() Config {
	var conf Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		// viper.BindEnv(conf.MYSQL_DB_HOST)
		// viper.BindEnv(conf.MYSQL_DB_PORT)
		// viper.BindEnv(conf.MYSQL_DB_USER)
		// viper.BindEnv(conf.MYSQL_DB_PASS)
		// viper.BindEnv(conf.MYSQL_DB_NAME)

		// viper.BindEnv(conf.FACEBOOK_AUTH_CLIENT)
		// viper.BindEnv(conf.FACEBOOK_AUTH_SECRET)

		// viper.BindEnv(conf.GOOGLE_AUTH_CLIENT)
		// viper.BindEnv(conf.GOOGLE_AUTH_SECRET)

		conf.MYSQL_DB_HOST = os.Getenv("MYSQL_DB_HOST")
		conf.MYSQL_DB_PORT = os.Getenv("MYSQL_DB_PORT")
		conf.MYSQL_DB_USER = os.Getenv("MYSQL_DB_USER")
		conf.MYSQL_DB_PASS = os.Getenv("MYSQL_DB_PASS")
		conf.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")

		conf.FACEBOOK_AUTH_CLIENT = os.Getenv("FACEBOOK_AUTH_CLIENT")
		conf.FACEBOOK_AUTH_SECRET = os.Getenv("FACEBOOK_AUTH_SECRET")

		conf.GOOGLE_AUTH_CLIENT = os.Getenv("GOOGLE_AUTH_CLIENT")
		conf.GOOGLE_AUTH_SECRET = os.Getenv("GOOGLE_AUTH_SECRET")

		conf.MIDTRANS_SERVER_KEY = os.Getenv("MIDTRANS_SERVER_KEY")
		conf.MIDTRANS_CLIENT_KEY = os.Getenv("MIDTRANS_CLIENT_KEY")
		conf.MIDTRANS_MERCHANT_ID = os.Getenv("MIDTRANS_MERCHANT_ID")

		conf.JWT_SECRET = os.Getenv("JWT_SECRET")
		conf.JWT_EXPIRED, _ = strconv.Atoi(os.Getenv("JWT_EXPIRED"))

		conf.REDIS_ENDPOINT = os.Getenv("REDIS_ENDPOINT")
		conf.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

		conf.GOOGLE_STORAGE_BUCKET_NAME = os.Getenv("GOOGLE_STORAGE_BUCKET_NAME")
		conf.GOOGLE_STORAGE_PRIVATE_KEY = os.Getenv("GOOGLE_STORAGE_PRIVATE_KEY")
		conf.GOOGLE_STORAGE_IAM_EMAIL = os.Getenv("GOOGLE_STORAGE_IAM_EMAIL")
		conf.GOOGLE_STORAGE_EXPIRED_TIME, _ = strconv.Atoi(os.Getenv("GOOGLE_STORAGE_EXPIRED_TIME"))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}
