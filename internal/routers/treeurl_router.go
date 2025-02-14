package routers

import (
	"profile-api/api/treeurls/controllers"

	"github.com/gin-gonic/gin"
)

func TreeRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	treeGroup := r.Group("/tree")
	{
		treeGroup.POST("/create", compControllers.Create)
	}
}
