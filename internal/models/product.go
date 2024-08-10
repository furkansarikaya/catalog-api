package models

type Product struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:nvarchar(max)"`
	Price       float64
	Stock       int `gorm:"default:0"`
}

func (Product) TableName() string {
	return "products"
}
