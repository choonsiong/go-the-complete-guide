package main

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events": models.GetAllEvents(),
	})
}

func newEvent(c *gin.Context) {
	var e models.Event
	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	e.ID = 1
	e.UserID = 1
	e.DateTime = time.Now()

	e.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   e,
	})
}
