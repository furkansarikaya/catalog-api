package repositories

import (
	"github.com/furkansarikaya/catalog-api/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Save(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := p.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Save(product *models.Product) error {
	return p.db.Save(product).Error
}

func (p *productRepository) Update(product *models.Product) error {
	// Burada Save() metodu da kullanılabilir, ama explicit olarak Update kullanmak istenirse:
	return p.db.Model(&models.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&models.Product{}, id).Error
}
