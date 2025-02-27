package main

import (
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

// Dependencies:
// go get -u github.com/gin-gonic/gin

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", newEvent)

	server.Run(":8080")
}
