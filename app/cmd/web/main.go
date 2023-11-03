package main

import (
	"flag"
	"fmt"
	config "gg/conf"
	"gg/modules/book"
	"gg/modules/user"

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
		panic(err)
	}

	return db
}

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/v1")
	{
		book.GetBooksRoutes(v1, db)
		user.GetUserRoutes(v1, db)
	}
}

func SetupMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
}

func main() {
	settings := getConfig()
	r := gin.Default()
	db := getDB(&settings.Database)

	SetupMiddlewares(r)
	InitRoutes(r, db)

	r.Run(settings.Server.GetListenAddr())
}
