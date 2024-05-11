package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/hello", greeting)

	router.Run("localhost:" + port)
}

func greeting(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world!"})
}