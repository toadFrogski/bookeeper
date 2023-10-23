package app

import (
	"gg/pkg/config"
	"gg/pkg/routes/v1"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *config.ApplicationConfig
	router *gin.Engine
}

func ListenAndServe(config *config.ApplicationConfig) (*Application, error) {
	var app Application

	router := gin.Default()

	routes.GetV1Routes(router)

	if err := router.Run(config.Server.GetListenAddr()); err != nil {
		return nil, err
	}

	app.router = router
	app.config = config

	return &app, nil
}
