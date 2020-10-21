package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	IngredientsRoutes(router)
	CategoriesRoutes(router)
	UsersRoutes(router)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Welcome to Recipes Hub API"})
	})

	return router
}
