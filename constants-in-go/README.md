# Constants in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=n-gA1DCl3yA)**

## Introduction

In Go, a constant is a name that refers to a fixed value that doesn't change while the program runs. Constants are useful for values you want to use in multiple places without accidentally changing them.

You define constants using the `const` keyword. Unlike variables, constants must be assigned a value at the time they are declared. You can't update a constant after it's been created.

Go supports constants for simple types like numbers, strings, and booleans.

## Uses / Use Cases

* Defining fixed values like tax rates or limits
* Making your code easier to read and update
* Avoiding accidental changes to values that should never vary
* Improving maintainability by giving names to repeated values

## Example

Here's a short program that uses constants:

```go
package main

import "fmt"

const maxUsers = 100
const welcomeMessage = "Welcome to Go Bytes"

func main() {
    fmt.Println(welcomeMessage)
    fmt.Println("The maximum number of users is", maxUsers)
}
```

Let's break it down:

* `const maxUsers = 100`
  This defines a constant named `maxUsers` with an integer value of 100.

* `const welcomeMessage = "Welcome to Go Bytes"`
  This defines a string constant that stores a greeting message.

Constants do not use `:=` because that's only for variables. You must assign a constant its value using `=` when you declare it.

**Note**: Go has a unique feature called *typed* and *untyped* constants. When you declare a constant without specifying a type (like our examples above), Go creates an "untyped" constant that can be used with different types as needed. This flexibility is one of Go's strengths.

## Expected Output

```
Welcome to Go Bytes
The maximum number of users is 100
```

## Challenge

Write a program that defines two constants:

* Your favorite number
* A short quote you like

Then print both values using `fmt.Println()`.

Remember:

* Use `const` for both
* You must assign the value at the time you declare the constant

## Solution

Here's one possible solution:

```go
package main

import "fmt"

const favoriteNumber = 7
const quote = "Keep it simple."

func main() {
    fmt.Println("My favorite number is", favoriteNumber)
    fmt.Println("Quote of the day:", quote)
}
```

If your output looks different, that's totally fine as long as you used constants and printed them successfully.

## Summary

You learned:

* How to declare constants using the `const` keyword
* Constants must be assigned at declaration and cannot change
* Constants are great for values that stay fixed throughout your program
* Go supports constants for basic types like numbers, strings, and booleans
* You can print constants just like variables

Constants help make your code more predictable, maintainable, and readable.


