package app

import "github.com/furkansarikaya/catalog-api/internal/handlers"

func mapUrls() {
	router.GET("/ping", handlers.PingHandler)
}
