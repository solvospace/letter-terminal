package routes

import (
	"backend/internal/modules/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterV1Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	auth.Routes(v1.Group("/auth"))

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "App is running",
		})
	})
}
