# Floats in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=rVqB3so80Mk)**

## Introduction

Go provides two floating-point types: `float32` and `float64`. These represent decimal numbers with different levels of precision.

In this lesson, you'll learn the difference between these types, when to use each one, and why Go defaults to `float64` for most floating-point operations.

Understanding floats helps you choose the right type for your calculations and avoid precision issues.

## Uses / Use Cases

* Working with decimal numbers like prices, measurements, or percentages
* Performing calculations that require fractional precision
* Choosing between precision and memory usage
* Understanding when Go infers `float64` automatically

## Example

```go
package main

import "fmt"

func main() {
    var price32 float32 = 19.99
    var price64 float64 = 19.99

    fmt.Println("float32:", price32)
    fmt.Println("float64:", price64)

    // Go infers float64 by default
    temperature := 22.5
    fmt.Printf("Inferred type: %T\n", temperature)
}
```

## Expected Output

```
float32: 19.99
float64: 19.99
Inferred type: float64
```

### Explanation

* `float32` uses 32 bits and provides about 7 decimal digits of precision.
* `float64` uses 64 bits and provides about 15 decimal digits of precision.
* When you write a decimal literal like `22.5`, Go infers `float64` by default.
* `float64` is the standard choice in Go because it offers better precision with minimal performance cost on modern hardware.

## More Examples

### Precision Comparison

```go
var small32 float32 = 0.123456789
var small64 float64 = 0.123456789

fmt.Println("float32:", small32) // May lose precision
fmt.Println("float64:", small64) // Preserves more digits
```

### When to Use Each

Use `float32` when:
* Memory is constrained (embedded systems, large arrays)
* You need less precision and want to save space

Use `float64` when:
* You need maximum precision (most common case)
* Working with financial calculations or scientific data
* You're unsure which to choose (default choice)

## Challenge

Create a program that:

1. Declares a `float32` variable for a price
2. Declares a `float64` variable for a temperature
3. Uses `fmt.Printf` with `%T` to print the type of both variables
4. Declares a third variable using `:=` with a decimal value and prints its type

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    var price float32 = 29.99
    var temp float64 = 25.5

    fmt.Printf("Price type: %T\n", price)
    fmt.Printf("Temperature type: %T\n", temp)

    height := 1.75
    fmt.Printf("Height type: %T\n", height)
}
```

If your solution works differently but still shows the types of all three variables, that's perfectly fine. The goal is to understand the difference between `float32` and `float64`, and see that Go infers `float64` by default.

## Summary

* Go provides `float32` (32-bit) and `float64` (64-bit) for decimal numbers.
* `float64` offers better precision and is Go's default for decimal literals.
* Use `float32` when memory is a concern, otherwise prefer `float64`.
* Decimal literals like `3.14` are inferred as `float64` by default.

Understanding these types helps you write more precise and efficient Go programs.

---



