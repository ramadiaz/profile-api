package services

import (
	"profile-api/api/likes/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context) (*dto.CurrentLikes, *exceptions.Exception)
	FindCurrentLikes(ctx *gin.Context) (*dto.CurrentLikes, *exceptions.Exception)
}
