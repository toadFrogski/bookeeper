package main

import (
	"gg/database"
	routes "gg/handlers"
	"gg/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

const SOCKET_PATH = "/var/run/www.sock"

// @title GG backend API
// @version 1.0
//
// @BasePath /v1
func main() {
	r := gin.Default()
	database.DB = database.InitDB()

	initMiddlewares(r)
	initRoutes(r)

	os.Remove(SOCKET_PATH)
	r.RunUnix(SOCKET_PATH)
}

func initRoutes(r *gin.Engine) {
	routes.GetAuthRoutes(r)
	v1 := r.Group("/v1")
	{
		routes.GetBooksRoutes(v1)
		routes.GetUserRoutes(v1)
	}
}

func initMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.PanicHandleMiddleware())
	r.Use(middlewares.JwtAuthMiddleware())
}
