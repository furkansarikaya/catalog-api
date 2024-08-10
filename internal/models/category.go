package models

type Category struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100);not null"`
}

func (Category) TableName() string {
	return "category"
}
