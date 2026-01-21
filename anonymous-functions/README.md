# Anonymous Functions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=SD1yUQlpmcg)**

## Introduction

In Go, functions do not always need a name. You can create functions inline and use them directly. These are called *anonymous functions* or *function literals*.

Anonymous functions are useful when you need a small piece of behaviour that is only used in one place. They help keep code close to where it is used and avoid creating extra named functions.

In this lesson, you will learn how to write anonymous functions, assign them to variables, and call them just like regular functions.

## Uses / Use Cases

* Writing short, local pieces of behaviour.
* Avoiding unnecessary named functions.
* Passing behaviour around as values.
* Creating simple one-off functions.
* Improving readability for small operations.

## Defining an Anonymous Function

An anonymous function looks like a normal function, but without a name.

```go
package main

import "fmt"

func main() {
    greet := func() {
        fmt.Println("Hello!")
    }

    greet()
}
```

Expected output:

```
Hello!
```

Explanation:

* The function has no name.
* It is assigned directly to a variable.
* Calling the variable runs the function.

This is a function literal.

## Anonymous Functions with Parameters and Return Values

Anonymous functions can take parameters and return values.

```go
package main

import "fmt"

func main() {
    add := func(a int, b int) int {
        return a + b
    }

    result := add(3, 4)
    fmt.Println(result)
}
```

Expected output:

```
7
```

Explanation:

* The function signature is defined inline.
* The variable type is inferred.
* The function behaves like any other function.

## Calling an Anonymous Function Immediately

You can also define and call an anonymous function in one step.

```go
package main

import "fmt"

func main() {
    func(message string) {
        fmt.Println(message)
    }("Hello from an anonymous function")
}
```

Expected output:

```
Hello from an anonymous function
```

Explanation:

* The function is defined inside `main`.
* It is immediately invoked.
* This pattern is useful for short, self contained logic.

## Anonymous Functions as Function Values

Anonymous functions are still function values, so they can be stored and passed around.

```go
package main

import "fmt"

func main() {
    var operation func(int, int) int

    operation = func(a int, b int) int {
        return a * b
    }

    fmt.Println(operation(4, 5))
}
```

Expected output:

```
20
```

Explanation:

* The variable has a function type.
* The anonymous function matches the required signature.
* The function can be called through the variable.

## What This Lesson Does Not Cover

Anonymous functions are often used with variables from their surrounding scope. This is called a *closure*. That topic is important, but it is covered in a separate lesson.

For now, focus on anonymous functions as unnamed, inline functions.

## Challenge

Write a program that:

1. Declares a variable of a function type.
2. Assigns an anonymous function that takes two integers and returns an integer.
3. Calls the function and prints the result.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func main() {
    var calculate func(int, int) int

    calculate = func(a int, b int) int {
        return a + b
    }

    fmt.Println(calculate(6, 7))
}
```

Expected output:

```
13
```

If you didn't get this right the first time, that's completely normal. The important thing is recognising that functions can be created inline and used just like named functions.

## Summary

You have now learned how anonymous functions work in Go. Here is what you covered:

* Anonymous functions are functions without names.
* They are written as function literals.
* They can be assigned to variables.
* They can take parameters and return values.
* They can be called immediately or stored for later use.

Anonymous functions are a core building block for many Go patterns, including closures and higher-order functions.
