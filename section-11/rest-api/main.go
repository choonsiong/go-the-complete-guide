package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"log"
)

// Dependencies:
// go get -u github.com/gin-gonic/gin
// go get github.com/mattn/go-sqlite3
// go get -u golang.org/x/crypto
// go get -u github.com/golang-jwt/jwt/v5

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	log.Fatal(server.Run(":8080"))
}
