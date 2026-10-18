package routes

import (
	"backend/internal/modules/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	auth.RegisterRoutes(v1.Group("/auth"))

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "App is running",
		})
	})
}
