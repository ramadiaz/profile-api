package services

import (
	"profile-api/api/incognitos/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Incognitos) *exceptions.Exception
}
