package config

import (
	"fmt"
	"github.com/furkansarikaya/catalog-api/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	SSLMode       string
}

// LoadConfig, .env dosyasını ve ortam değişkenlerinden yapılandırmayı yükler
func LoadConfig() *Config {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default environment variables")
	}

	cfg := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBName:        getEnv("DB_NAME", "catalog_db"),
		SSLMode:       getEnv("SSL_MODE", "disable"),
	}
	return cfg
}

// getEnv, ortam değişkenini okur, yoksa varsayılan değeri döndürür
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetDatabaseConnectionString, PostgreSQL bağlantı dizesini döndürür
func (c *Config) GetDatabaseConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.SSLMode,
	)
}

// SetupDatabase, PostgreSQL veritabanı bağlantısını başlatır ve döndürür
func SetupDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.GetDatabaseConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	// Veritabanı şemasını güncelle
	err = db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.ProductCategory{},
	)
	if err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	}

	return db, nil
}

func init() {
	// Yapılandırmayı yükle ve hataları kontrol et
	cfg := LoadConfig()
	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		log.Fatal("Gerekli veritabanı yapılandırmaları eksik!")
	}
}
