package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
