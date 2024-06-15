package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"strconv"
	"dev.ssouza/rest-api/service"
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
	router.GET("/echo", greeting)
	router.POST("/echo", greetingPost)

	router.Run("localhost:" + port)
}

func greeting(c *gin.Context) {
	_, authErr := strconv.Atoi(c.Query("authId"))
	_, ansErr := strconv.Atoi(c.Query("answerId"))

	if authErr != nil || ansErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad request")
	}

	c.IndentedJSON(http.StatusOK, service.CreatEndpoint())
}

func greetingPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post message"})
}