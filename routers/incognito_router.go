package routers

import (
	"profile-api/api/incognitos/controllers"

	"github.com/gin-gonic/gin"
)

func IncognitoRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	incognitoGroup := r.Group("/incognito")
	{
		incognitoGroup.POST("/create", compControllers.Create)
	}
}
