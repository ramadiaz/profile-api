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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Blogs) (*dto.BlogOutput, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapBlogInputToModel(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.CreateSlug(data.Title)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.Create(ctx, tx, input)
	if err != nil {
		return nil, err
	}

	blogData, err := s.repo.FindByUUID(ctx, tx, input.UUID)
	if err != nil {
		return nil, err
	}

	result := mapper.MapBlogModelToOutput(*blogData)

	return &result, nil
}

func (s *CompServicesImpl) CreateFeaturedBlog(ctx *gin.Context, data dto.FeaturedBlogs) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapFeaturedBlogInputToModel(data)

	err := s.repo.CreateFeaturedBlog(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindFeaturedBlogs(ctx *gin.Context) (*dto.FeaturedBlogOutput, *exceptions.Exception) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hotBlog, err := s.repo.FindHotBlog(ctx, tx)
	if err != nil {
		return nil, err
	}
	hotBlogResult := mapper.MapBlogModelToOutput(hotBlog.Blog)

	featuredBlogs, err := s.repo.FindFeaturedBlogs(ctx, tx)
	if err != nil {
		return nil, err
	}

	var featuredBlogResult []dto.BlogOutput

	for _, item := range featuredBlogs {
		featuredBlogResult = append(featuredBlogResult, mapper.MapBlogModelToOutput(item.Blog))
	}

	latestBlogs, err := s.repo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var latestBlogResult []dto.BlogOutput

	for _, item := range latestBlogs {
		latestBlogResult = append(latestBlogResult, mapper.MapBlogModelToOutput(item))
	}

	results := dto.FeaturedBlogOutput{
		HotBlog:       hotBlogResult,
		FeaturedBlogs: featuredBlogResult,
		Latest:        latestBlogResult,
	}

	return &results, nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) ([]dto.BlogOutput, *exceptions.Exception) {
	data, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var result []dto.BlogOutput

	for _, item := range data {
		result = append(result, mapper.MapBlogModelToOutput(item))
	}

	return result, nil
}

func (s *CompServicesImpl) FindBySlug(ctx *gin.Context, slug string) (*dto.BlogOutput, *exceptions.Exception) {
	data, err := s.repo.FindBySlug(ctx, s.DB, slug)
	if err != nil {
		return nil, err
	}

	output := mapper.MapBlogModelToOutput(*data)
	return &output, nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.BlogOutput, *exceptions.Exception) {
	data, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}

	output := mapper.MapBlogModelToOutput(*data)
	return &output, nil
}

func (s *CompServicesImpl) Delete(ctx *gin.Context, uuid string) *exceptions.Exception {
	return s.repo.Delete(ctx, s.DB, uuid)
}
