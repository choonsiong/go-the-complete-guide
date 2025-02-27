package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/events"
	"github.com/gin-gonic/gin"
)

// Dependencies:
// go get -u github.com/gin-gonic/gin

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
