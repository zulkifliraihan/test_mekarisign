# Response Format Documentation

This application uses a standardized response format similar to Laravel's ReturnResponser trait.

## Response Helper Location

The response helper is located at `internal/helpers/response.go`

## Success Response Format

### Structure
```json
{
  "response_code": 200,
  "response_status": "successfully-{type}",
  "message": "Success message",
  "data": {...},
  "redirect": null
}
```

### Response Types

| Type | Status Code | Default Message | Response Status |
|------|-------------|-----------------|-----------------|
| `Created` | 201 | Data successfully created! | successfully-created |
| `Updated` | 200 | Data successfully updated! | successfully-updated |
| `Deleted` | 200 | Data successfully deleted! | successfully-deleted |
| `Uploaded` | 200 | Data successfully uploaded! | successfully-uploaded |
| `OngoingUpload` | 200 | Data successfully ongoing upload! | successfully-ongoing-upload |
| `Downloaded` | 200 | Data successfully downloaded! | successfully-downloaded |
| `Searched` | 200 | Data successfully searched! | successfully-searched |
| `Get` | 200 | Data successfully get! | successfully-get |

### Usage in Handler

```go
// Import the helpers package
import "test_mekari/internal/helpers"

// Success with default message
helpers.Success(w, helpers.Created, todoData, nil, nil)

// Success with custom message
msg := "Custom success message"
helpers.Success(w, helpers.Updated, todoData, &msg, nil)

// Success with redirect
msg := "Todo created"
redirect := "/todos/1"
helpers.Success(w, helpers.Created, todoData, &msg, &redirect)
```

### Example Success Responses

#### GET Request (List todos)
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

#### POST Request (Create todo)
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

#### DELETE Request
```json
{
  "response_code": 200,
  "response_status": "successfully-deleted",
  "message": "Todo deleted successfully"
}
```

---

## Error Response Format

### Structure
```json
{
  "response_code": 400,
  "response_status": "failed-{type}",
  "message": "Error message",
  "errors": {...}
}
```

### Error Types

| Function | Status Code | Response Status | Default Message |
|----------|-------------|-----------------|-----------------|
| `ErrorValidator` | 422 | failed-validation | Error! The request not expected! |
| `ErrorNotFound` | 404 | failed-not-found | Error! The resource not found! |
| `ErrorAuthentication` | 401 | failed-authentication | Error! The authentication failed! |
| `ErrorServer` | 400 | failed-server | Internal Server Error! |
| `ErrorBadRequest` | 400 | failed-bad-request | Bad Request! |

### Usage in Handler

```go
// Validation error
helpers.ErrorValidator(w, validationErrors, nil)

// Not found error with custom message
msg := "Todo not found"
helpers.ErrorNotFound(w, "todo_id: 123", &msg)

// Server error
msg := "Failed to create todo"
helpers.ErrorServer(w, err.Error(), &msg)

// Bad request
msg := "Invalid user_id parameter"
helpers.ErrorBadRequest(w, err.Error(), &msg)

// Authentication error
msg := "Invalid credentials"
helpers.ErrorAuthentication(w, "Unauthorized access", &msg)
```

### Example Error Responses

#### Validation Error (422)
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "todo text cannot be empty"
}
```

#### Not Found Error (404)
```json
{
  "response_code": 404,
  "response_status": "failed-not-found",
  "message": "Error! The resource not found!",
  "errors": "user_id not found"
}
```

#### Bad Request Error (400)
```json
{
  "response_code": 400,
  "response_status": "failed-bad-request",
  "message": "Invalid user_id parameter",
  "errors": "strconv.Atoi: parsing \"abc\": invalid syntax"
}
```

#### Server Error (400)
```json
{
  "response_code": 400,
  "response_status": "failed-server",
  "message": "Internal Server Error!",
  "errors": "database connection failed"
}
```

---

## Implementation Example

### Handler Example

```go
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    var req dto.CreateTodoRequest

    // Decode request body
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&req); err != nil {
        msg := "Invalid request payload"
        helpers.ErrorValidator(w, err.Error(), &msg)
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
        if err == service.ErrInvalidTodoText {
            helpers.ErrorValidator(w, err.Error(), nil)
            return
        }
        msg := "Failed to create todo"
        helpers.ErrorServer(w, err.Error(), &msg)
        return
    }

    helpers.Success(w, helpers.Created, todo, nil, nil)
}
```

### Try-Catch Pattern (Similar to Laravel)

In Go, we use error handling, similar to try-catch in Laravel:

```go
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    // Get ID from URL
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        msg := "Invalid todo ID"
        helpers.ErrorBadRequest(w, err.Error(), &msg)
        return
    }

    // Decode request
    var req dto.CreateTodoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        msg := "Invalid request payload"
        helpers.ErrorValidator(w, err.Error(), &msg)
        return
    }
    defer r.Body.Close()

    // Update todo - similar to Laravel service call
    todo, err := h.service.UpdateTodo(id, req)
    if err != nil {
        // Handle specific errors
        if err == repository.ErrTodoNotFound {
            helpers.ErrorNotFound(w, err.Error(), nil)
            return
        }
        if err == service.ErrInvalidTodoText {
            helpers.ErrorValidator(w, err.Error(), nil)
            return
        }
        // Generic server error
        msg := "Failed to update todo"
        helpers.ErrorServer(w, err.Error(), &msg)
        return
    }

    // Success response
    helpers.Success(w, helpers.Updated, todo, nil, nil)
}
```

---

## Comparison with Laravel

### Laravel (PHP)
```php
try {
    $service = $this->employeeService->update((int) $id, $request->validated());
    if ($service['status']) {
        return $this->success($service['response'], $service['data']);
    }
    if (($service['response'] ?? null) === 'not-found') {
        return $this->errorNotFound($service['errors'] ?? ['Task not found']);
    }
    return $this->errorServer($service['errors'] ?? ['Unexpected error']);
} catch (\Throwable $th) {
    return $this->errorServer($th->getMessage());
}
```

### Go (Equivalent)
```go
todo, err := h.service.UpdateTodo(id, req)
if err != nil {
    if err == repository.ErrTodoNotFound {
        helpers.ErrorNotFound(w, err.Error(), nil)
        return
    }
    msg := "Failed to update todo"
    helpers.ErrorServer(w, err.Error(), &msg)
    return
}

helpers.Success(w, helpers.Updated, todo, nil, nil)
```

---

## Benefits

1. **Consistent Response Format**: All API responses follow the same structure
2. **Easy to Maintain**: Centralized response handling
3. **Type-Safe**: Uses constants for response types
4. **Flexible**: Custom messages supported per response
5. **Similar to Laravel**: Familiar pattern for Laravel developers
6. **Automatic Logging**: Server errors are logged automatically

---

## Available Helper Functions

### Success Helpers
- `helpers.Success(w, responseType, data, message, redirect)`

### Error Helpers
- `helpers.ErrorValidator(w, errors, message)` - 422 Validation Error
- `helpers.ErrorNotFound(w, errors, message)` - 404 Not Found
- `helpers.ErrorAuthentication(w, errors, message)` - 401 Unauthorized
- `helpers.ErrorServer(w, errors, message)` - 400 Server Error
- `helpers.ErrorBadRequest(w, errors, message)` - 400 Bad Request

### Parameters
- `w http.ResponseWriter` - HTTP response writer
- `responseType helpers.ResponseType` - Type of success response (Created, Updated, etc)
- `data interface{}` - Response data (can be nil)
- `errors interface{}` - Error details (can be nil)
- `message *string` - Custom message (can be nil for default message)
- `redirect *string` - Redirect URL (can be nil)

---

## Notes

- All server errors are automatically logged using `log.Printf`
- Default messages are used if a custom message is not provided
- Response status follows the pattern: `successfully-{type}` for success and `failed-{type}` for errors
- This format makes it easier for the frontend to handle responses consistently
