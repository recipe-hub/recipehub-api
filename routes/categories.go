package routes

import (
	"recipehub-api/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/v1")
	categories := v1.Group("/categories")
	{
		categories.GET("/", controllers.FindCategories)
		categories.GET("/:id", controllers.FindCategoriesById)
		categories.POST("/", controllers.CreateCategories)
		categories.PUT("/:id", controllers.UpdateCategories)
		categories.DELETE("/:id", controllers.DeleteCategories)
	}

	return router
}
