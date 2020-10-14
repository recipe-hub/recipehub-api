package main

import (
	"net/http"
	models "recipehub-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Wellcome to Books Go API"})
	})

	// Run the server
	r.Run()
}
