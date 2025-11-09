# Collaborative Todo List API

A collaborative team todo list application built with Go and Gorilla Mux following the Service Layer pattern.

## Features

- Create, read, update, and delete todos
- Collaborative team functionality - todos are associated with users
- Filter todos by user
- In-memory storage with thread-safe operations
- RESTful API design
- CORS enabled for frontend integration
- Clean architecture with service layer pattern

## Tech Stack

- **Language**: Go 1.21
- **Router**: Gorilla Mux
- **Architecture**: Service Layer Pattern
- **Storage**: In-memory (slice with sync.RWMutex)

## AI-Assisted Optimizations

This project was developed with AI assistance for code optimization and documentation:

**AI Contributions:**
- ✅ **Code Structure Optimization** - Separated models, DTOs, and routes for clean architecture
- ✅ **Error Handling Enhancement** - Human-readable error messages instead of technical errors
- ✅ **Development Experience** - Hot reload setup, environment variables support
- ✅ **Data Management** - User seeder implementation
- ✅ **Frontend Development** - Responsive web interface with Tailwind CSS
- ✅ **Comprehensive Documentation** - 7 detailed .md files covering all aspects

**Human Developer:**
- Core business logic and Service Layer Pattern implementation
- API endpoint design and requirements
- Project architecture decisions

**See [ACTION_PLAN.md](./ACTION_PLAN.md) for complete AI optimization history and all prompts used.**

## Project Structure

```
test_mekari/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point (clean, only initialization)
├── internal/
│   ├── models/
│   │   ├── user.go              # User model
│   │   └── todo.go              # Todo model
│   ├── dto/
│   │   ├── todo_request.go      # Request DTOs
│   │   └── response.go          # Response DTOs (deprecated)
│   ├── helpers/
│   │   └── response.go          # Standardized response helper
│   ├── routes/
│   │   └── routes.go            # All route definitions
│   ├── repository/
│   │   ├── todo_repository.go   # Data access layer
│   │   └── user_seeder.go       # User data seeder
│   ├── service/
│   │   └── todo_service.go      # Business logic layer
│   ├── handler/
│   │   └── todo_handler.go      # HTTP handlers
│   └── middleware/
│       └── cors.go              # CORS & logging middleware
├── go.mod
├── go.sum
├── ACTION_PLAN.md               # AI optimization & documentation history
├── DESIGN.md                    # Design documentation & architecture
├── RESPONSE_FORMAT.md           # Response format documentation
├── ERROR_HANDLING.md            # Human-readable error messages documentation
├── SEEDER.md                    # Database seeder documentation
├── FRONTEND.md                  # Frontend documentation
└── README.md                    # This file
```

## Architecture

The application follows the **Service Layer Pattern** with clear separation of concerns:

```
Client → Handler → Service → Repository → Data Store
```

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic and validation
- **Repository Layer**: Data access and manipulation
- **Data Store**: In-memory slice with mutex for thread-safety

See [DESIGN.md](./DESIGN.md) for detailed architecture diagrams and communication flows.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git (optional)

### Installation

1. Clone or download the project:
```bash
cd test_mekari
```

2. Install dependencies:
```bash
go mod download
```

3. Install Air for hot reload (optional, like npm run dev):
```bash
go install github.com/air-verse/air@latest
```

4. Run the application:

**Option A: With Hot Reload (Recommended for Development)**
```bash
make dev
# or
air
```
The server will auto-restart on code changes (like npm run dev in Node.js)

**Option B: Standard Run**
```bash
make run
# or
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Alternative: Build and Run

```bash
# Build the binary
make build
# or
go build -o todo-api cmd/api/main.go

# Run the binary
./todo-api
```

### Development with Hot Reload

For development, use `make dev` or `air` command. The server will automatically restart when you make changes to any `.go` file:

```bash
make dev
```

Features:
- ✅ Auto-restart on file changes
- ✅ Fast rebuild
- ✅ Colored output
- ✅ Build error logs

### Environment Variables (.env)

This application supports a `.env` file for configuration.

1. Copy `.env.example` to `.env`:
```bash
cp .env.example .env
```

2. Edit `.env` as needed:
```bash
# .env
APP_PORT=8081
```

3. Run the application (it will automatically load .env):
```bash
make dev
# or
go run cmd/api/main.go
```

**Available Environment Variables:**
- `APP_PORT` - Server port (default: 8080)

### Custom Port

**Option 1: Using .env file (Recommended)**
```bash
# Edit .env file
APP_PORT=3000

# Run
make dev
```

**Option 2: Using environment variable directly**
```bash
APP_PORT=3000 go run cmd/api/main.go
# or
APP_PORT=3000 make run-port
```

## API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Get All Users

**Endpoint:** `GET /users`

**Description:** Retrieve all available users

**Example Request:**
```bash
curl http://localhost:8080/users
```

**Example Response:**
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
    },
    {
      "id": 3,
      "name": "Bob Johnson",
      "email": "bob@example.com",
      "created_at": "0001-01-01T00:00:00Z"
    }
  ]
}
```

#### 2. Get All Todos

**Endpoint:** `GET /todos`

**Description:** Retrieve all todos or filter by user

**Query Parameters:**
- `user_id` (optional): Filter todos by user ID

**Example Request:**
```bash
# Get all todos
curl http://localhost:8080/todos

# Get todos for user 1
curl http://localhost:8080/todos?user_id=1
```

**Example Response:**
```json
{
  "response_code": 200,
  "response_status": "successfully-get",
  "message": "Data successfully get!",
  "data": [
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
}
```

#### 3. Create Todo

**Endpoint:** `POST /todos`

**Description:** Create a new todo

**Important:** `user_id` must exist in the system. Use `GET /users` to see available users.

**Request Body:**
```json
{
  "text": "Buy groceries",
  "user_id": 1,
  "completed": false
}
```

**Validation:**
- `text`: Cannot be empty
- `user_id`: Must be a valid user ID (returns 404 if user not found)

**Example Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Buy groceries",
    "user_id": 1,
    "completed": false
  }'
```

**Example Response:**
```json
{
  "response_code": 201,
  "response_status": "successfully-created",
  "message": "Data successfully created!",
  "data": {
    "id": 1,
    "text": "Buy groceries",
    "completed": false,
    "user_id": 1,
    "created_by": "John Doe",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

**Error Response (User not found):**
```json
{
  "response_code": 404,
  "response_status": "failed-not-found",
  "message": "Error! The resource not found!",
  "errors": "user_id not found"
}
```

#### 4. Delete Todo

**Endpoint:** `DELETE /todos/{id}`

**Description:** Delete a todo by ID

**URL Parameters:**
- `id`: Todo ID to delete

**Example Request:**
```bash
curl -X DELETE http://localhost:8080/todos/1
```

**Example Response:**
```json
{
  "response_code": 200,
  "response_status": "successfully-deleted",
  "message": "Todo deleted successfully"
}
```

#### 5. Update Todo

**Endpoint:** `PUT /todos/{id}`

**Description:** Update a todo

**Request Body:**
```json
{
  "text": "Buy groceries and cook dinner",
  "user_id": 1,
  "completed": true
}
```

**Example Request:**
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Buy groceries and cook dinner",
    "user_id": 1,
    "completed": true
  }'
```

#### 6. Toggle Todo Status

**Endpoint:** `PATCH /todos/{id}/toggle`

**Description:** Toggle the completed status of a todo

**Example Request:**
```bash
curl -X PATCH http://localhost:8080/todos/1/toggle
```

#### 7. Health Check

**Endpoint:** `GET /health`

**Description:** Check if the API is running

**Example Request:**
```bash
curl http://localhost:8080/health
```

**Example Response:**
```json
{
  "response_code": 200,
  "response_status": "successfully-get",
  "message": "Service is running",
  "data": {
    "status": "healthy",
    "service": "Collaborative Todo List API"
  }
}
```

#### 8. API Information

**Endpoint:** `GET /`

**Description:** Get API information and available endpoints

**Example Request:**
```bash
curl http://localhost:8080/
```

**Example Response:**
```json
{
  "response_code": 200,
  "response_status": "successfully-get",
  "message": "Welcome to Collaborative Todo List API",
  "data": {
    "name": "Collaborative Todo List API",
    "version": "1.0.0",
    "endpoints": {
      "GET /users": "Get all users",
      "GET /todos": "Get all todos (optional: ?user_id=1 to filter by user)",
      "POST /todos": "Create a new todo (user_id must exist)",
      "DELETE /todos/{id}": "Delete a todo",
      "PUT /todos/{id}": "Update a todo",
      "PATCH /todos/{id}/toggle": "Toggle todo completed status",
      "GET /health": "Health check"
    }
  }
}
```

### Error Responses

All error responses follow this standardized format:

```json
{
  "response_code": 404,
  "response_status": "failed-not-found",
  "message": "Error! The resource not found!",
  "errors": "Error details here"
}
```

**Error Types:**

| Status Code | Response Status | Default Message |
|-------------|-----------------|-----------------|
| 422 | failed-validation | Error! The request not expected! |
| 404 | failed-not-found | Error! The resource not found! |
| 401 | failed-authentication | Error! The authentication failed! |
| 400 | failed-server | Internal Server Error! |
| 400 | failed-bad-request | Bad Request! |

**Human-Readable Error Messages:**

This API returns user-friendly error messages instead of technical errors. For example:

❌ **Before:** `"json: cannot unmarshal string into Go struct field CreateTodoRequest.completed of type bool"`

✅ **After:** `"Field 'completed' must be a boolean value (true or false)"`

**See [ERROR_HANDLING.md](./ERROR_HANDLING.md) for complete error handling documentation.**

**See [RESPONSE_FORMAT.md](./RESPONSE_FORMAT.md) for complete response format documentation.**

**HTTP Status Codes:**
- `200 OK`: Success
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid input / Server error
- `404 Not Found`: Resource not found
- `422 Unprocessable Entity`: Validation error

## Sample Users

The application comes with 3 pre-configured sample users created via **seeder** on startup:

| User ID | Name | Email |
|---------|------|-------|
| 1 | John Doe | john@example.com |
| 2 | Jane Smith | jane@example.com |
| 3 | Bob Johnson | bob@example.com |

**See [SEEDER.md](./SEEDER.md) for seeder documentation and how to add more users.**

## Testing the API

### Using cURL

```bash
# Create a todo for John Doe
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Review code", "user_id": 1, "completed": false}'

# Create a todo for Jane Smith
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Write documentation", "user_id": 2, "completed": false}'

# Get all todos
curl http://localhost:8080/todos

# Get todos for user 1
curl http://localhost:8080/todos?user_id=1

# Toggle todo completion
curl -X PATCH http://localhost:8080/todos/1/toggle

# Delete a todo
curl -X DELETE http://localhost:8080/todos/1
```

### Using Postman

1. Import the endpoints into Postman
2. Set `Content-Type: application/json` header for POST/PUT requests
3. Use the examples above as request bodies

## Development

### Running Tests

```bash
go test ./...
```

### Code Structure

The codebase follows Go best practices:

- **Package organization**: Clear separation by feature and layer
- **Dependency injection**: Services and repositories are injected
- **Interface-driven**: Easy to mock and test
- **Thread-safe**: Uses sync.RWMutex for concurrent access
- **Error handling**: Proper error propagation and HTTP status codes

### Adding New Features

To add a new feature:

1. Add model to `internal/models/`
2. Add repository methods to `internal/repository/`
3. Add service logic to `internal/service/`
4. Add HTTP handlers to `internal/handler/`
5. Register routes in `cmd/api/main.go`

## Design Decisions

### Why Service Layer Pattern?

1. **Separation of Concerns**: Each layer has a single responsibility
2. **Testability**: Easy to unit test each layer independently
3. **Maintainability**: Changes in one layer don't affect others
4. **Scalability**: Easy to add new features or swap implementations

### Thread Safety

- Uses `sync.RWMutex` for safe concurrent access
- Multiple readers can access data simultaneously
- Writers have exclusive access during modifications

### In-Memory Storage

- Fast performance for demo/testing purposes
- Easy to understand and debug
- Can be replaced with a database later by only changing the repository layer

## Limitations

- **No Persistence**: Data is lost when the server restarts
- **No Authentication**: All endpoints are publicly accessible
- **Limited Validation**: Basic validation only
- **No Real-time Updates**: No WebSocket support for live collaboration

## Future Enhancements

- Add persistent database (PostgreSQL/MySQL)
- Implement authentication and authorization
- Add WebSocket for real-time collaboration
- Implement todo sharing and assignment
- Add tags, categories, and priorities
- Add due dates and reminders
- Implement pagination for large datasets
- Add unit and integration tests
- Add API documentation with Swagger

## Credits

### Development
- **Core Development**: Manual implementation of Service Layer Pattern, business logic, and API design
- **AI Optimization**: Code refactoring, structure improvements, and error handling enhancement
- **AI Documentation**: Comprehensive .md files and guides (ACTION_PLAN.md, DESIGN.md, etc.)
- **Frontend**: AI-generated responsive web interface with Tailwind CSS

### AI Tools Used
- **Model**: Claude Sonnet 4.5 (Anthropic)
- **Purpose**: Code optimization, documentation, and frontend development
- **Approach**: Collaborative - AI assisted human-written codebase

**For detailed AI contribution history, see [ACTION_PLAN.md](./ACTION_PLAN.md)**

## License

This project is created for technical assessment purposes.

## Author

Technical Test - Collaborative Todo List Application
