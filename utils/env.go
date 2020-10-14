package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetPort() string {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "9000"
		log.Printf("Defaulting to port %s", PORT)
	}

	return PORT
}
