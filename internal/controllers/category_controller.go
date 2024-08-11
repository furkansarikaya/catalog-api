// internal/controllers/category_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/furkansarikaya/catalog-api/internal/dtos"
	"github.com/furkansarikaya/catalog-api/internal/services"
	"github.com/gin-gonic/gin"
)

var categoryService services.CategoryService

func InitCategoryController(service services.CategoryService) {
	categoryService = service
}

func GetAllCategories(ctx *gin.Context) {
	categories, err := categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := categoryService.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func CreateCategory(ctx *gin.Context) {
	var categoryDTO dtos.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newCategory, err := categoryService.CreateCategory(categoryDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusCreated, newCategory)
}

func UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var categoryDTO dtos.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	categoryDTO.ID = uint(id)

	updatedCategory, err := categoryService.UpdateCategory(categoryDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = categoryService.DeleteCategory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
