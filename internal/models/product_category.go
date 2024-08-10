package models

type ProductCategory struct {
	ID         uint `gorm:"primaryKey"`
	ProductId  uint `json:"product_id"`
	CategoryId uint `json:"category_id"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}
