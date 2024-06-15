package api

import (
	"os"
	"github.com/gin-gonic/gin"
)

func InitWebServer() {
	port     := os.Getenv("PORT")	

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	InitRoutes(router)
	router.Run("localhost:" + port)
}
