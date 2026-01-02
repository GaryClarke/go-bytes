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

