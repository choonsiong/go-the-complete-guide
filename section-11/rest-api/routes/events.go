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
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event by id",
			"id":      id,
			"error":   err.Error(),
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
	//token := c.Request.Header.Get("Authorization")
	//if token == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "unauthorized request",
	//	})
	//	return
	//}
	//
	//email, uid, err := utils.VerifyToken(token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "unauthorized request",
	//		"error":   err.Error(),
	//	})
	//	return
	//}

	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get email from context",
		})
		return
	}

	uid := c.GetInt("userID")

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to event struct",
			"error":   err.Error(),
		})
		return
	}

	event.UserID = uid
	event.DateTime = time.Now()

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"email":   email,
		"user_id": uid,
		"event":   event,
	})
}

// deleteEventByID delete an event by id
func deleteEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event by id",
			"id":      id,
			"error":   err.Error(),
		})
		return
	}

	if event.UserID != c.GetInt("userID") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized delete request",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete event",
			"id":      id,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// updateEventByID update an event by id
func updateEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse id param",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event by id",
			"id":      id,
			"error":   err.Error(),
		})
		return
	}

	if event.UserID != c.GetInt("userID") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized update request",
		})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json to event struct",
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = int(id)
	updatedEvent.DateTime = time.Now()

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update event",
			"id":      id,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"event":   updatedEvent,
	})
}
