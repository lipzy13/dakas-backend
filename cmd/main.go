package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lipzy13/dakas-backend.git/internal/config"
	"github.com/lipzy13/dakas-backend.git/internal/delivery"
	"github.com/lipzy13/dakas-backend.git/internal/repository"
	"github.com/lipzy13/dakas-backend.git/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	netHTTP "net/http"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gerobakRepo := repository.NewGerobakRepository(db)
	gerobakService := service.NewGerobakService(gerobakRepo)

	router := httprouter.New()

	delivery.SetupRoutes(router, gerobakService)

	log.Println("Starting server on :8080")
	if err := netHTTP.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
