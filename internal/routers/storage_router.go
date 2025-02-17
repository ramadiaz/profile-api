package routers

import (
	"profile-api/api/storages/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	storageGroup := r.Group("/storage")
	{
		imageGroup := storageGroup.Group("/image")
		{
			imageGroup.POST("/upload", compControllers.Images)
		}
	}
}
