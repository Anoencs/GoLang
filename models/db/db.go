package db

import (
	"log"
	"model/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
}

func (database *Database) Connect() *gorm.DB {
	dbURL := "postgres://imeter_user:iMeter@123@10.0.0.232:5432/imeter?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func (database *Database) Init() {

	dbURL := "postgres://imeter_user:iMeter@123@10.0.0.232:5432/imeter?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.Migrator().CreateTable(&models.StaffEntity{})
	db.Migrator().CreateTable(&models.StaffRoute{})
	db.Migrator().CreateTable(&models.Token{})

}
