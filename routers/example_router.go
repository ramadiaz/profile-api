package routers

import (
	"profile-api/api/example/controllers"

	"github.com/gin-gonic/gin"
)

func ExampleRoutes(r *gin.RouterGroup, exampleController controllers.CompControllers) {
	hotelGroup := r.Group("/example")
	{
		hotelGroup.POST("/create", exampleController.Create)
	}
}
