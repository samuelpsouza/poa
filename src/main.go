package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", greeting)

	router.Run("localhost:8080")
}

func greeting(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world!"})
}