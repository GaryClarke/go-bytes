# Variadic Functions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Sometimes you want to write a function that accepts any number of arguments. For example, adding numbers together, collecting messages, or building a list of values. In Go, you can do this with a variadic function.

A variadic function uses `...` before the final parameter type to say, "You can pass as many of these as you like." Inside the function, the parameter behaves like a slice. This means you can loop over the values, process them, and return whatever you need.

In this lesson, you will learn how to define variadic functions, how to call them, and how the arguments appear inside the function.

## Uses / Use Cases

* Summing values without creating a slice manually.
* Building flexible functions that accept a variable number of arguments.
* Creating convenience helpers for logging, formatting, or calculations.
* Making functions easier to call without extra setup.

## Example: A Simple Variadic Function

Here is a function that adds any number of integers together.

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))
    fmt.Println(sum(10, 20))
    fmt.Println(sum())
}
```

Expected output:

```
6
30
0
```

Explanation:

* `nums ...int` allows the function to accept zero, one, or many integers.
* Inside the function, `nums` is a slice.
* Calling `sum()` with no arguments gives an empty slice, so the total is zero.

Variadic functions provide flexibility without extra work.

## Example: Mixing Fixed and Variadic Parameters

Variadic parameters can be used alongside regular parameters, as long as the variadic one is last.

```go
package main

import "fmt"

func greet(prefix string, names ...string) {
    for _, name := range names {
        fmt.Println(prefix, name)
    }
}

func main() {
    greet("Hello", "Alice", "Bob", "John")
}
```

Expected output:

```
Hello Alice
Hello Bob
Hello John
```

Explanation:

* `prefix` is a single value.
* `names` collects all remaining arguments into a slice.
* The order must always be: fixed parameters first, variadic last.

## Challenge

Write a function called `multiplyAll` that:

1. Accepts any number of integers as a variadic parameter.
2. Returns the product of all numbers.
3. Returns `1` if no numbers are provided.

Then call it with:

* A single number
* Several numbers
* No numbers

Print the results.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func multiplyAll(nums ...int) int {
    if len(nums) == 0 {
        return 1
    }

    result := 1
    for _, n := range nums {
        result *= n
    }
    return result
}

func main() {
    fmt.Println(multiplyAll(5))
    fmt.Println(multiplyAll(2, 3, 4))
    fmt.Println(multiplyAll())
}
```

Expected output:

```
5
24
1
```

If you did not get it exactly right the first time, that is completely normal. Variadic functions can be tricky when you first encounter them. The important thing is understanding how the `...` syntax collects arguments into a slice.

## Summary

You have now learned how variadic functions work in Go. Here is what you covered:

* A variadic parameter uses `...` before its type.
* Inside the function, that parameter becomes a slice.
* You can call the function with zero, one, or many arguments.
* You can mix fixed and variadic parameters as long as the variadic one is last.
* Variadic functions are useful for flexible, readable function calls.

These patterns appear throughout the Go standard library, so understanding them will help you recognise and use them confidently.

Note: You can also pass a slice to variadic functions using `...`, which we'll cover in a future lesson.

