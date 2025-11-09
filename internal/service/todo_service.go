package service

import (
	"test_mekari/internal/dto"
	"test_mekari/internal/models"
	"test_mekari/internal/repository"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidTodoText = errors.New("todo text cannot be empty")
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrUserNotFound    = errors.New("user_id not found")
	ErrUnauthorized    = errors.New("unauthorized to perform this action")
)

// TodoService handles business logic for todos
type TodoService struct {
	repo *repository.TodoRepository
}

// NewTodoService creates a new instance of TodoService
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// GetAllTodos returns all todos
func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	todos := s.repo.FindAll()
	return todos, nil
}

// GetAllUsers returns all users
func (s *TodoService) GetAllUsers() ([]models.User, error) {
	users := s.repo.GetAllUsers()
	return users, nil
}

// GetTodosByUser returns todos filtered by user ID
func (s *TodoService) GetTodosByUser(userID int) ([]models.Todo, error) {
	if userID <= 0 {
		return nil, ErrInvalidUserID
	}

	// Check if user exists
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	todos := s.repo.FindByUserID(userID)
	return todos, nil
}

// GetTodoByID returns a single todo by ID
func (s *TodoService) GetTodoByID(id int) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("invalid todo ID")
	}

	return s.repo.FindByID(id)
}

// CreateTodo creates a new todo
func (s *TodoService) CreateTodo(req dto.CreateTodoRequest) (*models.Todo, error) {
	// Validate input
	if err := s.validateTodoRequest(req); err != nil {
		return nil, err
	}

	// Get user information - must exist
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Create todo object
	now := time.Now()
	todo := &models.Todo{
		Text:      strings.TrimSpace(req.Text),
		Completed: req.Completed,
		UserID:    req.UserID,
		CreatedBy: user.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save to repository
	createdTodo, err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}

// DeleteTodo deletes a todo by ID
func (s *TodoService) DeleteTodo(id int) error {
	if id <= 0 {
		return errors.New("invalid todo ID")
	}

	// Check if todo exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Delete the todo
	return s.repo.Delete(id)
}

// ToggleTodo toggles the completed status of a todo
func (s *TodoService) ToggleTodo(id int) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("invalid todo ID")
	}

	// Find the todo
	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Toggle completed status
	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()

	// Update in repository
	return s.repo.Update(todo)
}

// UpdateTodo updates a todo
func (s *TodoService) UpdateTodo(id int, req dto.CreateTodoRequest) (*models.Todo, error) {
	if id <= 0 {
		return nil, errors.New("invalid todo ID")
	}

	// Validate input
	if err := s.validateTodoRequest(req); err != nil {
		return nil, err
	}

	// Find the existing todo
	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	todo.Text = strings.TrimSpace(req.Text)
	todo.Completed = req.Completed
	todo.UpdatedAt = time.Now()

	// Save changes
	return s.repo.Update(todo)
}

// validateTodoRequest validates a todo creation/update request
func (s *TodoService) validateTodoRequest(req dto.CreateTodoRequest) error {
	// Validate text
	if strings.TrimSpace(req.Text) == "" {
		return ErrInvalidTodoText
	}

	// Validate user ID
	if req.UserID <= 0 {
		return ErrInvalidUserID
	}

	return nil
}
