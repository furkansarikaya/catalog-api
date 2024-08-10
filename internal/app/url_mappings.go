package app

import "github.com/furkansarikaya/catalog-api/internal/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
