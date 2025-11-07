# Basic Functions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=HDgV7CA8Ae8)**

## Introduction

A function is a block of code that performs a specific task.

Functions allow you to break your program into smaller, reusable parts. You've already used one, `main()`, which is the special function where Go starts your program.

In this lesson, you'll learn how to define your own functions, pass data to them using parameters, and return results.

## Uses / Use Cases

* Reuse common logic across different parts of your program

* Break large problems into smaller, testable pieces

* Improve code readability and organization

## Example

Here's a simple function that greets a user by name:

```go
package main

import "fmt"

// greet takes a name and prints a message
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Alice")
    greet("Sam")
}
```

### Explanation

* `func greet(name string)`

  This defines a function named `greet`.

  It takes one parameter, `name`, which must be a string.

* `fmt.Println("Hello,", name)`

  This prints a greeting using the provided name.

* `greet("Alice")` and `greet("Sam")`

  These are function calls that pass different values to the `greet` function.

### Returning a value

Functions can also return a result. Here's an example:

```go
package main

import "fmt"

// square returns the square of a number
func square(n int) int {
    return n * n
}

func main() {
    result := square(5)
    fmt.Println("5 squared is", result)
}
```

* `func square(n int) int`

  This declares a function that takes an `int` and returns an `int`.

* `return n * n`

  This sends a value back to the caller.

* `result := square(5)`

  Calls the function and stores the returned value in `result`.

## Expected Output

```
Hello, Alice
Hello, Sam
5 squared is 25
```

## Challenge

Write a function named `double` that takes an integer and returns twice its value.

In your `main()` function, call `double(4)` and print the result using `fmt.Println()`.

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func double(n int) int {
    return n * 2
}

func main() {
    result := double(4)
    fmt.Println("Double of 4 is", result)
}
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!  
The important thing is understanding how functions can take parameters and return values.

## Summary

You learned how to:

* Define a function with `func name(params) { }`

* Call a function with specific values

* Pass arguments to functions

* Return values using the `return` keyword

Functions help organize your Go programs into logical, testable parts. You'll use them constantly in every Go project.

