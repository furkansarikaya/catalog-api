package services

import (
	"github.com/furkansarikaya/catalog-api/internal/dtos"
	"github.com/furkansarikaya/catalog-api/internal/models"
	"github.com/furkansarikaya/catalog-api/internal/repositories"
)

type ProductService interface {
	GetAllProducts() ([]dtos.ProductDto, error)
	GetProductByID(id uint) (*dtos.ProductDto, error)
	CreateProduct(productDTO dtos.ProductDto) (*dtos.ProductDto, error)
	UpdateProduct(productDTO dtos.ProductDto) (*dtos.ProductDto, error)
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo         repositories.ProductRepository
	productCategoryRepo repositories.ProductCategoryRepository
}

func (p productService) GetAllProducts() ([]dtos.ProductDto, error) {
	products, err := p.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var productsDTOS []dtos.ProductDto

	for _, product := range products {
		categories, err := p.productCategoryRepo.GetCategoriesByProductID(product.ID)
		if err != nil {
			return nil, err
		}
		productDTO := toProductDTO(product)
		productDTO.Categories = toCategoryDTOsByProduct(categories)
		productsDTOS = append(productsDTOS, productDTO)
	}
	return productsDTOS, nil
}

func (p productService) GetProductByID(id uint) (*dtos.ProductDto, error) {
	product, err := p.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	categories, err := p.productCategoryRepo.GetCategoriesByProductID(product.ID)
	if err != nil {
		return nil, err
	}
	productDTO := toProductDTO(*product)
	productDTO.Categories = toCategoryDTOsByProduct(categories)

	return &productDTO, nil
}

func (p productService) CreateProduct(productDTO dtos.ProductDto) (*dtos.ProductDto, error) {
	product := toProductEntity(productDTO)

	if err := p.productRepo.Save(&product); err != nil {
		return nil, err
	}

	// Update product-category relationships
	if err := p.productCategoryRepo.UpdateProductCategories(product.ID, productDTO.CategoryIDs); err != nil {
		return nil, err
	}

	result := toProductDTO(product)
	return &result, nil
}

func (p productService) UpdateProduct(productDTO dtos.ProductDto) (*dtos.ProductDto, error) {
	product := toProductEntity(productDTO)
	if err := p.productRepo.Update(&product); err != nil {

		return nil, err
	}

	// Update product-category relationships
	if err := p.productCategoryRepo.UpdateProductCategories(product.ID, productDTO.CategoryIDs); err != nil {
		return nil, err
	}

	result := toProductDTO(product)
	return &result, nil
}

func (p productService) DeleteProduct(id uint) error {
	return p.productRepo.Delete(id)
}

func NewProductService(productRepo repositories.ProductRepository, productCategoryRepo repositories.ProductCategoryRepository) ProductService {
	return &productService{
		productRepo:         productRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

// toProductDTO converts a Product model to ProductDTO
func toProductDTO(product models.Product) dtos.ProductDto {
	return dtos.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

// toProductEntity converts a ProductDTO to Product model
func toProductEntity(dto dtos.ProductDto) models.Product {
	return models.Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
	}
}

// toCategoryDTOs converts a slice of Category models to a slice of CategoryDTOs
func toCategoryDTOsByProduct(categories []models.Category) []dtos.CategoryDTO {
	var categoryDTOs []dtos.CategoryDTO
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, dtos.CategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return categoryDTOs
}
