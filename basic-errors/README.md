# Basic Errors in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Go does not use exceptions for error handling. Instead, errors are ordinary values that you create, return, and check just like anything else. This style keeps control flow simple and makes it easy to see what your code is doing.

In this lesson, you will learn how to create an error, how to return a value alongside an error, and how to check whether an error occurred by comparing it with `nil`.

## Uses / Use Cases

- Creating simple error messages using the standard library.
- Returning both a result and an error from a function.
- Checking for problems by comparing an error with `nil`.
- Stopping a function early when something goes wrong.
- Matching the most common Go pattern for error handling.

## Example

Here is a function that converts a number of minutes to seconds. It returns both a result and an error if the input is invalid:

```go
package main

import (
    "errors"
    "fmt"
)

func toSeconds(minutes int) (int, error) {
    if minutes < 0 {
        return 0, errors.New("minutes cannot be negative")
    }
    return minutes * 60, nil
}

func main() {
    seconds, err := toSeconds(5)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Seconds:", seconds)
}
```

Let’s break this down:

- **Creating an error**  
  `errors.New("minutes cannot be negative")`  
  This makes a simple error value with a message.

- **Returning a value and an error**  
  The function returns two things. If something is wrong, we return a default value and the error. If everything is fine, we return the real result and `nil`.

- **Checking the error**  
  In `main()`, we look at the second return value. If it is not `nil`, something went wrong. If it is `nil`, we know the result is safe to use.

This pattern appears everywhere in Go and will quickly become second nature.

## Expected Output

```
Seconds: 300
```

If you passed a negative number, you would see:
```
Error: minutes cannot be negative
```

## Challenge

Write a function called `squareRoot` that takes an integer and returns:

1. The square root (use `float64` for the result)
2. An error if the input is negative

In `main()`, call the function with both valid and invalid inputs to test your code. Print the error when one occurs. Print the result when everything is fine.

Run it using:
```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here’s one way to solve it:

```go
package main

import (
    "errors"
    "fmt"
    "math"
)

func squareRoot(n int) (float64, error) {
    if n < 0 {
        return 0, errors.New("cannot take the square root of a negative number")
    }
    return math.Sqrt(float64(n)), nil
}

func main() {
    result, err := squareRoot(16)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    result, err = squareRoot(-9)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

## Summary

You have now learned the basics of Go’s error handling style. Here is what you covered:

- Errors are normal values created with `errors.New`.
- Functions often return a result and an error together.
- You check for errors by comparing with `nil`.
- Returning early when something is wrong keeps code clear.
- This pattern replaces exceptions and encourages simple control flow.

This foundation prepares you for more advanced error handling as you continue learning Go.
