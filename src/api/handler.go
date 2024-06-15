package api

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"dev.ssouza/rest-api/service"
)

func Greeting(c *gin.Context) {
	_, authErr := strconv.Atoi(c.Query("authId"))
	_, ansErr := strconv.Atoi(c.Query("answerId"))

	if authErr != nil || ansErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad request")
	}

	c.IndentedJSON(http.StatusOK, service.CreatEndpoint())
}

func GreetingPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post message"})
}