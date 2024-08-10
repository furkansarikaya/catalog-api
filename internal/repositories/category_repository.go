package repositories

import (
	"github.com/furkansarikaya/catalog-api/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]*models.Category, error)
	FindById(id uint) (*models.Category, error)
	Save(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return *categoryRepository{db}
}

func (c categoryRepository) FindAll() ([]*models.Category, error) {
	var categories []*models.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepository) FindById(id uint) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepository) Save(category *models.Category) error {
	return c.db.Save(category).Error
}

func (c categoryRepository) Update(category *models.Category) error {
	return c.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(category).Error
}

func (c categoryRepository) Delete(id uint) error {
	return c.db.Delete(&models.Category{}, id).Error
}
