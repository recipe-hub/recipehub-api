package routes

import (
	"recipehub-api/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/v1")
	users := v1.Group("/users")
	{
		users.GET("/", controllers.FindUsers)
		users.GET("/:id", controllers.FindUserById)
		users.POST("/", controllers.CreateUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
