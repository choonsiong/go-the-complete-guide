package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// getEvents get all events
func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get all events",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"events":  events,
	})
}

// getEventByID get an event by id
func getEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err,
		})
		return
	}

	event, err := models.GetEventByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event by id",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"event":   event,
	})
}

// newEvent insert a new event
func newEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to event struct",
			"error":   err,
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.DateTime = time.Now()

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save event",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"event":   event,
	})
}

// updateEventByID update an event by id
func updateEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err,
		})
		return
	}

	_, err = models.GetEventByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event by id",
			"error":   err,
		})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to event struct",
			"error":   err,
		})
		return
	}

	updatedEvent.ID = int(id)
	updatedEvent.DateTime = time.Now()

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update event",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"event":   updatedEvent,
	})
}
