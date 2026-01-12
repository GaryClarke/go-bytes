# Multiple Assignment in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=nlp1NUZt5VQ)**

## Introduction

Go allows you to assign multiple values at once using a feature called *multiple assignment*. This lets you set, update, or swap values in a single statement without needing temporary variables.

Multiple assignment is used throughout Go, especially when working with function return values, swapping variables, and updating related values together. Once you are familiar with it, your code becomes shorter and easier to read.

In this lesson, you will learn how multiple assignment works and see some of the most common patterns.

## Uses / Use Cases

* Swapping values without a temporary variable.
* Assigning multiple return values from a function.
* Updating several related variables at the same time.
* Writing clearer and more concise code.

## Example: Assigning Multiple Values

You can assign more than one variable in a single statement.

```go
package main

import "fmt"

func main() {
    a, b := 10, 20

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

Explanation:

* `a, b := 10, 20` assigns `10` to `a` and `20` to `b`.
* The values are assigned in order from left to right.
* This works with both short declaration (`:=`) and regular assignment (`=`).

## Expected Output

```
a: 10
b: 20
```

## Example: Swapping Values

One of the most well-known uses of multiple assignment is swapping values.

```go
package main

import "fmt"

func main() {
    x := 5
    y := 10

    x, y = y, x

    fmt.Println("x:", x)
    fmt.Println("y:", y)
}
```

Expected output:

```
x: 10
y: 5
```

Explanation:

* Go evaluates the right side first.
* The values are then assigned to the left side.
* No temporary variable is needed.

This is safe, clear, and idiomatic Go.

## Example: Multiple Assignment with Function Returns

Many Go functions return more than one value. Multiple assignment lets you capture them cleanly.

```go
package main

import "fmt"

func getPoint() (int, int) {
    return 3, 7
}

func main() {
    x, y := getPoint()
    fmt.Println("x:", x, "y:", y)
}
```

Explanation:

* `getPoint` returns two values.
* Multiple assignment assigns each return value to a variable.
* This pattern is extremely common in Go.

## Challenge

Write a program that:

1. Declares two integer variables.
2. Prints their initial values.
3. Swaps their values using multiple assignment.
4. Prints the values again to confirm they were swapped.

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
    a := 3
    b := 8

    fmt.Println("Before swap:", a, b)

    a, b = b, a

    fmt.Println("After swap:", a, b)
}
```

Expected output:

```
Before swap: 3 8
After swap: 8 3
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!
The important part is understanding that all right hand values are evaluated before any assignment happens.

## Aside: Updating Related Values Together

Multiple assignment is useful when updating related values at the same time.

```go
package main

import "fmt"

func main() {
    a, b := 1, 2

    a, b = a+1, b+1

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

Explanation:

* Both expressions on the right are evaluated first.
* The results are then assigned together.
* This avoids subtle bugs that can happen with sequential updates.

## Extra Credit

Try updating both values in a single assignment, such as increasing both by 10 after the swap.

## Summary

You have now learned how multiple assignment works in Go. Here is what you covered:

* Go supports assigning multiple values in a single statement.
* This is commonly used when working with function return values.
* Swapping values can be done safely without temporary variables.
* Multiple assignment evaluates all right hand values first.
* This feature helps keep Go code concise and expressive.

Multiple assignment is a small feature, but it appears everywhere in Go code, so becoming comfortable with it will make your programs cleaner and easier to reason about.


