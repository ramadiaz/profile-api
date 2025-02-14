package services

import (
	"profile-api/api/likes/dto"
	"profile-api/api/likes/repositories"
	"profile-api/models"
	"profile-api/pkg/exceptions"
	"profile-api/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mssola/user_agent"
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

func (s *CompServicesImpl) Create(ctx *gin.Context) (*dto.CurrentLikes, *exceptions.Exception) {
	clientIP := ctx.ClientIP()
	userAgent := ctx.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	name, _ := ua.Browser()

	data := models.Likes{
		IP:      clientIP,
		Browser: name,
		OS:      ua.OS(),
		Device:  ua.Platform(),
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.Create(ctx, tx, data)
	if err != nil {
		return nil, err
	}

	likeData, err := s.repo.FindLast(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &dto.CurrentLikes{
		Count: likeData.ID,
	}, nil
}

func (s *CompServicesImpl) FindCurrentLikes(ctx *gin.Context) (*dto.CurrentLikes, *exceptions.Exception) {
	data, err := s.repo.FindLast(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.CurrentLikes{
		Count: data.ID,
	}, nil
}
