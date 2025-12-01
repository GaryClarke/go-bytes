# Introduction to Methods and Receivers

**[Watch this lesson on YouTube](https://youtube.com/watch?v=DnB39HAhEy4)**

## Introduction

Go lets you attach functions to types. These attached functions are called **methods**, and the type you attach them to is known as the **receiver**.

Methods give your types behaviour, which helps you organise code and keep related logic together.

You have already written standalone functions. Methods are similar, but they have one important difference. A method is defined with a receiver, which tells Go which type the method belongs to.

In this lesson, you will learn how to create a method, how to call it, and when methods make your code clearer.

## Uses / Use Cases

- Adding behaviour to your own custom types.
- Grouping related functions around a single type.
- Improving readability by keeping data and actions together.
- Creating helper methods, such as `FullName()` or `Area()`.
- Building more advanced types as your programs grow.

## Example

Here is a simple type called `User` with a method that returns the full name:

```go
package main

import "fmt"

type User struct {
    FirstName string
    LastName  string
}

func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

func main() {
    user := User{FirstName: "Alice", LastName: "Smith"}
    fmt.Println(user.FullName())
}
```

Let's break this down:

- **Defining a type**
  `type User struct { ... }` creates a simple struct with two fields.

- **Adding a method**
  `func (u User) FullName() string` defines a method named `FullName`.
  The part in parentheses before the method name is the receiver.
  It tells Go that this method belongs to the `User` type and that inside the method the value will be called `u`.

- **Calling the method**
  Once you have a `User` value, you call the method with dot syntax.
  `user.FullName()` looks and behaves just like calling a function on an object.

This pattern keeps code tidy, especially when your type has multiple actions associated with it.

## Expected Output

```
Alice Smith
```

## Challenge

Create a struct type called `Rectangle` with two integer fields:

- `Width`
- `Height`

Add a method called `Area` that returns the area as an integer.

In `main()`, create a rectangle and print the result of calling `Area()`.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
}
```

## Summary

You have now learned the basics of methods and receivers in Go. Here is what you covered:

- Methods are functions that belong to a type.
- A receiver tells Go which type the method is associated with.
- Methods are called using dot syntax, just like fields.
- You can use methods to give your types meaningful behaviour.
- Keeping related actions inside methods helps your code stay organised and clear.

This foundation prepares you for more advanced concepts, including pointer receivers and interface-based design.

