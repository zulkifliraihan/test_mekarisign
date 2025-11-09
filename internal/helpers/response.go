package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	ResponseCode   int         `json:"response_code"`
	ResponseStatus string      `json:"response_status"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
	Redirect       *string     `json:"redirect,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	ResponseCode   int         `json:"response_code"`
	ResponseStatus string      `json:"response_status"`
	Message        string      `json:"message"`
	Errors         interface{} `json:"errors,omitempty"`
}

// ResponseType represents different types of success responses
type ResponseType string

const (
	Created        ResponseType = "created"
	Updated        ResponseType = "updated"
	Deleted        ResponseType = "deleted"
	Uploaded       ResponseType = "uploaded"
	OngoingUpload  ResponseType = "ongoing-upload"
	Downloaded     ResponseType = "downloaded"
	Searched       ResponseType = "searched"
	Get            ResponseType = "get"
)

// responseFormat holds the status code and default message for each response type
var responseFormats = map[ResponseType]struct {
	Code    int
	Message string
}{
	Created:       {201, "Data successfully created!"},
	Updated:       {200, "Data successfully updated!"},
	Deleted:       {200, "Data successfully deleted!"},
	Uploaded:      {200, "Data successfully uploaded!"},
	OngoingUpload: {200, "Data successfully ongoing upload!"},
	Downloaded:    {200, "Data successfully downloaded!"},
	Searched:      {200, "Data successfully searched!"},
	Get:           {200, "Data successfully get!"},
}

// Success returns a success JSON response
func Success(w http.ResponseWriter, responseType ResponseType, data interface{}, message *string, redirect *string) {
	format, exists := responseFormats[responseType]
	if !exists {
		format = struct {
			Code    int
			Message string
		}{200, "Successfully Action!"}
	}

	// Use custom message if provided, otherwise use default
	finalMessage := format.Message
	if message != nil {
		finalMessage = *message
	}

	response := SuccessResponse{
		ResponseCode:   format.Code,
		ResponseStatus: "successfully-" + string(responseType),
		Message:        finalMessage,
		Data:           data,
		Redirect:       redirect,
	}

	writeJSON(w, format.Code, response)
}

// ErrorValidator returns a validation error JSON response
func ErrorValidator(w http.ResponseWriter, errors interface{}, message *string) {
	finalMessage := "Error! The request not expected!"
	if message != nil {
		finalMessage = *message
	}

	response := ErrorResponse{
		ResponseCode:   http.StatusUnprocessableEntity,
		ResponseStatus: "failed-validation",
		Message:        finalMessage,
		Errors:         errors,
	}

	writeJSON(w, http.StatusUnprocessableEntity, response)
}

// ErrorNotFound returns a not found error JSON response
func ErrorNotFound(w http.ResponseWriter, errors interface{}, message *string) {
	finalMessage := "Error! The resource not found!"
	if message != nil {
		finalMessage = *message
	}

	response := ErrorResponse{
		ResponseCode:   http.StatusNotFound,
		ResponseStatus: "failed-not-found",
		Message:        finalMessage,
		Errors:         errors,
	}

	writeJSON(w, http.StatusNotFound, response)
}

// ErrorAuthentication returns an authentication error JSON response
func ErrorAuthentication(w http.ResponseWriter, errors interface{}, message *string) {
	finalMessage := "Error! The authentication failed!"
	if message != nil {
		finalMessage = *message
	}

	response := ErrorResponse{
		ResponseCode:   http.StatusUnauthorized,
		ResponseStatus: "failed-authentication",
		Message:        finalMessage,
		Errors:         errors,
	}

	writeJSON(w, http.StatusUnauthorized, response)
}

// ErrorServer returns a server error JSON response
func ErrorServer(w http.ResponseWriter, errors interface{}, message *string) {
	// Log the error
	log.Printf("Server Error: %v", errors)

	finalMessage := "Internal Server Error!"
	if message != nil {
		finalMessage = *message
	}

	response := ErrorResponse{
		ResponseCode:   http.StatusBadRequest,
		ResponseStatus: "failed-server",
		Message:        finalMessage,
		Errors:         errors,
	}

	writeJSON(w, http.StatusBadRequest, response)
}

// ErrorBadRequest returns a bad request error JSON response
func ErrorBadRequest(w http.ResponseWriter, errors interface{}, message *string) {
	finalMessage := "Bad Request!"
	if message != nil {
		finalMessage = *message
	}

	response := ErrorResponse{
		ResponseCode:   http.StatusBadRequest,
		ResponseStatus: "failed-bad-request",
		Message:        finalMessage,
		Errors:         errors,
	}

	writeJSON(w, http.StatusBadRequest, response)
}

// ParseJSONError converts technical JSON decode errors into human-readable messages
func ParseJSONError(err error) string {
	if err == nil {
		return "Invalid request payload"
	}

	errMsg := err.Error()

	// Type mismatch errors
	if strings.Contains(errMsg, "cannot unmarshal") {
		// Extract field name and expected type
		if strings.Contains(errMsg, "bool") {
			if strings.Contains(errMsg, "completed") {
				return "Field 'completed' must be a boolean value (true or false)"
			}
			return "One of the fields must be a boolean value (true or false)"
		}
		if strings.Contains(errMsg, "number") || strings.Contains(errMsg, "int") {
			if strings.Contains(errMsg, "user_id") {
				return "Field 'user_id' must be a number"
			}
			return "One of the fields must be a number"
		}
		if strings.Contains(errMsg, "string") {
			if strings.Contains(errMsg, "text") {
				return "Field 'text' must be a string"
			}
			return "One of the fields must be a string"
		}
		return "Invalid data type for one or more fields"
	}

	// Missing field errors
	if strings.Contains(errMsg, "missing") {
		return "Required field is missing"
	}

	// Invalid JSON syntax
	if strings.Contains(errMsg, "invalid character") {
		return "Invalid JSON format. Please check your request payload"
	}

	// EOF errors
	if strings.Contains(errMsg, "EOF") {
		return "Empty request body. Please provide valid JSON data"
	}

	// Unknown field errors
	if strings.Contains(errMsg, "unknown field") {
		return "Request contains unknown field(s)"
	}

	// Default fallback
	return "Invalid request payload. Please check your data format"
}

// writeJSON is a helper function to write JSON response
func writeJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"response_code": 500, "response_status": "failed-server", "message": "Internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
