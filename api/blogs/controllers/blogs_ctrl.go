package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	CreateFeaturedBlog(ctx *gin.Context)
	FindFeaturedBlogs(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindBySlug(ctx *gin.Context)
	FindByUUID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
