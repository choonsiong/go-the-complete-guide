package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signUp register a new user
func signUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to user struct",
			"error":   err.Error(),
		})
		return
	}

	savedUser, err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"event":   savedUser,
	})
}
