package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

type AuthenticationMethod interface {

}

type Authentication struct {
	Id int
	Key string
	Method AuthenticationMethod
}

type Request struct {
	Id int
	Body string
	Headers []string
	Method string
}

type Response struct {
	Id int
	Key string
	Body string
	StatusCode int
}

type Endpoint struct {
	Id int
	Key string
	Auth AuthenticationMethod
	DefaultResponse Response 
	Requests []Request
	Origin string
	Received string
}


var auths = []Authentication{}
var answers = []Response{}

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
	authId, authErr := strconv.Atoi(c.Query("authId"))
	answerId, ansErr := strconv.Atoi(c.Query("answerId"))

	if authErr != nil || ansErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad request")
	}

	fmt.Println(authId)

	for _, res := range answers {
        if res.Id == answerId {
            c.IndentedJSON(http.StatusOK, res)
            return
        }
    }

	c.IndentedJSON(http.StatusOK, gin.H{"message": "answer not found"})
}

func greetingPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post message"})
}