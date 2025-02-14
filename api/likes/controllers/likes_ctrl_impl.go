package controllers

import (
	"net/http"
	"profile-api/api/likes/dto"
	"profile-api/api/likes/services"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}


func (h *CompControllersImpl) Create(ctx *gin.Context) {
	data, err := h.services.Create(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "data created successfully",
		Body:    data,
	})
}

func (h *CompControllersImpl) FindCurrentLikes(ctx *gin.Context) {
	data, err := h.services.FindCurrentLikes(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Body:    data,
	})
}