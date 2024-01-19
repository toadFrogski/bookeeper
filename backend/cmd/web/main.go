package main

import (
	config "gg/conf"
	"gg/database"
	"gg/handlers"
	"gg/middlewares"

	"github.com/gin-gonic/gin"
)

// @title GG backend API
// @version 1.0
//
// @BasePath /v1
func main() {
	r := gin.Default()
	database.DB = database.InitDB()

	initMiddlewares(r)
	initRoutes(r)
	r.Run(config.GetListenAddr())
}

func initRoutes(r *gin.Engine) {
	handlers.GetAuthRoutes(r)
	v1 := r.Group("/v1")
	{
		handlers.GetBooksRoutes(v1)
		handlers.GetUserRoutes(v1)
	}
}

func initMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.PanicHandleMiddleware())
	r.Use(middlewares.JwtAuthMiddleware())
}
