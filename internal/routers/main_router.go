package routers

import (
	"profile-api/internal/injectors"
	"profile-api/pkg/middleware"
	"profile-api/storages"

	publicInjectors "profile-api/injectors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, storage *s3.Client, memory *storages.Memory, validate *validator.Validate) {
	internalController := injectors.InitializeAuthController(validate)
	treeController := publicInjectors.InitializeTreeController(db, validate, memory)
	blogController := publicInjectors.InitializeBlogController(db, validate, memory)
	storageController := publicInjectors.InitializeStorageController(db, storage, validate)

	AuthRoutes(r, internalController)

	r.Use(middleware.InternalMiddleware())
	TreeRoutes(r, treeController)
	BlogRoutes(r, blogController)
	StorageRoutes(r, storageController)
}
