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

func (r *CompRepositoriesImpl) FindHotBlog(ctx *gin.Context, tx *gorm.DB) (*models.FeaturedBlogs, *exceptions.Exception) {
	var data models.FeaturedBlogs

	result := tx.
		Preload("Blog").
		Where("type = ?", models.Hot).
		First(&data).
		Order("created_at DESC")
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindFeaturedBlogs(ctx *gin.Context, tx *gorm.DB) ([]models.FeaturedBlogs, *exceptions.Exception) {
	var data []models.FeaturedBlogs

	result := tx.
		Preload("Blog").
		Where("type = ?", models.Featured).
		Find(&data).
		Order("created_at DESC")
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Blogs, *exceptions.Exception) {
	var data []models.Blogs

	result := tx.
		Preload("Tags").
		Find(&data).
		Order("created_at DESC")
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Blogs, *exceptions.Exception) {
	var data models.Blogs

	result := tx.
		Preload("Tags").
		Where("slug = ?", slug).
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Blogs, *exceptions.Exception) {
	var data models.Blogs

	result := tx.
		Preload("Tags").
		Where("uuid = ?", uuid).
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	result := tx.Where("uuid = ?", uuid).Delete(&models.Blogs{})
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}
