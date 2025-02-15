package services

import (
	"profile-api/api/blogs/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Blogs) *exceptions.Exception
}