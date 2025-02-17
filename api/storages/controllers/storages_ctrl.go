package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Images(ctx *gin.Context)
}