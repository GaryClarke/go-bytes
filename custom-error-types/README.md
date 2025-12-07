# Custom Error Types in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In an earlier lesson, you learned how to create basic errors with `errors.New` and how to return a value and an error from a function. Now that you understand interfaces, you can go a step further. In Go, an error is not just a string. It is an interface that any type can satisfy.

This opens the door to richer error handling. You can create custom error types that carry extra information beyond a simple string message.

In this lesson, you will learn what the `error` interface looks like and how to implement your own custom error types.

## Uses / Use Cases

- Creating errors with more detail than a simple string.
- Carrying extra information in your errors, such as which field failed validation.
- Designing functions that can return different types of errors.
- Understanding how real programs create meaningful error messages.

## Error is an Interface

The `error` type is defined like this:

```go
type error interface {
    Error() string
}
```

Any type that has an `Error() string` method satisfies the error interface. This is why:

- `errors.New()` works, because it returns a type that implements `Error() string`.
- You can write your own error types.
- You can pass different kinds of errors to functions that expect an `error`.
- You can return `nil` when you want to say there is no error.

If you've been wondering why stuff like this works: `fmt.Println(err)`, here's the answer. When Go's formatting functions (like `fmt.Println`, `fmt.Printf`, or `fmt.Sprintf`) need to print an error, they automatically call the `Error()` method on that error value. This is why you can print any error directly, and it will show a meaningful message. The formatting system knows that errors have an `Error()` method, so it calls it for you.

This is Go's way of describing error behaviour through a simple interface.

## Example: Creating a Custom Error Type

Here is a custom error type that carries more information than a string alone.

```go
package main

import "fmt"

type ValidationError struct {
    Field string
}

func (v ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field: %s", v.Field)
}

func validateName(name string) error {
    if name == "" {
        return ValidationError{Field: "name"}
    }
    return nil
}

func main() {
    err := validateName("")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Name is valid")
}
```

Let's break this down:

- **Custom error type**
  `ValidationError` is a struct with a field called `Field`.

- **Satisfying the error interface**
  The `Error()` method returns a string.
  Because of this, `ValidationError` is now an error.

- **Using the custom error**
  `validateName` returns `ValidationError` when something is wrong, and `nil` when everything is fine.

- **Checking the error**
  In `main()`, the code checks `err != nil` and handles the failure.

Custom errors give you more context without changing how errors behave.

## Expected Output

```
Error: validation failed for field: name
```

## Challenge

Create a custom error type named `NotFoundError` with a field called `Resource`.

1. Implement the `Error() string` method so it returns a message describing which resource was not found.

2. Write a function `findUser(id int)` that always returns a `NotFoundError` for this exercise.

3. In `main()`, call `findUser`, check the error, and print it.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

type NotFoundError struct {
    Resource string
}

func (e NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Resource)
}

func findUser(id int) error {
    return NotFoundError{Resource: "user"}
}

func main() {
    err := findUser(7)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("User found")
}
```

## Summary

You have now learned how to create custom error types in Go. Here is what you covered:

- `error` is an interface with one method, `Error() string`.
- Any type that implements `Error()` becomes an error.
- You can create custom error types to carry extra information beyond a simple string.
- When you print an error, Go automatically calls the `Error()` method.
- Custom errors let you provide more context about what went wrong.

These patterns form the basis of how Go manages failures. You will use them in almost every real program you write.

