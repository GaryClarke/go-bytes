# Working with the strings Package in Go

**[Watch this lesson on YouTube](https://youtu.be/LLz8djgE20U)**

## Introduction

Go provides a built in package called `strings` that contains many helpful functions for working with text. Whether you are checking if a string contains certain characters, cleaning user input, or transforming text to a different case, the `strings` package gives you simple tools to do it.

Strings in Go are immutable. This means they cannot be changed once created. Every function in the `strings` package returns a new string, rather than modifying the original. Understanding this will help you work confidently with the functions provided.

In this lesson, you will learn the most commonly used functions in the `strings` package and see how to apply them in real world situations.

## Uses / Use Cases

- Checking if user input contains required text.
- Validating email addresses, filenames, or commands.
- Cleaning text from files or forms.
- Converting text to upper or lower case.
- Replacing parts of a string with new content.

## Example: Checking Strings

The `strings` package provides several functions that return booleans based on what the string contains.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    email := "alice@example.com"

    fmt.Println("Contains @:", strings.Contains(email, "@"))
    fmt.Println("Starts with alice:", strings.HasPrefix(email, "alice"))
    fmt.Println("Ends with .com:", strings.HasSuffix(email, ".com"))
}
```

Let's break this down:

- `strings.Contains` checks if a substring appears anywhere in the string.
- `strings.HasPrefix` checks how the string begins.
- `strings.HasSuffix` checks how the string ends.

These are very common checks when validating input.

### Expected Output

```
Contains @: true
Starts with alice: true
Ends with .com: true
```

## Example: Transforming Strings

Here are some functions that return a new version of the string.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    name := "  Sarah  "

    cleaned := strings.TrimSpace(name)
    upper := strings.ToUpper(cleaned)
    lower := strings.ToLower(cleaned)

    fmt.Println("Cleaned:", cleaned)
    fmt.Println("Upper:", upper)
    fmt.Println("Lower:", lower)
}
```

Explanation:

- `strings.TrimSpace` removes leading and trailing whitespace.
- `strings.ToUpper` converts to upper case.
- `strings.ToLower` converts to lower case.
- The original `name` string is unchanged because strings are immutable.

### Expected Output

```
Cleaned: Sarah
Upper: SARAH
Lower: sarah
```

## Example: Replacing Text

You can replace parts of a string using `strings.Replace`.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "hello bob bob bob"

    replacedOnce := strings.Replace(s, "bob", "john", 1)
    replacedAll := strings.Replace(s, "bob", "john", -1)

    fmt.Println("Replace one:", replacedOnce)
    fmt.Println("Replace all:", replacedAll)
}
```

Explanation:

- The final argument controls how many replacements to make.
- Use `1` to replace the first match.
- Use `-1` to replace all occurrences.
- The original string remains unchanged.

### Expected Output

```
Replace one: hello john bob bob
Replace all: hello john john john
```

## Challenge

Write a program that processes a string using at least three different `strings` functions. For example:

1. Trim any extra spaces.
2. Check whether it ends with `.txt`.
3. Convert it to upper case.

Choose any string you like, and print the results of each step. Make sure you use at least three different `strings` functions.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    filename := "  report.txt  "

    trimmed := strings.TrimSpace(filename)
    valid := strings.HasSuffix(trimmed, ".txt")
    upper := strings.ToUpper(trimmed)

    fmt.Println("Trimmed:", trimmed)
    fmt.Println("Valid .txt:", valid)
    fmt.Println("Upper:", upper)
}
```

### Expected Output

```
Trimmed: report.txt
Valid .txt: true
Upper: REPORT.TXT
```

## Summary

You have now learned the most useful functions in Go's `strings` package. Here is what you covered:

- The `strings` package provides tools for searching, checking, cleaning, and transforming text.
- Strings are immutable, so functions return new strings rather than modifying the original.
- `Contains`, `HasPrefix`, and `HasSuffix` return booleans.
- `ToUpper`, `ToLower`, `TrimSpace`, and `Replace` return new strings.
- These functions are common in real programs, especially when processing input.

As your programs grow, you will use the `strings` package often to make text handling simple and reliable.

