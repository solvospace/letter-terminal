package app

import (
	"backend/internal/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	router := gin.Default()
	routes.RegisterRoutes(router)

	return &App{
		router: router,
	}
}

func (a *App) Run() error {
	return a.router.Run(":8080")
}
