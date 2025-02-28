package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authenticate)
	authGroup.POST("/events", newEvent)
	authGroup.PUT("/events/:id", updateEventByID)
	authGroup.DELETE("/events/:id", deleteEventByID)

	router.POST("/signup", signUp)
	router.POST("/login", login)
}
