package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID := c.GetUint("user_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticated request",
		"user_id": userID,
	})
}
