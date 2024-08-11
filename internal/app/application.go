package app

import (
	"github.com/furkansarikaya/catalog-api/config"
	"github.com/furkansarikaya/catalog-api/internal/controllers"
	"github.com/furkansarikaya/catalog-api/internal/repositories"
	"github.com/furkansarikaya/catalog-api/internal/services"
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

	// Veritabanı bağlantısı
	db, err := config.SetupDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Repository ve Service katmanları
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)

	// CategoryController'ı başlat
	controllers.InitCategoryController(categoryService)

	mapUrls()

	log.Println("Server is running on", cfg.ServerAddress)

	errRouter := router.Run(cfg.ServerAddress)
	if errRouter != nil {
		log.Fatal(errRouter)
	}
}
