package routers

import (
	"profile-api/internal/injectors"
	"profile-api/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.Use(middleware.GzipResponseMiddleware())
	internalController := injectors.InitializeAuthController(validate)

	AuthRoutes(r, internalController)
}
