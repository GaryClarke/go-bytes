# Understanding `nil` in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In Go, `nil` represents the absence of a value. It isn’t the same as `0`, an empty string, or `false`. It literally means *nothing*.

You’ll often see `nil` used with:

- **Pointers**
- **Slices**
- **Maps**
- **Interfaces**

In this lesson you’ll learn what `nil` means, when it’s used, and how to check for it.

## Uses / Use Cases

- Safely checking if a pointer, map, or slice is initialized
- Returning `nil` from a function to indicate “nothing to return”
- Avoiding panics when working with potentially uninitialized data

## Example

```go
package main

import "fmt"

func main() {
    var numbers []int
    var colors map[string]string
    var name *string

    fmt.Println(numbers == nil) // true
    fmt.Println(colors == nil)  // true
    fmt.Println(name == nil)    // true
}
```

### Explanation

- We declared a **slice**, a **map**, and a **pointer**.
- None were initialized, so their default value is `nil`.
- We use `== nil` to check whether the value is uninitialized.

### Common Scenarios

#### 1. Nil pointers

```go
var age *int
fmt.Println(age == nil) // true
```

#### 2. Nil maps

```go
var user map[string]string

if user == nil {
    user = map[string]string{}
}
```

#### 3. Nil slices

```go
var data []int

if data == nil {
    fmt.Println("data slice is nil")
}
```

#### 4. Nil interfaces

```go
var x interface{}
fmt.Println(x == nil) // true
```

## Challenge

Try declaring the following without initializing them:

- A `map[string]int`
- A `[]string` slice
- A `*float64` pointer

Then write `if` statements to check if each one is `nil`, and print a message if it is.

## Solution

Great job if you attempted this challenge! Here’s one possible solution:

```go
package main

import "fmt"

func main() {
    var scores map[string]int
    var names []string
    var value *float64

    if scores == nil {
        fmt.Println("scores map is nil")
    }

    if names == nil {
        fmt.Println("names slice is nil")
    }

    if value == nil {
        fmt.Println("value pointer is nil")
    }
}
```

If your code prints different messages but still checks for `nil` before using the values, that’s perfectly fine. The goal is to practise spotting uninitialised data.

## Summary

- `nil` means *no value*.
- It’s the default zero value for pointers, slices, maps, and interfaces.
- Always check for `nil` before using these types to avoid panics.
- You can use `== nil` for safe comparisons.

## Coming Up

In the next lesson we’ll explore **zero values**, the defaults assigned when you don’t initialise a variable.

---
