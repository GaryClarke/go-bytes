# Introduction to Interfaces

**[Watch this lesson on YouTube](https://youtube.com/watch?v=W3xS52aR0dM)**

## Introduction

Interfaces are how Go describes behaviour. An interface lists one or more method signatures. Any type that has those methods automatically satisfies the interface. You do not need to write anything special or declare that your type implements it. Go checks it for you.

This makes your code flexible. You can write functions that accept an interface, then pass in different concrete types that behave in the required way.

In this lesson, you will learn what an interface is, how types satisfy an interface, and how interfaces help you write clean and reusable code.

## Uses / Use Cases

* Grouping types by shared behaviour.

* Writing functions that accept many different types.

* Making code easier to extend without changing existing functions.

* Swapping implementations for testing.

* Designing clear and simple program structures.

## Example

Here is an interface called `Shape` and two types that satisfy it by providing an `Area()` method.

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    c := Circle{Radius: 3}
    r := Rectangle{Width: 5, Height: 4}

    printArea(c)
    printArea(r)
}
```

Let's break this down:

* **Defining an interface**

  `type Shape interface { Area() float64 }` describes behaviour.

  Any type with an `Area()` method that returns a `float64` satisfies it.

* **Implicit satisfaction**

  Neither `Circle` nor `Rectangle` says they implement `Shape`.

  They simply have the required method, so Go considers them valid shapes.

* **Using the interface**

  The `printArea` function accepts a `Shape`.

  This means it will work with any type that meets the interface's requirements.

* **Calling methods through the interface**

  Inside `printArea`, calling `s.Area()` uses the correct method for the concrete type.

Interfaces allow you to depend on behaviour rather than concrete structures.

## Expected Output

```
Area: 28.274333882308138
Area: 20
```

## Challenge

Create a small logging system using an interface.

1. Define an interface called `Logger` with one method:

   ```
   Log(message string)
   ```

2. Create two types that satisfy this interface:

   * `ConsoleLogger` which prints the message to the terminal.

   * `PrefixedLogger` which also prints to the terminal but adds a fixed prefix, for example `[INFO]`.

3. Write a function called `runLog` that accepts a `Logger` and a message, then calls the `Log` method.

4. In `main()`, create both loggers and pass them to `runLog` with different messages.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
    fmt.Println(message)
}

type PrefixedLogger struct {
    Prefix string
}

func (p PrefixedLogger) Log(message string) {
    fmt.Println(p.Prefix, message)
}

func runLog(l Logger, msg string) {
    l.Log(msg)
}

func main() {
    plain := ConsoleLogger{}
    info := PrefixedLogger{Prefix: "[INFO]"}

    runLog(plain, "Starting application")
    runLog(info, "User logged in")
}
```

## Summary

You have learned the basics of interfaces in Go. Here is what you covered:

* Interfaces describe behaviour through method signatures.

* A type satisfies an interface automatically by having the required methods.

* There is no keyword or special declaration needed.

* Functions that accept interfaces can work with many different types.

* Interfaces help you build flexible, testable, and reusable code.


