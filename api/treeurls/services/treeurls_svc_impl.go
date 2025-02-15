package services

import (
	"fmt"
	"profile-api/api/treeurls/dto"
	"profile-api/api/treeurls/repositories"
	"profile-api/models"
	"profile-api/pkg/exceptions"
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
	treeData, err := compRepositories.FindAll(nil, db)
	if err != nil {
		fmt.Println("Error fetching tree data:", err)
	}

	memory.Set("tree", treeData)

	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
		memory:   memory,
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

	go s.MemorizedTree()

	return nil
}

func (s *CompServicesImpl) FindByShortURL(ctx *gin.Context, shortURL string) (*dto.TreeURLOutput, *exceptions.Exception) {
	memoryData, ok := s.memory.Get("tree")
	if ok {
		treeData, ok := memoryData.([]models.TreeURLs)
		if ok {
			for _, tree := range treeData {
				if tree.ShortURL == shortURL {
					output := mapper.MapTreeURLModelToOutput(tree)
					return &output, nil
				}
			}
		}
	}

	data, err := s.repo.FindByShortURL(ctx, s.DB, shortURL)
	if err != nil {
		return nil, err
	}

	output := mapper.MapTreeURLModelToOutput(*data)
	return &output, nil
}

func (s *CompServicesImpl) MemorizedTree() *exceptions.Exception {
	treeData, err := s.repo.FindAll(nil, s.DB)
	if err != nil {
		return err
	}

	s.memory.Set("tree", treeData)

	return nil
}
