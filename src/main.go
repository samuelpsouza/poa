package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

type AuthenticationTyper interface {

}

type Authentication struct {
	id int
	key string
	method AuthenticationTyper
}

type Request struct {
	id int
	body string
	headers []string
	method string
}

type Response struct {
	id int
	key string
	body string
	statusCode int
}

type Endpoint struct {
	id int
	key string
	auth AuthenticationTyper
	defaultResponse Response 
	requests []Request
	origin string
	received string
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