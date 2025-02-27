package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", newEvent)
	router.PUT("/events/:id", updateEventByID)
	router.DELETE("/events/:id", deleteEventByID)
}
