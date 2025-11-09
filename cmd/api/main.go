package main

import (
	"log"
	"net/http"
	"os"

	"test_mekari/internal/handler"
	"test_mekari/internal/repository"
	"test_mekari/internal/routes"
	"test_mekari/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (if exists)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using default values")
	}

	// Get port from environment or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize layers (Dependency Injection)
	todoRepo := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// Setup routes
	router := routes.SetupRoutes(todoHandler)

	// Start server
	log.Printf("🚀 Server starting on port %s...", port)
	log.Printf("📍 API available at http://localhost:%s", port)
	log.Printf("💚 Health check: http://localhost:%s/health", port)
	log.Printf("📚 API docs: http://localhost:%s/", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}
