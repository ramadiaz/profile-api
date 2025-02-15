package routers

import (
	"profile-api/internal/injectors"
	"profile-api/pkg/middleware"

	publicInjectors "profile-api/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	internalController := injectors.InitializeAuthController(validate)
	treeController := publicInjectors.InitializeTreeController(db, validate)
	blogController := publicInjectors.InitializeBlogController(db, validate)

	AuthRoutes(r, internalController)

	r.Use(middleware.InternalMiddleware())
	TreeRoutes(r, treeController)
	BlogRoutes(r, blogController)
}
