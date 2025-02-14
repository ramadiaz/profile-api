package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindByShortURL(ctx *gin.Context)
}
