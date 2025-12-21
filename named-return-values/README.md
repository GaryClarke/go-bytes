# Named Return Values in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Go allows you to give names to your return values right in the function signature. These are called *named return values*. When you use them, Go creates variables with those names as soon as the function begins. You can then assign to them directly inside the function, and a bare `return` statement will return their current values.

Named returns can make some functions easier to read, especially when there are multiple results or early returns. However, they should be used deliberately and not overused.

In this lesson, you will learn how named return values work, how to use them, and when they are helpful.

## Uses / Use Cases

* Returning multiple values without repeating them in the return statement.
* Functions where results are built gradually.
* Improving clarity when the meaning of each result is important.
* Simplifying code that needs several early returns.

## Example: Simple Named Return Values

Here is a function that uses a named return value for clarity.

```go
package main

import "fmt"

func square(n int) (result int) {
    result = n * n
    return
}

func main() {
    fmt.Println(square(5))
}
```

Expected output:

```
25
```

Explanation:

* The function signature declares a return variable called `result`.
* Inside the function, you assign directly to `result`.
* A bare `return` statement returns the current value of `result`.

This is the simplest use of named return values.

## Example: Multiple Named Return Values

Named returns become especially useful when there are several outputs.

```go
package main

import "fmt"

func divide(a, b float64) (quotient float64, ok bool) {
    if b == 0 {
        ok = false
        return
    }

    quotient = a / b
    ok = true
    return
}

func main() {
    q, ok := divide(10, 2)
    fmt.Println(q, ok)

    q, ok = divide(10, 0)
    fmt.Println(q, ok)
}
```

Expected output:

```
5 true
0 false
```

Explanation:

* `quotient` and `ok` are declared when the function starts.
* You set them as needed inside the function.
* `return` sends them back without listing them explicitly.

This pattern is common when returning a value plus a boolean status.

## Example: Building a Result Step by Step

Named returns can make sense when building a result through several steps.

```go
package main

import "fmt"

func process(n int) (output int) {
    output = n * 2
    output = output + 3
    output = output * 4
    return
}

func main() {
    fmt.Println(process(5))
}
```

Expected output:

```
44
```

Here it is clear which variable is being modified throughout the function.

## A Note of Caution

Named return values can improve clarity in some functions, but using them everywhere can make code harder to read. Use them when:

* the names add meaning, and
* returning the variables directly improves readability.

Avoid them when:

* you find yourself searching the function to understand what is being returned, or
* a simple explicit `return value1, value2` would be clearer.

## Challenge

Write a function called `stats` that:

1. Accepts a slice of integers.
2. Uses named return values `total int` and `count int`.
3. Calculates the sum of all numbers and the number of elements.
4. Uses a bare `return` to return the named variables.

Then call the function and print the results.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func stats(nums []int) (total int, count int) {
    for _, n := range nums {
        total += n
        count++
    }
    return
}

func main() {
    total, count := stats([]int{3, 5, 7})
    fmt.Println("Total:", total)
    fmt.Println("Count:", count)
}
```

Expected output:

```
Total: 15
Count: 3
```

If you did not get it exactly right the first time, that is completely normal. Named return values can be tricky when you first encounter them. The important thing is understanding how Go creates the variables for you and how a bare return works.

## Summary

You have now learned how to use named return values in Go. Here is what you covered:

* Named return values are declared in the function signature.
* Go creates these variables when the function begins.
* You can assign to them directly.
* A bare `return` returns the current values of the named variables.
* Named returns help when results have clear meanings or when a function builds its result step by step.
* Use them carefully and only when they make the function easier to read.

Named return values are a useful tool to have in your Go toolbox, especially for functions with multiple outputs.



