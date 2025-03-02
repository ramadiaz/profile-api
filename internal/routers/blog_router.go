package routers

import (
	"profile-api/api/blogs/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	blogGroup := r.Group("/blog")
	{
		blogGroup.POST("/create", compControllers.Create)
		blogGroup.PUT("/update", compControllers.Update)
		blogGroup.DELETE("/delete", compControllers.Delete)
		blogGroup.GET("/all", compControllers.FindAll)
		blogGroup.GET("/slug", compControllers.FindBySlug)
		blogGroup.GET("/uuid", compControllers.FindByUUID)
		blogGroup.GET("/featured", compControllers.FindFeaturedBlogs)

		featuredGroup := blogGroup.Group("/featured")
		{
			featuredGroup.POST("/create", compControllers.CreateFeaturedBlog)
			featuredGroup.DELETE("/delete", compControllers.DeleteFeaturedBlogs)
			featuredGroup.GET("/memorized", compControllers.MemorizedFeaturedBlogs)
		}
	}
}
