package models

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
