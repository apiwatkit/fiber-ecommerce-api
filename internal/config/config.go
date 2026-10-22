package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string 
	AppPort string 
	AppURL  string

	TimeZone string

	AutoMigrate bool
	DBHost     string 
	DBPort     string 
	DBName     string 
	DBUser     string
	DBPass string
	DBSSL      string

	JWTSecret string
	JWTExpiresIn string

	AdminEmail		string
	AdminPassword 	string
	AdminFirstName	string
	AdminLastName	string
}

func LoadConfig() *Config  {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default values")
	}

	autoMigrate, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	if err != nil {
		autoMigrate = false
	}
	
	return &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		AppURL:  os.Getenv("APP_URL"),

		TimeZone: os.Getenv("TIME_ZONE"),

		AutoMigrate: autoMigrate,
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBSSL:      os.Getenv("DB_SSL"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiresIn: os.Getenv("JWT_EXPIRES_IN"),

		AdminEmail:		os.Getenv("ADMIN_EMAIL"),
		AdminPassword: 	os.Getenv("ADMIN_PASSWORD"),
		AdminFirstName:	os.Getenv("ADMIN_FIRST_NAME"),
		AdminLastName:	os.Getenv("ADMIN_LAST_NAME"),
	}
}