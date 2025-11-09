package routes

import (
	"net/http"
	"path/filepath"

	"test_mekari/internal/handler"
	"test_mekari/internal/helpers"
	"test_mekari/internal/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all application routes
func SetupRoutes(todoHandler *handler.TodoHandler) *mux.Router {
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware)

	// Define routes
	// User routes
	router.HandleFunc("/users", todoHandler.GetUsers).Methods("GET", "OPTIONS")

	// Todo routes
	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todos/{id}/toggle", todoHandler.ToggleTodo).Methods("PATCH", "OPTIONS")

	// Health check endpoint
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// API info endpoint
	router.HandleFunc("/api", apiInfoHandler).Methods("GET")

	// Root endpoint - serve frontend HTML
	router.HandleFunc("/", frontendHandler).Methods("GET")

	return router
}

// healthCheckHandler handles health check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthData := map[string]string{
		"status":  "healthy",
		"service": "Collaborative Todo List API",
	}
	msg := "Service is running"
	helpers.Success(w, helpers.Get, healthData, &msg, nil)
}

// apiInfoHandler handles API information requests
func apiInfoHandler(w http.ResponseWriter, r *http.Request) {
	apiInfo := map[string]any{
		"name":    "Collaborative Todo List API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"GET /users":               "Get all users",
			"GET /todos":               "Get all todos (optional: ?user_id=1 to filter by user)",
			"POST /todos":              "Create a new todo (user_id must exist)",
			"DELETE /todos/{id}":       "Delete a todo",
			"PUT /todos/{id}":          "Update a todo",
			"PATCH /todos/{id}/toggle": "Toggle todo completed status",
			"GET /health":              "Health check",
			"GET /api":                 "API documentation",
			"GET /":                    "Web interface",
		},
	}
	msg := "Welcome to Collaborative Todo List API"
	helpers.Success(w, helpers.Get, apiInfo, &msg, nil)
}

// frontendHandler serves the frontend HTML file
func frontendHandler(w http.ResponseWriter, r *http.Request) {
	// Get absolute path to views directory
	htmlPath := filepath.Join("views", "index.html")
	http.ServeFile(w, r, htmlPath)
}
