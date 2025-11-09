# Quick Start Guide

## Project Structure

```
test_mekari/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── models/
│   │   └── todo.go                    # Data models (Todo, User, DTOs)
│   ├── repository/
│   │   └── todo_repository.go         # Data access layer
│   ├── service/
│   │   └── todo_service.go            # Business logic layer
│   ├── handler/
│   │   └── todo_handler.go            # HTTP handlers
│   └── middleware/
│       └── cors.go                    # CORS middleware
├── DESIGN.md                          # Design & Architecture documentation
├── README.md                          # Full documentation
├── Makefile                           # Build automation
├── test_api.sh                        # API test script
├── go.mod
└── go.sum
```

## How to Run the Application

### Option 1: Hot Reload Development (Recommended - Like npm run dev)

```bash
# Install Air (once)
go install github.com/air-verse/air@latest

# Run with hot reload (auto-restart on code changes)
make dev
# or
air
```

**Benefits:**
- ✅ Automatically restarts on code changes
- ✅ Similar to `npm run dev` in Node.js
- ✅ No manual stop/start needed
- ✅ Fast rebuilds

### Option 2: Using Makefile

```bash
# List all available commands
make help

# Install dependencies
make deps

# Run the application (without hot reload)
make run

# Build the application
make build

# Run tests (server must be running in another terminal)
make test

# Clean build artifacts
make clean
```

### Option 3: Using Go Command

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/api/main.go

# Build the application
go build -o todo-api cmd/api/main.go

# Run binary
./todo-api
```

### Option 4: Custom Port

```bash
# With Makefile
PORT=3000 make run-port

# Or with Go
PORT=3000 go run cmd/api/main.go
```

## Testing API

### 1. Manual Testing with cURL

```bash
# Health check
curl http://localhost:8080/health

# List all users
curl http://localhost:8080/users

# Buat todo (user_id must exist!)
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Buy groceries", "user_id": 1, "completed": false}'

# List all todos
curl http://localhost:8080/todos

# Filter todos by user
curl http://localhost:8080/todos?user_id=1

# Toggle todo completion
curl -X PATCH http://localhost:8080/todos/1/toggle

# Delete todo
curl -X DELETE http://localhost:8080/todos/1
```

### 2. Automated Testing

```bash
# Terminal 1: Start server
make run

# Terminal 2: Run tests
make test

# Atau langsung
./test_api.sh
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users |
| GET | `/todos` | Get all todos (optional: ?user_id=X) |
| POST | `/todos` | Create new todo (user_id must exist) |
| DELETE | `/todos/{id}` | Delete todo |
| PUT | `/todos/{id}` | Update todo |
| PATCH | `/todos/{id}/toggle` | Toggle completed status |
| GET | `/health` | Health check |

## Sample Users

| ID | Name | Email |
|----|------|-------|
| 1 | John Doe | john@example.com |
| 2 | Jane Smith | jane@example.com |
| 3 | Bob Johnson | bob@example.com |

## Architecture - Service Layer Pattern

```
┌──────────────┐
│    Client    │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│   Handler    │  ← HTTP Request/Response
└──────┬───────┘
       │
       ▼
┌──────────────┐
│   Service    │  ← Business Logic & Validation
└──────┬───────┘
       │
       ▼
┌──────────────┐
│  Repository  │  ← Data Access
└──────┬───────┘
       │
       ▼
┌──────────────┐
│  Data Store  │  ← In-Memory (Slice + Mutex)
└──────────────┘
```

## Key Features

✅ **Service Layer Pattern** - Clean separation of concerns
✅ **Thread-Safe** - sync.RWMutex for concurrent access
✅ **CORS Enabled** - Support for frontend integration
✅ **Error Handling** - Proper HTTP status codes and error messages
✅ **RESTful API** - Following REST best practices
✅ **Validation** - Input validation in the service layer
✅ **Documentation** - Comprehensive docs and examples

## Troubleshooting

### Port already in use
```bash
# Use another port
PORT=3000 go run cmd/api/main.go
```

### Dependencies issues
```bash
go mod tidy
go mod download
```

### Build errors
```bash
# Clean and rebuild
make clean
make build
```

## Next Steps

1. Read [DESIGN.md](./DESIGN.md) to understand the detailed architecture
2. Read [README.md](./README.md) for complete API documentation
3. Run `./test_api.sh` to see all endpoints in action
4. Modify the code as needed

## Important Notes

- Data is stored in **memory** (not persistent)
- CORS enabled for **all origins** (development mode)
- Server runs on **port 8080** by default
- 3 sample users are available for testing (ID: 1, 2, 3)

---

**Happy Coding!** 🚀
