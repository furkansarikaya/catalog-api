package app

import (
	"github.com/furkansarikaya/catalog-api/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func init() {
	//Üretim ortamınd, Gin'i release modunda çalıştırın
	gin.SetMode(gin.ReleaseMode)

	//Yeni bir Gin Engine oluşturun
	router = gin.New()

	//Logger ve Recovery middleware'lerini manuel olarak ekleyin
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %s", err)
	}
}

func StartApplication() {
	cfg := config.LoadConfig()

	mapUrls()

	log.Println("Server is running on", cfg.ServerAddress)

	err := router.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
