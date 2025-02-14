package services

import (
	"profile-api/internal/auth/dto"
	"profile-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception)
}
