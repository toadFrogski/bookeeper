package database

import (
	config "gg/conf"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)

	}

	return db
}
