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
- **Testing invalid input for create handlers**: Use a table-driven test with cases for missing fields, invalid values, and malformed JSON; each case has payload, expected status (422 or 400), and optionally expected error keys (wantKeys); assert status then, for 422, decode response and check the errors object contains those keys
- **422 vs 400 in tests**: For create handlers, assert 422 when JSON is valid but validation fails (response has errors object); assert 400 when the body is not valid JSON (no errors object)

## Integration / Database Testing

- **Testing against the database**: Run handler tests with a real database (e.g. in-memory SQLite) so handlers, store, and SQL are exercised together; catches more bugs than mocking the store alone
- **In-memory SQLite for tests**: Use `sql.Open("sqlite", ":memory:")` so tests use a real database in RAM; fast, no disk, no manual cleanup
- **Test setup with Migrate and Seed**: In test setup, open the database, run Migrate and SeedIfEmpty, return an App with NewStores(db); use t.Cleanup to close the connection
- **Blank import for database drivers**: Use `_ "driver/package"` to register the driver with database/sql so sql.Open works

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

## Database Writes / Inserts- **ExecContext for INSERT**: Use `ExecContext` for INSERT, UPDATE, and DELETE (statements that don't return rows); use QueryContext/QueryRowContext for SELECT
- **LastInsertId()**: After `ExecContext` for an INSERT, call `LastInsertId()` on the Result to get the auto-generated ID and set it on your struct
- **Insert store methods**: Take a pointer to the entity (with required fields set, ID typically zero), run the INSERT with placeholders, set the struct's ID from `LastInsertId()`, and return the same pointer
- **Parameterized INSERT**: Use placeholders for values; omit auto-generated columns (e.g. `id`); let the database generate the ID## Handler and Store- **Wiring handler to store**: Call the store's Insert (or other write) method from the create handler with the domain model built from validated input
- **Handling insert errors**: On Insert failure, log the error and return 500 Internal Server Error; don't expose internal error details in the response
- **Returning the saved resource**: Use the value returned from Insert (with ID set) when writing the 201 response so the client receives the full created resource including its ID

## Borrowed / Inspiration

### singleflight - Duplicate Function Call Suppression

**Summary:** Use singleflight in front of the cache-miss path for any Go service with caching. It deduplicates concurrent calls for the same key so only one execution runs; all other callers wait and receive the same result. Prevents cache stampedes (thundering herd) where many requests hit the backend at once when cache expires.

**Key points:**
- Package: `golang.org/x/sync/singleflight`
- For a given key, only one execution of a function is in-flight at a time
- Concurrent callers for the same key wait and share the result
- Example: 1000 requests for same uncached data; without singleflight all hit DB; with singleflight 1 query runs, 999 wait and share
- Protects downstream services from redundant load
- Deduplicates expensive computations

**Links:**
- https://pkg.go.dev/golang.org/x/sync/singleflight

**Original inspiration:** Put singleflight in front of your cache-miss path for free deduplication. It ensures only one execution per key; other concurrent callers wait and receive the same result. Critical for preventing cache stampedes and protecting downstream services.