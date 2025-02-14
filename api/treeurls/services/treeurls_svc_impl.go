package services

import (
	"profile-api/api/treeurls/dto"
	"profile-api/api/treeurls/repositories"
	"profile-api/pkg/exceptions"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.TreeURLs) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapTreeURLInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindByShortURL(ctx *gin.Context, shortURL string) (*dto.TreeURLOutput, *exceptions.Exception) {
	data, err := s.repo.FindByShortURL(ctx, s.DB, shortURL)
	if err != nil {
		return nil, err
	}

	output := mapper.MapTreeURLModelToOutput(*data)
	return &output, nil
}
