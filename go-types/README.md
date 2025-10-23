# Go Types

**[Watch this lesson on YouTube](https://youtube.com/watch?v=xYERjE-zg-Y)**

## Introduction

In Go, every value has a **type**. A type tells the computer what kind of data you're working with, and what operations are allowed.
Types are core to how Go works. They help catch bugs before your program runs and make your code easier to understand.

This lesson introduces Go's basic types and shows how to declare and assign values using both **explicit types** and **type inference**.
You'll learn about `int`, `string`, `float64`, and `bool`, and how to use them in your own programs.

Understanding types early makes everything else in Go easier to learn.

## Uses / Use Cases

- Declaring variables with clear, predictable data
- Storing text, numbers, or true/false values
- Performing calculations and comparisons safely
- Writing code that behaves as expected

## Example

Here's a simple Go program that uses four different types:

```go
package main

import "fmt"

func main() {
    name := "Alice"        // string
    age := 35              // int
    isMember := true       // bool
    height := 1.82         // float64

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Member:", isMember)
    fmt.Println("Height:", height)
}
```

Let's walk through the types:

* `string` is used for text. Anything in double quotes.
* `int` is used for whole numbers.
* `bool` is used for true or false values.
* `float64` is used for decimal numbers.

The `:=` operator tells Go to figure out the type based on the value.
If you want to be explicit, you can use `var` and write the type yourself:

```go
var country string = "UK"
var year int = 2025
var isActive bool = false
var temperature float64 = 22.5
```

Both styles work. Use whichever helps make your code clearer.

> **Note:** Go won't let you mix incompatible types. For example, you can't add a string to a number. This helps prevent bugs by catching type errors before your program runs.

## Expected Output

```
Name: Alice
Age: 35
Member: true
Height: 1.82
```

## Challenge

Write a Go program that creates a simple user profile. Declare variables for:

1. A `string` for your favorite programming language
2. An `int` for how many years you've been programming  
3. A `bool` for whether you prefer coding in the morning or evening (true for morning, false for evening)
4. A `float64` for your programming confidence level (0.0 to 1.0)

Print each value with a descriptive label like:
```
Favorite Language: Go
Years Programming: 2
Morning Coder: true
Confidence Level: 0.8
```

Run the program using:

```
go run main.go
```

## Solution

Well done if you gave this a try. Here's one possible solution:

```go
package main

import "fmt"

func main() {
    favoriteLanguage := "Go"
    yearsProgramming := 1
    morningCoder := true
    confidenceLevel := 0.7

    fmt.Println("Favorite Language:", favoriteLanguage)
    fmt.Println("Years Programming:", yearsProgramming)
    fmt.Println("Morning Coder:", morningCoder)
    fmt.Println("Confidence Level:", confidenceLevel)
}
```

If your solution is different, that's fine.
The goal is to get comfortable declaring variables and working with basic types.

## Summary

In this lesson, you learned:

* Every value in Go has a type
* `string`, `int`, `float64`, and `bool` are the core types for text, numbers, and logic
* Use `:=` for automatic typing or `var` to be explicit
* Go's type system helps you write safe, predictable programs

You'll use types everywhere in Go, so this foundation will serve you well as you build real applications.
