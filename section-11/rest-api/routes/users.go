package routes

import (
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// login validates and login user
func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to user struct",
			"error":   err.Error(),
		})
		return
	}

	err = user.ValidateCredential()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to authenticate user",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to authenticate user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
		"token":   token,
	})
}

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

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"event":   user,
	})
}
