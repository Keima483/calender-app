package initializers

import (
	"fmt"
	"github.com/keima483/calender-app/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialiseDB(config *Config) error {
	DNS := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUserName,
		config.DBUserPassword,
		config.DBName,
		config.DBPort,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection failed")
		return err
	}
	return DB.AutoMigrate(&repository.Holiday{}, &repository.User{}, &repository.Task{})
}
