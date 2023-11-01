package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config *DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()))
	if err != nil {
		return nil, err
	}

	return db, nil
}
