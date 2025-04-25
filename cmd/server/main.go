package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nnnn24/url_shortener_service/internal/api/handlers"
	"github.com/nnnn24/url_shortener_service/internal/models"
	"github.com/nnnn24/url_shortener_service/internal/repository"
	"github.com/nnnn24/url_shortener_service/internal/service"
	"github.com/nnnn24/url_shortener_service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found")
	}

	cfg := config.Load()

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}

	if err := db.AutoMigrate(&models.Url{}); err != nil {
		log.Fatal("Failed to migrate db")
	}

	// Initialize repositories
	urlRepo := repository.NewURLRepository(db)

	// Initialize service
	urlService := service.NewURLService(urlRepo)

	// Initialize handlers
	urlHandler := handlers.NewURLHandler(urlService)

	// Set up Gin router
	router := gin.Default()

	// Initialize http routes
	initializeHttpRoutes(router, urlHandler)

	// Get port from environment variable or use default
	port := cfg.ServerPort

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func initializeHttpRoutes(router *gin.Engine, urlHandler *handlers.URLHandler) {
	// API group
	api := router.Group("/api")
	{
		// URL routes
		urls := api.Group("/urls")
		{
			urls.POST("", urlHandler.CreateURL)
			urls.GET(":shortCode", urlHandler.FindByShortCode)
			urls.PUT(":shortCode", urlHandler.UpdateURL)
			urls.DELETE(":shortCode", urlHandler.DeleteURL)
		}
	}
}
