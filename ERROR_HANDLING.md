# Error Handling Documentation

## Overview

This application implements comprehensive error handling with **human-readable error messages** to improve developer experience when integrating with the API.

## Human-Readable Error Messages

Instead of showing technical JSON decode errors like:
```
"json: cannot unmarshal string into Go struct field CreateTodoRequest.completed of type bool"
```

The API returns user-friendly messages like:
```
"Field 'completed' must be a boolean value (true or false)"
```

## Implementation

### ParseJSONError Helper

Located in `internal/helpers/response.go`, the `ParseJSONError()` function converts technical errors into human-readable messages.

```go
func ParseJSONError(err error) string {
    // Analyzes error message and returns human-friendly text
}
```

### Usage in Handlers

Both `CreateTodo` and `UpdateTodo` handlers use this helper:

```go
decoder := json.NewDecoder(r.Body)
if err := decoder.Decode(&req); err != nil {
    // Parse JSON error to human-readable message
    humanMsg := helpers.ParseJSONError(err)
    helpers.ErrorValidator(w, humanMsg, nil)
    return
}
```

## Error Types and Messages

### 1. Type Mismatch Errors

#### Boolean Type Error
**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": 1, "completed": "yes"}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Field 'completed' must be a boolean value (true or false)"
}
```

#### Number Type Error
**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": "one", "completed": false}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Field 'user_id' must be a number"
}
```

#### String Type Error
**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": 123, "user_id": 1, "completed": false}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Field 'text' must be a string"
}
```

### 2. Invalid JSON Syntax

**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", invalid json}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Invalid JSON format. Please check your request payload"
}
```

### 3. Empty Request Body

**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d ''
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Empty request body. Please provide valid JSON data"
}
```

### 4. Unknown Fields

**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": 1, "unknown_field": "value"}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Request contains unknown field(s)"
}
```

### 5. Missing Required Fields

**Request:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test"}'
```

**Response:**
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Error! The request not expected!",
  "errors": "Required field is missing"
}
```

## Error Detection Logic

The `ParseJSONError` function uses pattern matching to detect error types:

### Type Mismatch Detection
```go
if strings.Contains(errMsg, "cannot unmarshal") {
    if strings.Contains(errMsg, "bool") {
        if strings.Contains(errMsg, "completed") {
            return "Field 'completed' must be a boolean value (true or false)"
        }
    }
    // ... more patterns
}
```

### JSON Syntax Detection
```go
if strings.Contains(errMsg, "invalid character") {
    return "Invalid JSON format. Please check your request payload"
}
```

### Empty Body Detection
```go
if strings.Contains(errMsg, "EOF") {
    return "Empty request body. Please provide valid JSON data"
}
```

## Benefits

### 1. Better Developer Experience
- Errors are immediately understandable
- No need to decode technical Go error messages
- Faster debugging and integration

### 2. Improved API Documentation
- Error examples are clear and self-explanatory
- Easy to write client-side validation based on error messages
- Better error handling in frontend applications

### 3. Consistent Error Format
- All validation errors return 422 status code
- Consistent response structure
- Predictable error handling

## Extending Error Handling

### Adding New Error Patterns

To add a new error pattern, edit `internal/helpers/response.go`:

```go
func ParseJSONError(err error) string {
    // ... existing code ...

    // Add new pattern
    if strings.Contains(errMsg, "your_pattern") {
        return "Your human-readable message"
    }

    // ... rest of code ...
}
```

### Example: Add Array Type Error

```go
if strings.Contains(errMsg, "array") || strings.Contains(errMsg, "slice") {
    return "Field must be an array"
}
```

### Example: Add Object Type Error

```go
if strings.Contains(errMsg, "object") || strings.Contains(errMsg, "struct") {
    return "Field must be an object"
}
```

## Testing

### Manual Testing Script

Create a test script `test_errors.sh`:

```bash
#!/bin/bash

echo "Test 1: Boolean type error"
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": 1, "completed": "yes"}'

echo "\n\nTest 2: Number type error"
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": "one", "completed": false}'

echo "\n\nTest 3: Invalid JSON"
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", invalid}'

echo "\n\nTest 4: Empty body"
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d ''
```

Run with:
```bash
chmod +x test_errors.sh
./test_errors.sh
```

### Automated Testing

Example unit test (to be implemented):

```go
func TestParseJSONError(t *testing.T) {
    tests := []struct {
        name     string
        error    error
        expected string
    }{
        {
            name:     "Boolean type mismatch",
            error:    errors.New("json: cannot unmarshal string into Go struct field CreateTodoRequest.completed of type bool"),
            expected: "Field 'completed' must be a boolean value (true or false)",
        },
        {
            name:     "Number type mismatch",
            error:    errors.New("json: cannot unmarshal string into Go struct field CreateTodoRequest.user_id of type int"),
            expected: "Field 'user_id' must be a number",
        },
        // ... more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := helpers.ParseJSONError(tt.error)
            if result != tt.expected {
                t.Errorf("Expected %s, got %s", tt.expected, result)
            }
        })
    }
}
```

## Best Practices

### 1. Always Use ParseJSONError
When decoding JSON in handlers:
```go
// ✅ Good
if err := decoder.Decode(&req); err != nil {
    humanMsg := helpers.ParseJSONError(err)
    helpers.ErrorValidator(w, humanMsg, nil)
    return
}

// ❌ Bad
if err := decoder.Decode(&req); err != nil {
    helpers.ErrorValidator(w, err.Error(), nil)
    return
}
```

### 2. Keep Error Messages Actionable
Error messages should tell users:
- What went wrong
- What type is expected
- Example of correct format (when possible)

### 3. Log Technical Errors
While returning human-readable errors to users, still log technical details:
```go
if err := decoder.Decode(&req); err != nil {
    log.Printf("JSON decode error: %v", err) // Log technical details
    humanMsg := helpers.ParseJSONError(err)
    helpers.ErrorValidator(w, humanMsg, nil)
    return
}
```

## Future Enhancements

1. **Field-specific validation rules**
   - Min/max length for strings
   - Range validation for numbers
   - Email format validation

2. **Multiple error aggregation**
   - Return all validation errors at once
   - Instead of stopping at first error

3. **Internationalization (i18n)**
   - Support multiple languages
   - Accept-Language header support

4. **Custom error codes**
   - Unique error codes for each type
   - Easier for frontend error handling

5. **Structured error responses**
   - Include field name in structured format
   - Machine-readable error types

Example future response:
```json
{
  "response_code": 422,
  "response_status": "failed-validation",
  "message": "Validation failed",
  "errors": [
    {
      "field": "completed",
      "type": "type_mismatch",
      "expected": "boolean",
      "received": "string",
      "message": "Field 'completed' must be a boolean value (true or false)"
    }
  ]
}
```

## Summary

The human-readable error handling feature:
- ✅ Converts technical errors to friendly messages
- ✅ Works on all POST and PUT endpoints
- ✅ Handles multiple error scenarios
- ✅ Maintains consistent response format
- ✅ Easy to extend and customize
- ✅ Improves developer experience

For more information, see:
- `internal/helpers/response.go` - Implementation
- `internal/handler/todo_handler.go` - Usage examples
- `RESPONSE_FORMAT.md` - Response format documentation
