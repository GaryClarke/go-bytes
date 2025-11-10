# Math in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VntgJANyYyc)**

## Introduction

Go supports standard mathematical operations like addition, subtraction, multiplication, and division.
It also provides more advanced functions through the built-in `math` package.

In this lesson, you'll learn how to:

* Use basic math operators
* Perform float vs integer division
* Use Go's `math` package for square roots and constants like π

## Uses / Use Cases

* Doing calculations in business logic
* Processing user input or numerical data
* Using math constants like Pi in geometry-related code
* Performing square root or power operations

## Example

```go
package main

import "fmt"
import "math"

func main() {
    // Basic arithmetic
    a := 10
    b := 3

    fmt.Println("a + b =", a+b)
    fmt.Println("a - b =", a-b)
    fmt.Println("a * b =", a*b)
    fmt.Println("a / b =", a/b) // integer division
    fmt.Println("a % b =", a%b) // modulus (remainder)

    // Float division
    x := 10.0
    y := 3.0
    fmt.Println("x / y =", x/y)

    // Using parentheses for order of operations
    c := 12.0
    d := 5.0
    e := 2.0
    result := c + (d * e)
    fmt.Println("Result using parentheses:", result)

    // Using math package
    fmt.Println("Pi:", math.Pi)
    fmt.Println("Square root of 16:", math.Sqrt(16))
}
```

## Explanation

Go supports the standard arithmetic operators:

* `+` (addition)
* `-` (subtraction)
* `*` (multiplication)
* `/` (division)
* `%` (modulus, or remainder)

Some important notes:

* **Integer division**: When both operands are integers (e.g., `10 / 3`), Go performs **integer division**, which **truncates** the result. So `10 / 3` evaluates to `3`, not `3.333...`.
* **Floating-point division**: To get a fractional result, you must **use floats** (e.g., `10.0 / 3.0`) or **convert** integers to `float64` using `float64(x)` before dividing.
* **Modulus**: `%` gives the remainder after integer division, like `10 % 3 = 1`.
* **Parentheses**: Use parentheses to control the order of operations. For example, `result := c + (d * e)` ensures the multiplication happens before the addition.

Go also provides common mathematical constants and functions via the `math` package:

* `math.Pi` – the value of π
* `math.Sqrt(x)` – square root of `x`
* Other functions: `math.Pow`, `math.Abs`, `math.Sin`, etc.

To use these, import the math package at the top of your file:

```go
import "math"
```

## Expected Output

```
a + b = 13
a - b = 7
a * b = 30
a / b = 3
a % b = 1
x / y = 3.3333333333333335
Pi: 3.141592653589793
Square root of 16: 4
```

## Challenge

Write a program that:

1. Declares two `float64` variables: `price` and `taxRate`
2. Calculates the total cost by adding tax to the price (hint: tax is calculated as `price * taxRate`)
3. Prints the total using `fmt.Println`

## Solution

Great job if you attempted this challenge! Here’s one possible solution:

```go
package main

import "fmt"

func main() {
    price := 19.99
    taxRate := 0.08

    total := price + (price * taxRate)

    fmt.Println("Total cost:", total)
}
```

## Summary

In this lesson, you learned how to:

* Use arithmetic operators in Go
* Understand the difference between integer and float division
* Use the `math` package for more advanced calculations

These math skills will help you build everything from pricing calculators to data visualizations and statistical tools.

---
