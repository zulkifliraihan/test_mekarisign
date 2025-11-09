# Technical Test - Collaborative Todo List Application

## Submission Summary

The **Collaborative Todo List** application has been built using:
- **Language**: Go 1.21
- **Router**: Gorilla Mux
- **Architecture**: Service Layer Pattern
- **Storage**: In-memory with thread-safe operations

---

## ✅ Requirements Checklist

### Product Requirements
- ✅ Collaborative team todo list functionality
- ✅ Users can collaborate through the board
- ✅ Every user can filter the board based on user ID

### API Endpoints
- ✅ `GET /todos` - Return all todos as JSON (with optional user filter)
- ✅ `POST /todos` - Add a new todo
- ✅ `DELETE /todos/{id}` - Delete a todo by ID

### Additional Features Implemented
- ✅ `PUT /todos/{id}` - Update a todo
- ✅ `PATCH /todos/{id}/toggle` - Toggle completed status
- ✅ `GET /health` - Health check endpoint

### Technical Requirements
- ✅ Store todos in a simple slice (in-memory)
- ✅ Unique ID for each todo (using counter)
- ✅ Error handling (404 for not found, 400 for bad request, etc.)
- ✅ CORS enabled for frontend communication
- ✅ Service Layer Pattern implementation
- ✅ Thread-safe operations using sync.RWMutex

### Design Documentation
- ✅ **ERD** - Database/Data model diagram
- ✅ **Architecture Diagram** - Service layer pattern visualization
- ✅ **Communication Flow** - Request/response flow diagrams
- ✅ **Complete documentation** for junior developers

---

## 📁 Project Structure

```
test_mekari/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── models/
│   │   ├── user.go                    # User model
│   │   └── todo.go                    # Todo model
│   ├── dto/
│   │   ├── todo_request.go            # Request DTOs
│   │   └── response.go                # Response DTOs
│   ├── repository/
│   │   └── todo_repository.go         # Data access layer (thread-safe)
│   ├── service/
│   │   └── todo_service.go            # Business logic & validation
│   ├── handler/
│   │   └── todo_handler.go            # HTTP handlers
│   └── middleware/
│       └── cors.go                    # CORS & logging middleware
├── DESIGN.md                          # ⭐ Design documentation & diagrams
├── README.md                          # Complete API documentation
├── QUICKSTART.md                      # Quick start guide
├── TECHNICAL_TEST.md                  # This file
├── Makefile                           # Build automation
├── test_api.sh                        # Automated API testing script
├── go.mod
└── go.sum
```

---

## 🚀 How to Run

### Quick Start (3 steps)

```bash
# 1. Navigate to project directory
cd /Volumes/ssd_projects/test_mekari

# 2. Install dependencies
go mod download

# 3. Run the application
go run cmd/api/main.go
```

The server will run at: **http://localhost:8080**

### Alternative: Using Makefile

```bash
make run        # Run application
make build      # Build binary
make test       # Run tests (requires server running)
```

---

## 🧪 Testing the Application

### Option 1: Automated Test Script

```bash
# Terminal 1: Start server
go run cmd/api/main.go

# Terminal 2: Run tests
chmod +x test_api.sh
./test_api.sh
```

The test script runs 16 test cases covering all endpoints.

### Option 2: Manual Testing with cURL

```bash
# Get all users
curl http://localhost:8080/users

# Create a todo (user_id must exist!)
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Review code", "user_id": 1, "completed": false}'

# Get all todos
curl http://localhost:8080/todos

# Get todos for specific user
curl http://localhost:8080/todos?user_id=1

# Delete a todo
curl -X DELETE http://localhost:8080/todos/1
```

---

## 📋 API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/users` | Get all users | - |
| GET | `/todos` | Get all todos (optional: ?user_id=X) | - |
| POST | `/todos` | Create new todo (user_id must exist) | `{"text": "...", "user_id": 1, "completed": false}` |
| DELETE | `/todos/{id}` | Delete todo | - |
| PUT | `/todos/{id}` | Update todo | `{"text": "...", "user_id": 1, "completed": true}` |
| PATCH | `/todos/{id}/toggle` | Toggle completed | - |
| GET | `/health` | Health check | - |

### Example Response

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

---

## 🏗️ Architecture - Service Layer Pattern

```
┌─────────────────────────────────────────┐
│              Client/Browser              │
└───────────────┬─────────────────────────┘
                │ HTTP/JSON
                ▼
┌─────────────────────────────────────────┐
│          Handler Layer (HTTP)           │
│  • Request validation                   │
│  • Response formatting                  │
│  • HTTP status codes                    │
└───────────────┬─────────────────────────┘
                │
                ▼
┌─────────────────────────────────────────┐
│         Service Layer (Logic)           │
│  • Business logic                       │
│  • Input validation                     │
│  • Error handling                       │
└───────────────┬─────────────────────────┘
                │
                ▼
┌─────────────────────────────────────────┐
│       Repository Layer (Data)           │
│  • Data access                          │
│  • CRUD operations                      │
│  • Thread-safe operations               │
└───────────────┬─────────────────────────┘
                │
                ▼
┌─────────────────────────────────────────┐
│      Data Store (In-Memory Slice)       │
│  • sync.RWMutex for concurrency         │
└─────────────────────────────────────────┘
```

**Benefits:**
- Clear separation of concerns
- Easy to test each layer independently
- Maintainable and scalable
- Can easily swap storage implementation

---

## 🎯 Key Features

### 1. Thread-Safe Operations
```go
// Repository uses sync.RWMutex
type TodoRepository struct {
    todos   []models.Todo
    mu      sync.RWMutex
}
```

### 2. CORS Support
```go
// CORS middleware for frontend integration
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
```

### 3. Validation
```go
// Service layer performs input validation
if strings.TrimSpace(req.Text) == "" {
    return ErrInvalidTodoText
}
```

### 4. Error Handling
```go
// Proper HTTP status codes
if err == repository.ErrTodoNotFound {
    statusCode = http.StatusNotFound
}
```

---

## 📖 Documentation Files

1. **DESIGN.md** - Design documentation including:
   - ERD (Entity Relationship Diagram)
   - Architecture diagram
   - Communication flow diagrams
   - API specifications

2. **README.md** - Complete documentation:
   - Setup instructions
   - API documentation with examples
   - Architecture explanation
   - Development guide

3. **QUICKSTART.md** - Quick reference:
   - How to run
   - How to test
   - Quick examples

---

## 🔍 Sample Data

The application includes 3 sample users for testing:

| User ID | Name | Email |
|---------|------|-------|
| 1 | John Doe | john@example.com |
| 2 | Jane Smith | jane@example.com |
| 3 | Bob Johnson | bob@example.com |

---

## 💡 Design Decisions

### Why Service Layer Pattern?
- **Separation of Concerns**: Each layer has a clear responsibility
- **Testability**: Easy to unit test each layer
- **Maintainability**: Changes in one layer do not affect others
- **Scalability**: Easy to add new features

### Why In-Memory Storage?
- Matches requirements (no database needed)
- Fast for demo/testing
- Thread-safe with sync.RWMutex
- Easy to replace with a database later (only change the repository layer)

### Additional Enhancements Beyond Requirements
- GET /users endpoint to list all users
- User validation on todo creation (user_id must exist)
- Health check endpoint
- Toggle completion endpoint
- Update todo endpoint
- Comprehensive error handling with proper status codes
- Automated test script
- Build automation (Makefile)
- Complete documentation

---

## 📊 Test Coverage

The test script (`test_api.sh`) covers:
1. ✅ Health check
2. ✅ Get all users
3. ✅ Create todos
4. ✅ Get all todos
5. ✅ Filter todos by user
6. ✅ Update todos
7. ✅ Toggle completion
8. ✅ Delete todos
9. ✅ Error cases (user not found, invalid input)

---

## 🎓 For Junior Developers

The documentation is comprehensive so junior developers can:
- Understand the architecture pattern used
- See the communication flow between layers
- Understand separation of concerns
- Learn best practices (error handling, validation, concurrency)
- Easily extend the application with new features

See **DESIGN.md** for diagrams and complete explanations!

---

## ⚡ Performance

- Thread-safe for concurrent requests
- Fast in-memory operations
- Efficient mutex usage (RWMutex for multiple readers)

---

## 🚀 Production Considerations (Future)

Current implementation is for demo/testing. For production:
- [ ] Add persistent database (PostgreSQL/MySQL)
- [ ] Add authentication & authorization
- [ ] Add request validation middleware
- [ ] Add rate limiting
- [ ] Add logging framework
- [ ] Add metrics/monitoring
- [ ] Add unit & integration tests
- [ ] Add Docker support
- [ ] Add API documentation (Swagger/OpenAPI)

---

## 📝 Notes

- ✅ All requirements fulfilled
- ✅ Service layer pattern implemented
- ✅ Complete documentation provided
- ✅ Ready for demonstration
- ✅ Clean, maintainable code
- ✅ Following Go best practices

---

**Thank you for reviewing this submission!**

For questions or clarifications, please refer to the documentation files or run the test script to see the application in action.
