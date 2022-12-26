package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DBConnection -> return db instance
func DBConnection() (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	url := "postgres://postgres:1@10.10.10.194:5439/casbin_golang?sslmode=disable"
	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})

}
