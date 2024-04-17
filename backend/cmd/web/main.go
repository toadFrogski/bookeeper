package main

import (
	"bookeeper/database"
	routes "bookeeper/handlers"
	"bookeeper/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

const SOCKET_PATH = "/var/run/www.sock"

// @title Bookeeper backend API
// @version 1.0
// @host      localhost:80
// @schemes http https
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
	v1 := r.Group("/v1")
	{
		routes.GetAuthRoutes(v1)
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
