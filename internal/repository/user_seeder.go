package repository

import (
	"test_mekari/internal/models"
	"time"
)

// SeedUsers returns initial user data for the application
// This is called during repository initialization to populate sample users
func SeedUsers() map[int]models.User {
	now := time.Now()

	return map[int]models.User{
		1: {
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: now,
		},
		2: {
			ID:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			CreatedAt: now,
		},
		3: {
			ID:        3,
			Name:      "Bob Johnson",
			Email:     "bob@example.com",
			CreatedAt: now,
		},
	}
}

// AddUser adds a new user to the seed data
// This function can be used to add more users programmatically
func AddUserToSeed(users map[int]models.User, id int, name, email string) {
	users[id] = models.User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
