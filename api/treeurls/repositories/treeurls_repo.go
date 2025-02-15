package repositories

import (
	"profile-api/models"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.TreeURLs) *exceptions.Exception
	FindByShortURL(ctx *gin.Context, tx *gorm.DB, shortURL string) (*models.TreeURLs, *exceptions.Exception)
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.TreeURLs, *exceptions.Exception)
}
