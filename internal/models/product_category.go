package models

type ProductCategory struct {
	ID         uint     `gorm:"primaryKey"`
	ProductId  uint     `gorm:"not null"`
	Product    Product  `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;"`
	CategoryId uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE;"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}
