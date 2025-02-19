package routers

import (
	"profile-api/api/blogs/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	blogGroup := r.Group("/blog")
	{
		blogGroup.POST("/create", compControllers.Create)
		blogGroup.DELETE("/delete", compControllers.Delete)

		featuredGroup := blogGroup.Group("/featured")
		{
			featuredGroup.POST("/create", compControllers.CreateFeaturedBlog)
		}
	}
}
