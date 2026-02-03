# Lesson Suggestions

This file tracks lesson ideas and suggestions for the "Build Your First Go App" course.

## Testing Concepts

- **Testing error cases**: Test handlers that return 404 (not found) and 500 (server error) status codes
- **Table-driven tests**: Use table-driven tests to test multiple scenarios with different inputs and expected outputs
- **Completing table-driven tests**: Loop over test cases with `for _, tc := range tests`, use `t.Run()` to create subtests for each case, and run the function being tested inside the loop
- **Subtests with t.Run()**: Use `t.Run()` to create subtests that run independently, providing better failure messages that show which specific test case failed
- **Testing edge cases**: Test handlers with edge cases like zero values, empty strings, and boundary conditions
- **Testing handlers that modify data**: Test POST, PUT, and DELETE handlers that change database state
- **Testing POST endpoints**: Write tests for POST handlers that verify request body decoding, validation, and response structure

## Unit Testing

- **Unit tests vs integration tests**: Understand the difference between testing functions in isolation (unit tests) versus testing multiple components together (integration tests)
- **Testing functions in isolation**: Write tests that run individual functions without external dependencies like databases or HTTP servers
- **Arrange-act-assert pattern**: Structure tests with three clear phases: set up test data, call the function, verify results
- **Testing the happy path**: Start with tests that verify valid input works correctly before testing error cases
- **Checking empty maps**: Use `len(map) > 0` to verify that a map is empty (no errors) or contains entries (errors present)
- **Testing error maps**: Check error count with `len(errors) != len(wantKeys)` and verify specific error keys exist using map lookups with the `ok` variable
- **Asserting exact error matching**: Check both the length of error maps and the presence of specific keys to ensure errors match expectations exactly
- **Test file naming conventions**: Create test files with `_test.go` suffix in the same package as the code being tested
- **Running specific tests**: Use `go test -run TestName ./package/path` to run individual tests during development

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

## Database Writes / Inserts

- **ExecContext for INSERT**: Use `ExecContext` for INSERT, UPDATE, and DELETE (statements that don't return rows); use QueryContext/QueryRowContext for SELECT
- **LastInsertId()**: After `ExecContext` for an INSERT, call `LastInsertId()` on the Result to get the auto-generated ID and set it on your struct
- **Insert store methods**: Take a pointer to the entity (with required fields set, ID typically zero), run the INSERT with placeholders, set the struct's ID from `LastInsertId()`, and return the same pointer
- **Parameterized INSERT**: Use placeholders for values; omit auto-generated columns (e.g. `id`); let the database generate the ID## Handler and Store

- **Wiring handler to store**: Call the store's Insert (or other write) method from the create handler with the domain model built from validated input
- **Handling insert errors**: On Insert failure, log the error and return 500 Internal Server Error; don't expose internal error details in the response
- **Returning the saved resource**: Use the value returned from Insert (with ID set) when writing the 201 response so the client receives the full created resource including its ID
