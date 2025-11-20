# Type Conversion in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=O6QmV52q6dU)**

## Introduction

Go requires explicit type conversions when you want to change a value from one type to another. Unlike some languages, Go won't automatically convert types for you, which helps prevent bugs.

In this lesson, you'll learn how to convert between types, when conversions are needed, and common conversion patterns you'll use in Go programs.

Understanding type conversion helps you work with different data types and write code that compiles correctly.

## Uses / Use Cases

* Converting between numeric types (int to float64, float32 to float64)
* Converting strings to byte slices for processing
* Working with functions that require specific types
* Handling data from external sources that may be in different formats

## Example

```go
package main

import "fmt"

func main() {
    var count int = 42
    var price float64 = 19.99

    // Convert int to float64
    countAsFloat := float64(count)
    fmt.Println("Count as float:", countAsFloat)

    // Convert float64 to int (truncates decimal part)
    priceAsInt := int(price)
    fmt.Println("Price as int:", priceAsInt)

    // Convert string to byte slice
    text := "Hello"
    bytes := []byte(text)
    fmt.Println("Bytes:", bytes)
}
```

## Expected Output

```
Count as float: 42
Price as int: 19
Bytes: [72 101 108 108 111]
```

### Explanation

* `float64(count)` converts an `int` to `float64`. The syntax is `TypeName(value)`.
* `int(price)` converts a `float64` to `int`, truncating the decimal part (19.99 becomes 19).
* `[]byte(text)` converts a `string` to a slice of bytes. Each character becomes its byte value.
* Go requires explicit conversions. You can't mix types without converting them first.

## More Examples

### Converting Between Float Types

```go
var small float32 = 3.14
var large float64 = float64(small)
fmt.Println("Converted:", large)
```

### Converting Byte Slices Back to Strings

```go
bytes := []byte{72, 101, 108, 108, 111}
text := string(bytes)
fmt.Println("Text:", text) // "Hello"
```

### When Conversions Are Needed

Go won't automatically convert types, even for similar ones:

```go
var x int = 10
var y float64 = 3.5

// This won't compile without conversion:
// result := x + y  // Error!

// This works:
result := float64(x) + y
fmt.Println(result) // 13.5
```

## Challenge

Create a program that:

1. Declares two `int` variables: `a` with value `10` and `b` with value `3`
2. Performs integer division (`a / b`) and prints the result
3. Converts both `a` and `b` to `float64`
4. Performs float division with the converted values and prints the result

This demonstrates how type conversion lets you get fractional results from division.

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    // Integer division (truncates)
    resultInt := a / b
    fmt.Println("Integer division:", resultInt) // 3

    // Convert to float64 for fractional result
    aFloat := float64(a)
    bFloat := float64(b)
    resultFloat := aFloat / bFloat
    fmt.Println("Float division:", resultFloat) // 3.3333333333333335
}
```

If your solution works differently but still shows the difference between integer and float division using conversions, that's perfectly fine. The goal is to practice converting types to get the results you need.

## Summary

* Go requires explicit type conversions using the syntax `TypeName(value)`.
* Converting `float64` to `int` truncates the decimal part.
* You can convert strings to byte slices with `[]byte(string)`.
* You can convert byte slices back to strings with `string([]byte)`.
* Type conversions help you work with functions and operations that require specific types.

Understanding type conversion is essential for working with different data types in Go.

---

