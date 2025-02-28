package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// registerForEventWithID registers event for the logged-in user
func registerForEventWithID(c *gin.Context) {
	userID := c.GetInt("userID")

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(int(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "failed to get event by id",
			"event_id": eventID,
			"error":    err.Error(),
		})
		return
	}

	err = event.Register(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "failed to register user for event",
			"event_id": eventID,
			"user_id":  userID,
			"error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

// cancelRegistrationForEventWithID delete the event registration for logged-in user
func cancelRegistrationForEventWithID(c *gin.Context) {
	userID := c.GetInt("userID")

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(int(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "failed to get event by id",
			"event_id": eventID,
			"error":    err.Error(),
		})
		return
	}

	err = event.CancelRegistration(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "failed to cancel user registration for event",
			"event_id": eventID,
			"user_id":  userID,
			"error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
