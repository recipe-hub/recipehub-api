package main

import (
	"net/http"

	"recipehub-api/models"
	"recipehub-api/routes"
	"recipehub-api/utils"
)

func main() {
	utils.LoadEnv()

	models.ConnectDatabase()
	PORT := utils.GetPort()
	server := routes.Router()

	http.ListenAndServe(":"+PORT, server)
}
