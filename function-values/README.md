# Function Values in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=hnHDxo1YbaM)**

## Introduction

In Go, functions are values just like numbers or strings. You can assign them to variables, pass them around, and call them through those variables. This is called treating functions as *first-class values*.

This may feel unusual at first if you are used to languages where functions are only defined and called by name. In Go, treating functions as values opens the door to more flexible and reusable code.

In this lesson, you will learn how to assign functions to variables and call them like any other function.

## Uses / Use Cases

* Storing behaviour in variables.
* Choosing which function to run at runtime.
* Passing behaviour into other functions.
* Writing flexible and reusable code.
* Preparing for concepts like callbacks and higher-order functions.

## Assigning a Function to a Variable

You can assign a function to a variable just like any other value.

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello!")
}

func main() {
    sayHello := greet
    sayHello()
}
```

Expected output:

```
Hello!
```

Explanation:

* `greet` is a function.
* `sayHello` now holds a reference to that function.
* Calling `sayHello()` runs the same code as calling `greet()`.

The function is not executed during assignment. It only runs when you call it.

## Function Types

When you assign a function to a variable, the variable has a function type.

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    var operation func(int, int) int

    operation = add
    result := operation(2, 3)

    fmt.Println(result)
}
```

Expected output:

```
5
```

Explanation:

* `func(int, int) int` is a function type.
* The variable can hold any function with the same signature.
* This allows you to swap behaviour easily.

## Choosing Functions at Runtime

Because functions are values, you can decide which function to use at runtime.

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}

func main() {
    var calculate func(int, int) int

    calculate = add
    fmt.Println(calculate(10, 5))

    calculate = subtract
    fmt.Println(calculate(10, 5))
}
```

Expected output:

```
15
5
```

Explanation:

* The same variable holds different functions at different times.
* Each function shares the same signature.
* This pattern is useful when behaviour needs to change based on conditions.

## Anonymous Functions as Values

Functions do not need names. You can assign anonymous functions directly to variables.

```go
package main

import "fmt"

func main() {
    multiply := func(a int, b int) int {
        return a * b
    }

    fmt.Println(multiply(4, 5))
}
```

Expected output:

```
20
```

Explanation:

* The function has no name.
* It is created and assigned in one step.
* This is common for short, local behaviour.

## Challenge

Write a program that:

1. Defines two functions that take two integers and return an integer.
2. Assigns one of them to a variable of a function type.
3. Calls the function through the variable.
4. Reassigns the variable to the other function and calls it again.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func multiply(a int, b int) int {
    return a * b
}

func divide(a int, b int) int {
    return a / b
}

func main() {
    var op func(int, int) int

    op = multiply
    fmt.Println(op(10, 2))

    op = divide
    fmt.Println(op(10, 2))
}
```

Expected output:

```
20
5
```

If you didn't get this right the first time, that's completely normal. The key idea is understanding that functions can be stored and called just like other values.

## Summary

You have now learned how function values work in Go. Here is what you covered:

* Functions are first-class values.
* You can assign functions to variables.
* Function variables have specific function types.
* You can change behaviour by reassigning functions.
* Anonymous functions can be used as values.

This concept is foundational for more advanced patterns in Go, including callbacks, function parameters, and functional-style techniques.
