package repositories

import (
	"profile-api/models"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.TreeURLs) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindByShortURL(ctx *gin.Context, tx *gorm.DB, shortURL string) (*models.TreeURLs, *exceptions.Exception) {
	var data models.TreeURLs

	result := tx.Where("short_url = ?", shortURL).First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}
