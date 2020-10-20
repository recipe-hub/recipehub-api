package routes

import (
	"recipehub-api/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/v1")
	ingredients := v1.Group("/categories")
	{
		ingredients.GET("/", controllers.FindCategories)
		ingredients.GET("/:id", controllers.FindCategoriesById)
		ingredients.POST("/", controllers.CreateCategories)
		ingredients.PUT("/:id", controllers.UpdateCategories)
		ingredients.DELETE("/:id", controllers.DeleteCategories)
	}

	return router
}
