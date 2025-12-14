package routes

import (
	"net/http"

	"github.com/SBPH-Matthew/gin-gonic-be/cmd/internal/handlers"
	"github.com/SBPH-Matthew/gin-gonic-be/cmd/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "Hello, World!",
		})
	})

	r.GET("/me", middleware.AuthRequired(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "Hello User",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/user/create", handlers.CreateUser)
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}

}
