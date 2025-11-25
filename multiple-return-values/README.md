# Multiple Return Values in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=Bh4LgJgDQRY)**

## Introduction

Go allows functions to return more than one value. This feature is built into the language and appears in many common patterns. Multiple return values let a function give you primary data along with useful extra information, all in one call.

In this lesson, you’ll learn how Go functions return multiple values and why this approach makes certain tasks clearer and more expressive.

## Uses / Use Cases

- Returning a value and a boolean to indicate success.
- Returning two related results at once, such as quotient and remainder.
- Splitting a calculation into two outputs instead of forcing separate function calls.
- Writing simple, predictable function signatures.

## Example

Here’s a simple example of a function that divides two integers and returns both the quotient and the remainder:

```go
package main

import "fmt"

func divide(a, b int) (int, int, bool) {
    if b == 0 {
        return 0, 0, false
    }
    return a / b, a % b, true
}

func main() {
    q, r, ok := divide(10, 3)
    if !ok {
        fmt.Println("Invalid division")
        return
    }
    fmt.Println("Quotient:", q)
    fmt.Println("Remainder:", r)
}
```

Let’s break this down:

- **Function signature**  
  `func divide(a, b int) (int, int, bool)`  
  This function returns three values: an integer quotient, an integer remainder, and a boolean that signals whether the result is valid.

- **Returning multiple values**  
  When `b` is zero, we return zeros plus `false` to indicate failure. When the inputs are valid, we return the computed values plus `true`.

- **Receiving multiple values**  
  Calling the function gives us three results at once: `q`, `r`, and `ok`.

- **Boolean success flag**  
  We check `ok` before using the other values. This pattern is common in Go and easy to understand.

## Expected Output

```
Quotient: 3
Remainder: 1
```

If you tried dividing by zero, the output would be:
```
Invalid division
```

## Challenge

Write a function called `minMax` that takes two integers and returns:

1. The smaller number
2. The larger number
3. A boolean `ok` value that is `true` when the numbers are different and `false` when they are equal

In `main()`, call this function with any two numbers and print out the results.

If `ok` is `false`, print a message that the numbers are equal.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here’s one way to solve it:

```go
package main

import "fmt"

func minMax(a, b int) (int, int, bool) {
    if a == b {
        return a, b, false
    }
    if a < b {
        return a, b, true
    }
    return b, a, true
}

func main() {
    min, max, ok := minMax(4, 9)
    if !ok {
        fmt.Println("The numbers are equal")
        return
    }
    fmt.Println("Min:", min)
    fmt.Println("Max:", max)
}
```

## Summary

You’ve learned how Go functions return multiple values and how this feature supports clear and expressive code. Here’s what you covered:

- A function can return more than one result.
- You can unpack multiple values on the caller side using multiple variables.
- Boolean flags like `ok` make it easy to check whether a result is usable.
- Returning several values at once keeps code simple and avoids unnecessary extra function calls.
