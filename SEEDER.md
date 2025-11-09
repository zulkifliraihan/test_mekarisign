# Database Seeder

## User Seeder

This application uses a **seeder** to create initial user data when the application starts.

### File Location

```
internal/repository/user_seeder.go
```

### How It Works

1. The `SeedUsers()` function is called when `NewTodoRepository()` runs
2. Automatically creates 3 sample users with a `CreatedAt` timestamp
3. Users are stored in in-memory storage (map)

### Sample Users

| User ID | Name | Email |
|---------|------|-------|
| 1 | John Doe | john@example.com |
| 2 | Jane Smith | jane@example.com |
| 3 | Bob Johnson | bob@example.com |

### Seeder Code

```go
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
```

## Add New Users

### Method 1: Edit Seeder File (Recommended)

Edit `internal/repository/user_seeder.go` and add a new user:

```go
func SeedUsers() map[int]models.User {
    now := time.Now()

    return map[int]models.User{
        1: {
            ID:        1,
            Name:      "John Doe",
            Email:     "john@example.com",
            CreatedAt: now,
        },
        // ... existing users ...
        4: {  // New user
            ID:        4,
            Name:      "Alice Wonder",
            Email:     "alice@example.com",
            CreatedAt: now,
        },
    }
}
```

### Method 2: Programmatically (Advanced)

Use the `AddUserToSeed()` function if you want to add users dynamically:

```go
users := SeedUsers()
AddUserToSeed(users, 4, "Alice Wonder", "alice@example.com")
```

## User Validation

This seeder ensures that:

1. **Users are available when creating todos**
   - `user_id` validation works correctly
   - A `user_id not found` error appears if the user does not exist

2. **Data is consistent with timestamps**
   - All users have a `CreatedAt` timestamp
   - The API response shows a valid `created_at`

## Testing Validation

Try creating a todo with a non-existent `user_id`:

```bash
# User ID 999 does not exist in the seeder
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Test todo",
    "user_id": 999,
    "completed": false
  }'
```

**Response:**
```json
{
  "response_code": 404,
  "response_status": "failed-not-found",
  "message": "Error! The resource not found!",
  "errors": "user_id not found"
}
```

## Benefits of the Seeder Pattern

1. **Separation of Concerns**: Seed data is separated from repository logic
2. **Maintainability**: Easy to add/edit sample users
3. **Consistency**: All fields (including `CreatedAt`) are properly populated
4. **Testing**: Initial data is always consistent for testing
5. **Reusability**: The seeder function can be called from anywhere

## Future Enhancement

If using a database, this seeder can be replaced with:
- Database migrations
- SQL seed files
- A database seeder library (e.g., GORM seeders)

For now, this pattern is suitable for in-memory storage and development/testing purposes.
