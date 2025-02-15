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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Blogs) *exceptions.Exception {
	var existingTags []models.BlogTags
	var newTags []models.BlogTags

	for _, tag := range data.Tags {
		var existingTag models.BlogTags

		if err := tx.Where("tag = ?", tag.Tag).First(&existingTag).Error; err != nil {
			if err == gorm.ErrRecordNotFound {

				newTags = append(newTags, tag)
			} else {

				return exceptions.ParseGormError(tx, err)
			}
		} else {
			existingTags = append(existingTags, existingTag)
		}
	}

	data.Tags = append(existingTags, newTags...)

	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
