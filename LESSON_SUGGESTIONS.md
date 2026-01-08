# Lesson Suggestions

This file tracks lesson ideas and suggestions for the "Build Your First Go App" course.

## Testing Concepts

- **Testing error cases**: Test handlers that return 404 (not found) and 500 (server error) status codes
- **Table-driven tests**: Use table-driven tests to test multiple scenarios with different inputs and expected outputs
- **Testing edge cases**: Test handlers with edge cases like zero values, empty strings, and boundary conditions
- **Testing handlers that modify data**: Test POST, PUT, and DELETE handlers that change database state
- **Testing POST endpoints**: Write tests for POST handlers that verify request body decoding, validation, and response structure

## Development Approaches

- **Outside-in development**: Build endpoints by starting with stubbed responses and working backwards to implement full logic
- **Stubbing responses**: Use hardcoded values to verify endpoint structure before implementing complex operations
- **Manual API testing**: Use curl, Postman, or similar tools to test endpoints during development

## API Design Patterns

- **Request structs**: Create separate structs for incoming request data, distinct from domain models
- **JSON struct tags**: Use struct tags to map JSON field names to struct fields
- **Package organization**: Organize request structs in dedicated packages (e.g., `internal/request`) separate from domain models
- **Input validation structs**: Use request structs to validate client input independently of database schema

## JSON Decoding and Request Bodies

- **JSON decoding from request bodies**: Use `json.NewDecoder(r.Body).Decode()` to parse incoming JSON data into structs
- **Request body as io.Reader**: Understand that `r.Body` is an `io.Reader` that can only be read once
- **Error handling for malformed JSON**: Return 400 Bad Request status codes when JSON decoding fails due to invalid client data
- **Pointer arguments for decoding**: Pass pointers to `Decode()` so it can modify struct fields

## Input Validation

- **Input validation functions**: Create validation functions that check required fields and valid value ranges before processing user input
- **Validation error maps**: Use `map[string]string` to return field-specific validation errors, where keys are field names and values are error messages
- **422 Unprocessable Entity status code**: Use 422 status code for validation errors when the request is well-formed but semantically invalid
- **Required field validation**: Check for empty strings and missing values to ensure required fields are provided
- **Numeric range validation**: Validate that numeric values are within acceptable bounds (e.g., positive integers, reasonable ranges)

