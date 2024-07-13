package main

import (
	"log"

	"github.com/halowahyudi/rest-api-structure-golang/internal/config"
	"github.com/halowahyudi/rest-api-structure-golang/internal/db"
	"github.com/halowahyudi/rest-api-structure-golang/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig("config/config.yaml")
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    // Initialize the database connection
    if err := db.Init(cfg); err != nil {
        log.Fatalf("Could not initialize database: %v", err)
    }

    // Ensure the database connection is closed when the main function ends
    defer db.Close()

    // Set up Gin router
    r := gin.Default()
    routes.SetupRoutes(r)

    // Start the server
    if err := r.Run(":" + cfg.Server.Port); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
