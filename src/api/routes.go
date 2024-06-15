package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/echo", Greeting)
	router.POST("/echo", GreetingPost)
}
