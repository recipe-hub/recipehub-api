package main

import (
	"log"
	"net/http"
	"os"

	models "recipehub-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.ConnectDatabase()

	server := gin.Default()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "9000"
		log.Printf("Defaulting to port %s", PORT)
	}

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Welcome to Recipes Hub API"})
	})

	server.Run(":" + PORT)
}
