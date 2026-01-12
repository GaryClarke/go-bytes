# String to Number Conversion in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=fUXPAnHX6nE)**

## Introduction

When working with user input, files, or external data, values often arrive as strings. To perform calculations or comparisons, you usually need to convert those strings into numbers. Likewise, numbers often need to be converted back into strings for display or output.

Go provides the `strconv` package for converting between strings and numeric types. It offers clear, explicit functions that return errors when conversion fails, making your programs safer and more predictable.

In this lesson, you will learn how to convert strings to numbers and numbers back to strings using the most commonly used functions in `strconv`.

## Uses / Use Cases

* Converting user input into numbers.
* Reading numeric values from files or environment variables.
* Formatting numbers for output.
* Validating numeric input safely.
* Avoiding runtime panics from invalid conversions.

## Converting Strings to Integers with strconv.Atoi

The simplest way to convert a string to an integer is `strconv.Atoi`.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    text := "42"

    number, err := strconv.Atoi(text)
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }

    fmt.Println("Number:", number)
}
```

Explanation:

* `strconv.Atoi` converts a string to an `int`.
* It returns the converted number and an error.
* If the string is not a valid integer, an error is returned.
* Always check the error before using the converted value.

## Expected Output

```
Number: 42
```

If `text` were `"abc"`, the conversion would fail and the error would be handled.

## Converting Numbers to Strings with strconv.Itoa

To convert an integer to a string, use `strconv.Itoa`.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    number := 123

    text := strconv.Itoa(number)
    fmt.Println("String:", text)
}
```

Explanation:

* `strconv.Itoa` converts an `int` to a string.
* It does not return an error because the conversion is always valid.
* This is useful when you need to format numbers for output or display.

## Challenge

Write a program that:

1. Reads a string value representing a number.
2. Converts it to an integer using `strconv.Atoi`.
3. Adds 10 to the number.
4. Converts the result back to a string using `strconv.Itoa`.
5. Prints the final string.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    input := "25"

    number, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid input")
        return
    }

    number += 10

    result := strconv.Itoa(number)
    fmt.Println("Result:", result)
}
```

Expected output:

```
Result: 35
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!
The important thing is understanding how to convert between strings and numbers and handle errors properly.

## Aside: More Control with strconv.ParseInt and strconv.FormatInt

If you need more control over conversions, Go provides additional functions.

For converting strings to integers with specific sizes or bases, use `strconv.ParseInt`:

```go
value, err := strconv.ParseInt("255", 10, 64)
```

The second argument is the base (10 means decimal), and the third argument is the bit size (64 means a 64-bit integer). The result type is `int64`.

For converting `int64` values to strings, use `strconv.FormatInt`:

```go
text := strconv.FormatInt(int64(1000), 10)
```

The second argument is the base for formatting (10 means decimal).

These functions are useful when you need specific integer sizes or bases, but for most cases, `strconv.Atoi` and `strconv.Itoa` are simpler and sufficient.

## Summary

You have now learned how to convert between strings and numbers in Go. Here is what you covered:

* The `strconv` package handles string and number conversions.
* `strconv.Atoi` converts strings to `int`.
* `strconv.Itoa` converts `int` to string.
* String to number conversions can fail and must be checked.
* `strconv.ParseInt` and `strconv.FormatInt` offer more control for specific needs.

These conversions are extremely common when working with input and output, so becoming comfortable with `strconv` will help you write more robust Go programs.

