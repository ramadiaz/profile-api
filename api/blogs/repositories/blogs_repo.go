package repositories

import (
	"profile-api/models"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Blogs) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Blogs, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Blogs, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Blogs, *exceptions.Exception)
	Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
}
