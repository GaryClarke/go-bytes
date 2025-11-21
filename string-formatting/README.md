# String Formatting in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=GORNyBnoMPI)**

## Introduction

Go gives you powerful tools for formatting strings using the `fmt` package.

In this lesson, you'll learn the difference between `fmt.Println`, `fmt.Printf`, and `fmt.Sprintf`, and how to use format verbs like `%s`, `%d`, and `%v`.

These formatting tools help you control how values are displayed, useful when printing messages, building output strings, or logging structured data.

## Uses / Use Cases

* Formatting numbers and strings for display
* Building dynamic strings with values
* Controlling output layout (spacing, padding, precision)
* Creating formatted logs or labels

## Example

### 1. `fmt.Println`

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 35

    fmt.Println("Name:", name)  // Name: Alice
    fmt.Println("Age:", age)    // Age: 35
}
```

This is the simplest way to print values. It automatically adds spaces between arguments and a newline at the end.

### 2. `fmt.Printf`

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 35

    fmt.Printf("Name: %s\n", name)  // Name: Alice
    fmt.Printf("Age: %d\n", age)      // Age: 35
}
```

`fmt.Printf` lets you control the formatting using **format verbs**. It doesn't automatically add spaces between values, so you decide how it looks.

### 3. `fmt.Sprintf`

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 35

    message := fmt.Sprintf("My name is %s and I'm %d years old.", name, age)
    fmt.Println(message)  // My name is Alice and I'm 35 years old.
}
```

`fmt.Sprintf` works like `Printf`, but **returns** the string instead of printing it. Great for building messages to use elsewhere (logs, responses, etc.).

## Common Format Verbs

| Verb  | Description               | Example                        |
| ----- | ------------------------- | ------------------------------ |
| `%s`  | String                    | `"Go"` → `Go`                  |
| `%d`  | Integer (decimal)         | `42` → `42`                    |
| `%f`  | Float (decimal)          | `3.14` → `3.140000`            |
| `%v`  | Default format (any type) | `true` → `true`                |
| `%T`  | Print type of value       | `"text"` → `string`            |
| `%+v` | Struct with field names   | `{Name: "Alice"}` → `Name:Alice` |

## Challenge

1. Create variables for a product:
   * Name: `"Notebook"`
   * Price: `9.99`

2. Use `fmt.Printf` to print: `Notebook costs $9.99`

3. Use `fmt.Sprintf` to build the same string and store it in a variable.

4. Print the result using `fmt.Println`.

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    name := "Notebook"
    price := 9.99

    // Using fmt.Printf
    fmt.Printf("%s costs $%.2f\n", name, price)  // Notebook costs $9.99

    // Using fmt.Sprintf
    message := fmt.Sprintf("%s costs $%.2f", name, price)
    fmt.Println(message)  // Notebook costs $9.99
}
```

`%.2f` formats the float with 2 decimal places.

If your solution works differently but still formats the output correctly, that's perfectly fine. The goal is to practice using format verbs to control how values are displayed.

## Summary

You now know how to format output in Go using:

* `fmt.Println` for simple printing with spaces and newlines
* `fmt.Printf` for custom formatting with format verbs
* `fmt.Sprintf` to build formatted strings for later use

These tools give you full control over how your program displays information, both in the terminal and behind the scenes.

---


