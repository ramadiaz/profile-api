package routers

import (
	"profile-api/api/likes/controllers"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	likeGroup := r.Group("/like")
	{
		likeGroup.GET("/create", compControllers.Create)
		likeGroup.GET("/current", compControllers.FindCurrentLikes)
	}
}