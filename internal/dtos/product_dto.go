package dtos

type ProductDto struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Stock       int           `json:"stock"`
	Categories  []CategoryDTO `json:"categories"`
	CategoryIDs []uint        `json:"category_ids"`
}
