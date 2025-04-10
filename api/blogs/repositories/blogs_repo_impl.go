package repositories

import (
	"profile-api/api/blogs/dto"
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

func (r *CompRepositoriesImpl) CreateFeaturedBlog(ctx *gin.Context, tx *gorm.DB, data models.FeaturedBlogs) *exceptions.Exception {
	if data.Type == models.Hot {
		result := tx.Where("type = ?", models.Hot).Delete(&models.FeaturedBlogs{})
		if result.Error != nil {
			return exceptions.ParseGormError(tx, result.Error)
		}
	}

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
		Preload("Blog.Tags").
		Where("type = ?", models.Hot).
		Order("created_at DESC").
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindFeaturedBlogs(ctx *gin.Context, tx *gorm.DB) ([]models.FeaturedBlogs, *exceptions.Exception) {
	var data []models.FeaturedBlogs

	result := tx.
		Preload("Blog").
		Preload("Blog.Tags").
		Preload("Blog.FeaturedBlogs").
		Where("type = ?", models.Featured).
		Order("created_at DESC").
		Limit(5).
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Blogs, *exceptions.Exception) {
	var data []models.Blogs

	result := tx.
		Preload("Tags").
		Preload("FeaturedBlogs").
		Order("created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Blogs, *exceptions.Exception) {
	var data models.Blogs

	result := tx.
		Preload("Tags").
		Preload("FeaturedBlogs").
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
		Preload("FeaturedBlogs").
		Where("uuid = ?", uuid).
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Blogs) *exceptions.Exception {
	var existingBlog models.Blogs
	if err := tx.Where("uuid = ?", data.UUID).Preload("Tags").First(&existingBlog).Error; err != nil {
		return exceptions.ParseGormError(tx, err)
	}

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

	if err := tx.Model(&existingBlog).Updates(map[string]interface{}{
		"title":   data.Title,
		"content": data.Content,
	}).Error; err != nil {
		return exceptions.ParseGormError(tx, err)
	}

	if err := tx.Model(&existingBlog).Association("Tags").Clear(); err != nil {
		return exceptions.ParseGormError(tx, err)
	}

	data.Tags = append(existingTags, newTags...)
	if err := tx.Model(&existingBlog).Association("Tags").Replace(data.Tags); err != nil {
		return exceptions.ParseGormError(tx, err)
	}

	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	result := tx.Where("uuid = ?", uuid).Delete(&models.Blogs{})
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) DeleteFeaturedBlogs(ctx *gin.Context, tx *gorm.DB, data dto.FeaturedBlogs) *exceptions.Exception {
	result := tx.
		Where("blog_uuid = ?", data.BlogUUID).
		Where("type = ?", data.Type).
		Delete(&models.FeaturedBlogs{})
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}
