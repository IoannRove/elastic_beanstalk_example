package main

import (
	"app/api/api_v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
    "github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}


func main() {
	router := gin.Default()
	router.Use(cors.Default())
	api_v1.Api(router)
	router.Run()
}
