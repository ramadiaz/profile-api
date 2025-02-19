package services

import (
	"profile-api/api/blogs/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Blogs) (*dto.BlogOutput, *exceptions.Exception)
	FindAll(ctx *gin.Context) ([]dto.BlogOutput, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, slug string) (*dto.BlogOutput, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, uuid string) (*dto.BlogOutput, *exceptions.Exception)
	Delete(ctx *gin.Context, uuid string) *exceptions.Exception
}
