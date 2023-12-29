package initializers

import (
	// "github.com/joho/godotenv"
	// "log"
	"os"
)

type Config struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
	AdminEmail     string
	AdminPassword  string
	JwtSecret      string
}

func LoadConfig(path string) Config {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	config := Config{
		DBHost:         os.Getenv("PGHOST"),
		DBUserName:     os.Getenv("PGUSER"),
		DBUserPassword: os.Getenv("PGPASSWORD"),
		DBName:         os.Getenv("PGDATABASE"),
		DBPort:         os.Getenv("PGPORT"),
		AdminEmail:     os.Getenv("ADMINEMAIL"),
		AdminPassword:  os.Getenv("ADMINPASSWORD"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
	}
	return config
}
