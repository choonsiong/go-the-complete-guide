package main

import (
	models "example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"events":  events,
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

	err = e.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not save event",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   e,
	})
}
