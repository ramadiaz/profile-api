package services

import (
	"profile-api/api/treeurls/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.TreeURLs) *exceptions.Exception
	FindByShortURL(ctx *gin.Context, shortURL string) (*dto.TreeURLOutput, *exceptions.Exception)
}
