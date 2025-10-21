# Declaring Variables in Go (short syntax)

**[Watch this lesson on YouTube](https://youtube.com/watch?v=PJz01dH8ZN4)**

## Introduction

Programs often need to store and reuse data, things like names, numbers, or settings.
In Go, you store data in **variables**.

A variable is like a labeled box in memory that holds a value. You can give it a name, put data inside, and use that name whenever you need the value again.

In this lesson, you'll learn how to declare variables in Go using the short declaration syntax `:=`. This approach lets Go automatically determine the variable's type for you.

## Uses / Use Cases

- Store and reuse values like names or numbers
- Make programs dynamic instead of hard-coding values
- Understand how Go infers types automatically
- Build a foundation for working with data in later lessons

## Example

Here's a simple Go program that uses two variables:

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 35

    fmt.Println(name)
    fmt.Println(age)
}
```

Let's break it down:

* `name := "Alice"`
  Declares a variable called `name` and assigns it the value `"Alice"`.
  Go automatically detects that this is a **string** because the value is inside quotes.

* `age := 35`
  Declares a variable called `age` and assigns it the value `35`.
  Go recognizes this as an **integer**.

* `fmt.Println(name)` and `fmt.Println(age)`
  Print the values of the variables to the terminal.

To run this file:

1. Save it as `main.go`
2. Open your terminal in the same folder
3. Run:

```
go run main.go
```

## Expected Output

```
Alice
35
```

> **Note:** You must declare a variable before you can use it. If you try to print a variable that hasn't been declared yet, Go will give you an error.

## Challenge

Create a program that declares three variables using `:=` and prints them:

1. One variable should hold the name of a city
2. Another should hold a number representing the current year
3. A third should hold a decimal number (like a temperature or price)

Then use `fmt.Println()` to print all three values.

## Solution

Here's one possible solution:

```go
package main

import "fmt"

func main() {
    city := "London"
    year := 2025
    temperature := 22.5

    fmt.Println(city)
    fmt.Println(year)
    fmt.Println(temperature)
}
```

If you wrote something slightly different, that's perfectly fine.
What matters is that you declared variables using `:=` and printed their values successfully.

## Summary

You learned how to declare and use variables in Go with the short declaration syntax `:=`.
This lets Go automatically determine the correct type for each variable based on the value you assign.

You can now store data in your programs and print it to the terminal.
In the next lesson, you'll learn more about the **types** Go assigns to these variables and how to work with them directly.



