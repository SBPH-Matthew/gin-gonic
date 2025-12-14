package routes

import (
	"net/http"

	"github.com/SBPH-Matthew/gin-gonic-be/cmd/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "Hello, World!",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/user/create", handlers.CreateUser)
		api.POST("/register", handlers.Register)
	}
}
