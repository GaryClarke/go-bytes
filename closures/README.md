# Closures in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=6mAScS0dEHU)**

## Introduction

In the previous lessons, you learned that functions in Go can be treated as values and that anonymous functions can be defined inline. Closures build directly on those ideas.

A *closure* is a function that can access variables from the scope where it was created, even after that scope has finished executing. This allows functions to remember and work with data from their surrounding context.

In this lesson, you will learn what closures are, how they work, and why they can be useful.

## Uses / Use Cases

* Remembering state between function calls.
* Creating functions with built-in configuration.
* Encapsulating behaviour and data together.
* Avoiding global variables.
* Writing cleaner and more flexible code.

## Functions Can Capture Variables

A closure captures variables from its surrounding scope.

```go
package main

import "fmt"

func main() {
    message := "Hello"

    greet := func() {
        fmt.Println(message)
    }

    greet()
}
```

Expected output:

```
Hello
```

Explanation:

* The anonymous function uses `message`, even though it is not a parameter.
* `message` comes from the surrounding scope.
* This combination of function and captured variable is a closure.

## Captured Variables Are Remembered

Captured variables are not copied. The function remembers the variable itself.

```go
package main

import "fmt"

func main() {
    count := 0

    increment := func() {
        count++
        fmt.Println(count)
    }

    increment()
    increment()
    increment()
}
```

Expected output:

```
1
2
3
```

Explanation:

* The function captures `count`.
* Each call updates the same variable.
* The value is remembered between calls.

This is one of the most important aspects of closures.

## Returning a Closure from a Function

Closures become especially useful when they are returned from functions.

```go
package main

import "fmt"

func makeCounter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()

    fmt.Println(counter())
    fmt.Println(counter())
    fmt.Println(counter())
}
```

Expected output:

```
1
2
3
```

Explanation:

* `makeCounter` creates a local variable `count`.
* The returned function captures `count`.
* Even after `makeCounter` finishes, `count` still exists.
* Each call updates the same captured variable.

This pattern is common and very powerful.

## Each Closure Has Its Own State

Each call to `makeCounter` creates a new, independent closure.

```go
package main

import "fmt"

func makeCounter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}

func main() {
    a := makeCounter()
    b := makeCounter()

    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(b())
}
```

Expected output:

```
1
2
1
```

Explanation:

* `a` and `b` each have their own captured `count`.
* They do not affect each other.
* This makes closures useful for isolated state.

## Aside: Why Closures Are Usually Anonymous

In Go, closures are usually written as anonymous functions. This is not a rule, but a consequence of how the language works.

There are a few reasons for this:

* Named functions in Go cannot be declared inside other functions.
* Variables that a closure captures usually exist inside a function.
* Anonymous functions can be defined inline, right where those variables exist.

Because of this, anonymous functions are the natural way to create closures in Go. What makes a function a closure is not whether it has a name, but whether it captures variables from its surrounding scope.

## Closures vs Global Variables

Closures let you keep state without using global variables.

Instead of this:

```go
var count int
```

You can keep the state private inside a function. This leads to safer and more predictable code.

## Challenge

Write a function called `makeAdder` that:

1. Takes an integer value.
2. Returns a function.
3. The returned function should take another integer and add it to the original value.
4. Call the returned function multiple times and print the results.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func makeAdder(base int) func(int) int {
    return func(x int) int {
        return base + x
    }
}

func main() {
    addTen := makeAdder(10)

    fmt.Println(addTen(5))
    fmt.Println(addTen(20))
}
```

Expected output:

```
15
30
```

If you didn't get this right the first time, that's completely normal. The key idea is understanding that the returned function remembers the values it captured.

## Summary

You have now learned how closures work in Go. Here is what you covered:

* Closures are functions that capture variables from their surrounding scope.
* Captured variables are remembered between calls.
* Closures can be returned from functions.
* Each closure has its own independent state.
* Closures help avoid global variables and keep state local.

Closures are a powerful concept in Go that enable many advanced patterns, including state management and function factories.
