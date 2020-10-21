package routes

import (
	"recipehub-api/controllers"

	"github.com/gin-gonic/gin"
)

func IngredientsRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/v1")
	ingredients := v1.Group("/ingredients")
	{
		ingredients.GET("/", controllers.FindIngredients)
		ingredients.GET("/:id", controllers.FindIngredientsById)
		ingredients.POST("/", controllers.CreateIngredients)
		ingredients.PUT("/:id", controllers.UpdateIngredients)
		ingredients.DELETE("/:id", controllers.DeleteIngredients)
	}

	return router
}
