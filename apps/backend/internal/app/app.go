package app

import (
	"backend/internal/routes"
	"github.com/gin-gonic/gin"
	"os"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	router := gin.Default()
	routes.RegisterV1Routes(router)

	return &App{
		router: router,
	}
}

func (app *App) Run() error {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	return app.router.Run(port)
}
