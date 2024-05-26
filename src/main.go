package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
)

type Authentication struct {
	Id int
	Key string
}

type Request struct {
	Id int
	Body string
}

type Response struct {
	Id int
	Body string
	StatusCode int
}

var auths = []Authentication{
	{ID:1, Key: "Basic"}
}

var answers = []Response{
	{ID:1, "hello ans"}
}

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
	authId := c.Query("authId")
	answerId := c.Query("answerId")
	fmt.Println(authId)

	for _, res := range answers {
        if res.ID == answerId {
            c.IndentedJSON(http.StatusOK, res)
            return
        }
    }

	c.IndentedJSON(http.StatusOK, gin.H{"message": "answer not found"})
}

func greetingPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post message"})
}