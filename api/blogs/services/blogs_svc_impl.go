package services

import (
	"profile-api/api/blogs/dto"
	"profile-api/api/blogs/repositories"
	"profile-api/pkg/exceptions"
	"profile-api/pkg/helpers"
	"profile-api/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Blogs) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapBlogInputToModel(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.CreateSlug(data.Title)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}
