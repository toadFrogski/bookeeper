package main

import (
	"flag"
	"fmt"
	"gg/config"
	"gg/routes/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	defaultConfigPath = "config.yml"
)

func getConfig() *config.ApplicationConfig {
	configPath := flag.String("config", defaultConfigPath, "config file")
	flag.Parse()
	config, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		panic(err)
	}

	fmt.Printf("Config loaded: %+v\n", config)
	return config
}

func getDB(dbconfig *config.DatabaseConfig) *gorm.DB {
	db, err := config.ConnectToDB(dbconfig)
	if err != nil {
		fmt.Printf("Error connect to DB: %+v\n", err)
	}

	return db
}

func main() {
	settings := getConfig()
	r := gin.Default()
	db := getDB(&settings.Database)

	routes.GetV1Routes(r, db)

	r.Run(settings.Server.GetListenAddr())
}
