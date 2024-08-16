package controllers

import (
	"github.com/furkansarikaya/catalog-api/internal/dtos"
	"github.com/furkansarikaya/catalog-api/internal/services"
	"github.com/furkansarikaya/catalog-api/internal/utils/http_utils"
	"github.com/furkansarikaya/catalog-api/internal/utils/rest_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var categoryService services.CategoryService

func InitCategoryController(service services.CategoryService) {
	categoryService = service
}

func GetAllCategories(ctx *gin.Context) {
	categories, err := categoryService.GetAllCategories()
	if err != nil {
		respErr := rest_errors.NewInternalServerError("Internal Server Error", err)
		http_utils.RespondError(ctx, respErr)
		return
	}
	http_utils.RespondJson(ctx, http.StatusOK, categories)
}

func GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid category ID")
		http_utils.RespondError(ctx, respErr)
		return
	}

	category, err := categoryService.GetCategoryByID(uint(id))
	if err != nil {
		respErr := rest_errors.NewNotFoundError("Category not found")
		http_utils.RespondError(ctx, respErr)
		return
	}

	http_utils.RespondJson(ctx, http.StatusOK, category)
}

func CreateCategory(ctx *gin.Context) {
	var categoryDTO dtos.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid input")
		http_utils.RespondError(ctx, respErr)
		return
	}

	newCategory, err := categoryService.CreateCategory(categoryDTO)
	if err != nil {
		respErr := rest_errors.NewInternalServerError("Internal Server Error", err)
		http_utils.RespondError(ctx, respErr)
		return
	}

	http_utils.RespondJson(ctx, http.StatusCreated, newCategory)
}

func UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid category ID")
		http_utils.RespondError(ctx, respErr)
		return
	}

	var categoryDTO dtos.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid input")
		http_utils.RespondError(ctx, respErr)
		return
	}

	categoryDTO.ID = uint(id)

	updatedCategory, err := categoryService.UpdateCategory(categoryDTO)
	if err != nil {
		respErr := rest_errors.NewInternalServerError("Internal Server Error", err)
		http_utils.RespondError(ctx, respErr)
		return
	}

	http_utils.RespondJson(ctx, http.StatusOK, updatedCategory)
}

func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid category ID")
		http_utils.RespondError(ctx, respErr)
		return
	}

	err = categoryService.DeleteCategory(uint(id))
	if err != nil {
		respErr := rest_errors.NewInternalServerError("Internal Server Error", err)
		http_utils.RespondError(ctx, respErr)
		return
	}

	http_utils.RespondJson(ctx, http.StatusNoContent, nil)
}
