package controllers

import (
	"net/http"
	"profile-api/api/treeurls/dto"
	"profile-api/api/treeurls/services"
	"profile-api/pkg/exceptions"

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
	var data dto.TreeURLs

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Create(ctx, data)
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

func (h *CompControllersImpl) FindByShortURL(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")
	if shortURL == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	data, err := h.services.FindByShortURL(ctx, shortURL)
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
