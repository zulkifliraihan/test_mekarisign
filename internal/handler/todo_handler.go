package handler

import (
	"test_mekari/internal/dto"
	"test_mekari/internal/helpers"
	"test_mekari/internal/repository"
	"test_mekari/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TodoHandler handles HTTP requests for todos
type TodoHandler struct {
	service *service.TodoService
}

// NewTodoHandler creates a new instance of TodoHandler
func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

// GetTodos handles GET /todos
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	// Check for user_id query parameter
	userIDStr := r.URL.Query().Get("user_id")

	var todos interface{}
	var err error

	if userIDStr != "" {
		// Filter by user ID
		userID, parseErr := strconv.Atoi(userIDStr)
		if parseErr != nil {
			msg := "Invalid user_id parameter"
			helpers.ErrorBadRequest(w, parseErr.Error(), &msg)
			return
		}

		todos, err = h.service.GetTodosByUser(userID)
		if err != nil {
			msg := "Failed to retrieve todos"
			helpers.ErrorServer(w, err.Error(), &msg)
			return
		}
	} else {
		// Get all todos
		todos, err = h.service.GetAllTodos()
		if err != nil {
			msg := "Failed to retrieve todos"
			helpers.ErrorServer(w, err.Error(), &msg)
			return
		}
	}

	helpers.Success(w, helpers.Get, todos, nil, nil)
}

// GetUsers handles GET /users
func (h *TodoHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		msg := "Failed to retrieve users"
		helpers.ErrorServer(w, err.Error(), &msg)
		return
	}

	helpers.Success(w, helpers.Get, users, nil, nil)
}

// CreateTodo handles POST /todos
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTodoRequest

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		// Parse JSON error to human-readable message
		humanMsg := helpers.ParseJSONError(err)
		helpers.ErrorValidator(w, humanMsg, nil)
		return
	}
	defer r.Body.Close()

	// Create todo through service
	todo, err := h.service.CreateTodo(req)
	if err != nil {
		// Check for specific error types
		if err == service.ErrUserNotFound {
			helpers.ErrorNotFound(w, err.Error(), nil)
			return
		}
		if err == service.ErrInvalidTodoText || err == service.ErrInvalidUserID {
			helpers.ErrorValidator(w, err.Error(), nil)
			return
		}
		msg := "Failed to create todo"
		helpers.ErrorServer(w, err.Error(), &msg)
		return
	}

	helpers.Success(w, helpers.Created, todo, nil, nil)
}

// DeleteTodo handles DELETE /todos/{id}
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := "Invalid todo ID"
		helpers.ErrorBadRequest(w, err.Error(), &msg)
		return
	}

	// Delete todo through service
	if err := h.service.DeleteTodo(id); err != nil {
		if err == repository.ErrTodoNotFound {
			helpers.ErrorNotFound(w, err.Error(), nil)
			return
		}
		msg := "Failed to delete todo"
		helpers.ErrorServer(w, err.Error(), &msg)
		return
	}

	msg := "Todo deleted successfully"
	helpers.Success(w, helpers.Deleted, nil, &msg, nil)
}

// ToggleTodo handles PATCH /todos/{id}/toggle
func (h *TodoHandler) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := "Invalid todo ID"
		helpers.ErrorBadRequest(w, err.Error(), &msg)
		return
	}

	// Toggle todo through service
	todo, err := h.service.ToggleTodo(id)
	if err != nil {
		if err == repository.ErrTodoNotFound {
			helpers.ErrorNotFound(w, err.Error(), nil)
			return
		}
		msg := "Failed to toggle todo"
		helpers.ErrorServer(w, err.Error(), &msg)
		return
	}

	msg := "Todo toggled successfully"
	helpers.Success(w, helpers.Updated, todo, &msg, nil)
}

// UpdateTodo handles PUT /todos/{id}
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := "Invalid todo ID"
		helpers.ErrorBadRequest(w, err.Error(), &msg)
		return
	}

	var req dto.CreateTodoRequest

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		// Parse JSON error to human-readable message
		humanMsg := helpers.ParseJSONError(err)
		helpers.ErrorValidator(w, humanMsg, nil)
		return
	}
	defer r.Body.Close()

	// Update todo through service
	todo, err := h.service.UpdateTodo(id, req)
	if err != nil {
		if err == repository.ErrTodoNotFound {
			helpers.ErrorNotFound(w, err.Error(), nil)
			return
		}
		if err == service.ErrInvalidTodoText || err == service.ErrInvalidUserID {
			helpers.ErrorValidator(w, err.Error(), nil)
			return
		}
		msg := "Failed to update todo"
		helpers.ErrorServer(w, err.Error(), &msg)
		return
	}

	helpers.Success(w, helpers.Updated, todo, nil, nil)
}
