package repositories

import (
	"profile-api/models"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Incognitos) *exceptions.Exception
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Incognitos, *exceptions.Exception)
	FindAll(ctx *gin.Context, tx *gorm.DB) (*[]models.Incognitos, *exceptions.Exception)
	Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
}
