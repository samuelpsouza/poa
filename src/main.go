package main

import (
	"github.com/joho/godotenv"
	"dev.ssouza/rest-api/api"
)

func main() {
	godotenv.Load()
	api.InitWebServer()
}
