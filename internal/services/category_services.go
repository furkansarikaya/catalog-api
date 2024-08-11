package services

import (
	"github.com/furkansarikaya/catalog-api/internal/dtos"
	"github.com/furkansarikaya/catalog-api/internal/models"
	"github.com/furkansarikaya/catalog-api/internal/repositories"
)

type CategoryService interface {
	GetAllCategories() ([]dtos.CategoryDTO, error)
	GetCategoryByID(id uint) (*dtos.CategoryDTO, error)
	CreateCategory(categoryDTO dtos.CategoryDTO) (*dtos.CategoryDTO, error)
	UpdateCategory(categoryDTO dtos.CategoryDTO) (*dtos.CategoryDTO, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetAllCategories() ([]dtos.CategoryDTO, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return toCategoryDTOs(categories), nil
}

func (s *categoryService) GetCategoryByID(id uint) (*dtos.CategoryDTO, error) {
	category, err := s.categoryRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	categoryDTO := toCategoryDTO(category)
	return &categoryDTO, nil
}

func (s *categoryService) CreateCategory(categoryDTO dtos.CategoryDTO) (dtos.CategoryDTO, error) {
	category := models.Category{
		Name: categoryDTO.Name,
	}

	if err := s.categoryRepo.Save(&category); err != nil {
		return nil, err
	}

	return toCategoryDTO(category), nil
}

func (s *categoryService) UpdateCategory(categoryDTO dtos.CategoryDTO) (dtos.CategoryDTO, error) {
	category := models.Category{
		ID:   categoryDTO.ID,
		Name: categoryDTO.Name,
	}

	if err := s.categoryRepo.Update(&category); err != nil {
		return nil, err
	}

	return toCategoryDTO(category), nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}

func toCategoryDTO(category models.Category) dtos.CategoryDTO {
	return dtos.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}

func toCategoryDTOs(categories []*models.Category) []dtos.CategoryDTO {
	var categoryDTOs []dtos.CategoryDTO
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, toCategoryDTO(*category))
	}
	return categoryDTOs
}
