package config

import (
	"fmt"
	"os"
)

func GetListenAddr() string {
	host := os.Getenv("WEBAPP_HOST")
	port := os.Getenv("WEBAPP_PORT")

	return fmt.Sprintf("%s:%s", host, port)
}

func GetDSN() string {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, database, port,
	)
}
