package services

import (
	"profile-api/api/example/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompService interface {
	Create(ctx *gin.Context, data dto.ExampleInput) *exceptions.Exception
}
