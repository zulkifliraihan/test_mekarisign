# Collaborative Todo List - Design Documentation

## 1. Entity Relationship Diagram (ERD)

```
┌─────────────────┐
│      User       │
├─────────────────┤
│ id (int)        │ PK
│ name (string)   │
│ email (string)  │
│ created_at      │
└─────────────────┘
         │
         │ 1:N
         │
         ▼
┌─────────────────┐
│      Todo       │
├─────────────────┤
│ id (int)        │ PK
│ text (string)   │
│ completed (bool)│
│ user_id (int)   │ FK
│ created_by      │
│ created_at      │
│ updated_at      │
└─────────────────┘
```

## 2. Database Model

### User Entity
```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Todo Entity
```go
type Todo struct {
    ID        int       `json:"id"`
    Text      string    `json:"text"`
    Completed bool      `json:"completed"`
    UserID    int       `json:"user_id"`
    CreatedBy string    `json:"created_by"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## 3. Architecture Design (Service Layer Pattern)

```
┌─────────────────────────────────────────────────────────────┐
│                         Client (Browser/App)                 │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ HTTP/JSON
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Presentation Layer                      │
│  ┌────────────┐  ┌────────────┐  ┌────────────┐            │
│  │   Handler  │  │   Handler  │  │   Handler  │            │
│  │   (GET)    │  │   (POST)   │  │  (DELETE)  │            │
│  └────────────┘  └────────────┘  └────────────┘            │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ Function Calls
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       Service Layer                          │
│  ┌────────────────────────────────────────────────┐         │
│  │          TodoService (Business Logic)          │         │
│  │  • GetAllTodos(userID)                         │         │
│  │  • GetTodosByUser(userID)                      │         │
│  │  • CreateTodo(todo)                            │         │
│  │  • DeleteTodo(id, userID)                      │         │
│  │  • ToggleTodo(id)                              │         │
│  │  • Validation & Business Rules                 │         │
│  └────────────────────────────────────────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ Data Operations
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Repository Layer                        │
│  ┌────────────────────────────────────────────────┐         │
│  │         TodoRepository (Data Access)           │         │
│  │  • FindAll()                                   │         │
│  │  • FindByID(id)                                │         │
│  │  • FindByUserID(userID)                        │         │
│  │  • Create(todo)                                │         │
│  │  • Delete(id)                                  │         │
│  │  • Update(todo)                                │         │
│  └────────────────────────────────────────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Data Store (In-Memory)                  │
│                    []Todo (Slice with mutex)                 │
└─────────────────────────────────────────────────────────────┘
```

## 4. Communication Flow

### Flow 1: GET /todos (Get All Todos)
```
Client                Handler              Service              Repository           Data Store
  │                      │                    │                     │                    │
  │──GET /todos?user_id─▶│                    │                     │                    │
  │                      │                    │                     │                    │
  │                      │──GetTodosByUser──▶│                     │                    │
  │                      │                    │                     │                    │
  │                      │                    │──FindByUserID────▶│                    │
  │                      │                    │                     │                    │
  │                      │                    │                     │──Read from slice─▶│
  │                      │                    │                     │◀───return todos────│
  │                      │                    │◀──return todos──────│                    │
  │                      │◀──return todos─────│                     │                    │
  │◀───JSON response─────│                    │                     │                    │
```

### Flow 2: POST /todos (Create Todo)
```
Client                Handler              Service              Repository           Data Store
  │                      │                    │                     │                    │
  │──POST /todos────────▶│                    │                     │                    │
  │  {text, user_id}     │                    │                     │                    │
  │                      │                    │                     │                    │
  │                      │──CreateTodo──────▶│                     │                    │
  │                      │                    │                     │                    │
  │                      │                    │──Validate input     │                    │
  │                      │                    │──Generate ID        │                    │
  │                      │                    │                     │                    │
  │                      │                    │──Create(todo)─────▶│                    │
  │                      │                    │                     │                    │
  │                      │                    │                     │──Append to slice─▶│
  │                      │                    │                     │◀───return todo─────│
  │                      │                    │◀──return todo───────│                    │
  │                      │◀──return todo──────│                     │                    │
  │◀───JSON response─────│                    │                     │                    │
```

### Flow 3: DELETE /todos/{id} (Delete Todo)
```
Client                Handler              Service              Repository           Data Store
  │                      │                    │                     │                    │
  │──DELETE /todos/1────▶│                    │                     │                    │
  │                      │                    │                     │                    │
  │                      │──DeleteTodo(1)───▶│                     │                    │
  │                      │                    │                     │                    │
  │                      │                    │──FindByID(1)──────▶│                    │
  │                      │                    │                     │──Check exists────▶│
  │                      │                    │                     │◀───return todo─────│
  │                      │                    │◀──return todo───────│                    │
  │                      │                    │                     │                    │
  │                      │                    │──Delete(1)────────▶│                    │
  │                      │                    │                     │──Remove from slice│
  │                      │                    │                     │◀───success─────────│
  │                      │                    │◀──success───────────│                    │
  │                      │◀──success──────────│                     │                    │
  │◀───200 OK────────────│                    │                     │                    │
```

## 5. API Endpoints Documentation

**Note:** All API endpoints use standardized response format. See [RESPONSE_FORMAT.md](./RESPONSE_FORMAT.md) for complete details.

### Response Format
All successful responses follow this structure:
```json
{
  "response_code": 200,
  "response_status": "successfully-{type}",
  "message": "Success message",
  "data": {...}
}
```

All error responses follow this structure:
```json
{
  "response_code": 404,
  "response_status": "failed-{type}",
  "message": "Error message",
  "errors": "Error details"
}
```

### 1. GET /users
**Description:** Get all available users

**Response:**
```json
{
  "response_code": 200,
  "response_status": "successfully-get",
  "message": "Data successfully get!",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "email": "jane@example.com",
      "created_at": "0001-01-01T00:00:00Z"
    }
  ]
}
```

### 2. GET /todos
**Description:** Get all todos or filter by user

**Query Parameters:**
- `user_id` (optional): Filter todos by user ID

**Response:**
```json
[
  {
    "id": 1,
    "text": "Buy groceries",
    "completed": false,
    "user_id": 1,
    "created_by": "John Doe",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
]
```

### 3. POST /todos
**Description:** Create a new todo

**Important:** `user_id` must exist in the system (validated)

**Request Body:**
```json
{
  "text": "Buy groceries",
  "user_id": 1,
  "completed": false
}
```

**Validation:**
- `text`: Cannot be empty (returns 400)
- `user_id`: Must be a valid user ID (returns 404 if not found)

**Response:**
```json
{
  "id": 1,
  "text": "Buy groceries",
  "completed": false,
  "user_id": 1,
  "created_by": "John Doe",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

### 4. DELETE /todos/{id}
**Description:** Delete a todo by ID

**URL Parameters:**
- `id`: Todo ID to delete

**Response:**
```json
{
  "message": "Todo deleted successfully"
}
```

## 6. Project Structure

```
collaborative-todo/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── models/
│   │   ├── user.go              # User model
│   │   └── todo.go              # Todo model
│   ├── dto/
│   │   ├── todo_request.go      # Request DTOs
│   │   └── response.go          # Response DTOs
│   ├── repository/
│   │   └── todo_repository.go   # Data access layer
│   ├── service/
│   │   └── todo_service.go      # Business logic layer
│   ├── handler/
│   │   └── todo_handler.go      # HTTP handlers
│   └── middleware/
│       └── cors.go              # CORS middleware
├── go.mod
├── go.sum
├── DESIGN.md                    # This file
└── README.md                    # Setup instructions
```

## 7. Key Design Decisions

### Why Service Layer Pattern?
1. **Separation of Concerns**: Clear separation between HTTP handling, business logic, and data access
2. **Testability**: Each layer can be tested independently
3. **Maintainability**: Easy to modify one layer without affecting others
4. **Scalability**: Easy to add new features or swap implementations

### Concurrency Safety
- Use `sync.RWMutex` for thread-safe access to the in-memory slice
- Multiple goroutines can read simultaneously
- Write operations are properly locked

### Error Handling
- Custom error types for better error messages
- Proper HTTP status codes (200, 201, 400, 404, 500)
- JSON error responses

### CORS Configuration
- Allow all origins for development
- Configurable for production
- Support for preflight requests

## 8. Future Enhancements (Out of Scope)
- Persistent database (PostgreSQL/MySQL)
- Authentication & Authorization
- WebSocket for real-time collaboration
- Todo assignment and sharing
- Tags and categories
- Due dates and reminders
