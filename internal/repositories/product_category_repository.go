package repositories

import (
	"errors"
	"github.com/furkansarikaya/catalog-api/internal/models"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	UpdateProductCategories(productID uint, categoryIDList []uint) error
}

type productCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{db}
}

func (p productCategoryRepository) UpdateProductCategories(productID uint, categoryIDList []uint) error {
	//Start a transaction
	tx := p.db.Begin()

	//Delete ProductCategory records that are not in the new categoryIDList
	if err := tx.Where("product_id = ? AND category_id NOT IN ?", productID, categoryIDList).Delete(&models.ProductCategory{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//Insert new ProductCategory records for the IDs that do not exist in the database
	for _, categoryID := range categoryIDList {
		var existingRelation models.ProductCategory

		err := tx.Where("product_id = ? AND category_id = ?", productID, categoryID).First(&existingRelation).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRelation := models.ProductCategory{
				ProductId:  productID,
				CategoryId: categoryID,
			}
			if err := tx.Create(&newRelation).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}
