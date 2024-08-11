package app

import "github.com/furkansarikaya/catalog-api/internal/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/categories", controllers.GetAllCategories)
	router.GET("/categories/:id", controllers.GetCategoryByID)
	router.POST("/categories", controllers.CreateCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
}
