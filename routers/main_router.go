package routers

import (
	"net/http"
	"profile-api/injectors"
	"profile-api/storages"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, memory *storages.Memory, validate *validator.Validate) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	incognitoController := injectors.InitializeIncognitoController(db, validate)
	likeController := injectors.InitializeLikeController(db, validate)
	treeConroller := injectors.InitializeTreeController(db, validate, memory)
	blogConroller := injectors.InitializeBlogController(db, validate, memory)

	IncognitoRoutes(r, incognitoController)
	LikeRoutes(r, likeController)
	TreeRoutes(r, treeConroller)
	BlogRoutes(r, blogConroller)
}
