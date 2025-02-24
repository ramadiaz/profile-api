package services

import (
	"net/http"
	"profile-api/api/blogs/dto"
	"profile-api/api/blogs/repositories"
	"profile-api/pkg/exceptions"
	"profile-api/pkg/helpers"
	"profile-api/pkg/mapper"
	"profile-api/storages"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
	memory   *storages.Memory
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, memory *storages.Memory) CompServices {
	services := &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
		memory:   memory,
	}

	go services.MemorizedFeaturedBlogs()
	return services
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

	go s.MemorizedFeaturedBlogs()

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

	go s.MemorizedFeaturedBlogs()

	return nil
}

func (s *CompServicesImpl) FindFeaturedBlogs(ctx *gin.Context) (*dto.FeaturedBlogOutput, *exceptions.Exception) {
	memoryData, ok := s.memory.Get("featured_blogs")
	if ok {
		featuredData, ok := memoryData.(dto.FeaturedBlogOutput)
		if ok {
			return &featuredData, nil
		}
	}

	var results dto.FeaturedBlogOutput

	hotBlog, err := s.repo.FindHotBlog(ctx, s.DB)
	if err != nil && err.Status != http.StatusNotFound {
		return nil, err
	}

	if hotBlog != nil && hotBlog.BlogUUID != "" {
		output := mapper.MapBlogModelToOutput(hotBlog.Blog)
		results.HotBlog = &output
	} else {
		results.HotBlog = nil
	}

	featuredBlogs, err := s.repo.FindFeaturedBlogs(ctx, s.DB)
	if err != nil && err.Status != http.StatusNotFound {
		return nil, err
	}

	for _, item := range featuredBlogs {
		results.FeaturedBlogs = append(results.FeaturedBlogs, mapper.MapBlogModelToOutput(item.Blog))
	}

	latestBlogs, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	for _, item := range latestBlogs {
		results.Latest = append(results.Latest, mapper.MapBlogModelToOutput(item))

	}

	go s.memory.Set("featured_blogs", &results)

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

func (s *CompServicesImpl) DeleteFeaturedBlogs(ctx *gin.Context, data dto.FeaturedBlogs) *exceptions.Exception {
	return s.repo.DeleteFeaturedBlogs(ctx, s.DB, data)
}

func (s *CompServicesImpl) MemorizedFeaturedBlogs() *exceptions.Exception {
	var results dto.FeaturedBlogOutput

	hotBlog, err := s.repo.FindHotBlog(nil, s.DB)
	if err != nil && err.Status != http.StatusNotFound {
		return err
	}

	if hotBlog != nil && hotBlog.BlogUUID != "" {
		output := mapper.MapBlogModelToOutput(hotBlog.Blog)
		results.HotBlog = &output
	} else {
		results.HotBlog = nil
	}

	featuredBlogs, err := s.repo.FindFeaturedBlogs(nil, s.DB)
	if err != nil && err.Status != http.StatusNotFound {
		return err
	}

	for _, item := range featuredBlogs {
		results.FeaturedBlogs = append(results.FeaturedBlogs, mapper.MapBlogModelToOutput(item.Blog))
	}

	latestBlogs, err := s.repo.FindAll(nil, s.DB)
	if err != nil {
		return err
	}

	for _, item := range latestBlogs {
		results.Latest = append(results.Latest, mapper.MapBlogModelToOutput(item))

	}

	s.memory.Set("featured_blogs", &results)

	return nil
}
