package repository

import (
	"test_mekari/internal/models"
	"errors"
	"sync"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

// TodoRepository handles data access for todos
type TodoRepository struct {
	todos   []models.Todo
	users   map[int]models.User
	nextID  int
	mu      sync.RWMutex
}

// NewTodoRepository creates a new instance of TodoRepository
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos:  make([]models.Todo, 0),
		users:  SeedUsers(), // Use seeder function to populate initial users
		nextID: 1,
	}
}

// FindAll returns all todos
func (r *TodoRepository) FindAll() []models.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to prevent external modifications
	todosCopy := make([]models.Todo, len(r.todos))
	copy(todosCopy, r.todos)
	return todosCopy
}

// FindByID finds a todo by its ID
func (r *TodoRepository) FindByID(id int) (*models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, todo := range r.todos {
		if todo.ID == id {
			todoCopy := todo
			return &todoCopy, nil
		}
	}
	return nil, ErrTodoNotFound
}

// FindByUserID finds all todos for a specific user
func (r *TodoRepository) FindByUserID(userID int) []models.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userTodos := make([]models.Todo, 0)
	for _, todo := range r.todos {
		if todo.UserID == userID {
			userTodos = append(userTodos, todo)
		}
	}
	return userTodos
}

// Create creates a new todo
func (r *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = r.nextID
	r.nextID++

	r.todos = append(r.todos, *todo)

	// Return a copy
	todoCopy := *todo
	return &todoCopy, nil
}

// Update updates an existing todo
func (r *TodoRepository) Update(todo *models.Todo) (*models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = *todo
			todoCopy := *todo
			return &todoCopy, nil
		}
	}
	return nil, ErrTodoNotFound
}

// Delete deletes a todo by its ID
func (r *TodoRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, todo := range r.todos {
		if todo.ID == id {
			// Remove the todo from slice
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return ErrTodoNotFound
}

// GetUserByID retrieves a user by ID
func (r *TodoRepository) GetUserByID(userID int) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[userID]; exists {
		userCopy := user
		return &userCopy, nil
	}
	return nil, errors.New("user not found")
}

// GetAllUsers returns all users
func (r *TodoRepository) GetAllUsers() []models.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users
}
