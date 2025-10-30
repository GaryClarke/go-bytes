# Go Slices - Flexible Lists in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=pdvJVbUnMjc)**

## Introduction

In Go, **slices** are the most common way to store a list of values. You can think of a slice as a flexible, resizable version of an array.

Where arrays have a fixed size, slices grow or shrink as needed. They're built on top of arrays but offer far more convenience and power.

In this lesson, you'll learn what slices are, how to create them, how to access and update their values, and how to work with them in real code.

## Uses / Use Cases

* Grouping values like a list of names or numbers
* Dynamically adding or removing elements
* Iterating over collections of items
* Passing multiple values into a function

## Example

Here's a simple Go program using a slice of strings:

```go
package main

import "fmt"

func main() {
    names := []string{"Alice", "Bob", "Charlie"}

    fmt.Println("First name:", names[0])
    fmt.Println("All names:", names)

    names = append(names, "Diana")
    fmt.Println("After append:", names)
}
```

### Let's break it down:

* `[]string{"Alice", "Bob", "Charlie"}`
  This creates a slice of strings with 3 initial elements. The square brackets are empty (`[]`), meaning it's a slice, not a fixed-size array.

* `names[0]`
  This accesses the first item in the slice. Slice indexes start at 0.

* `append(names, "Diana")`
  This returns a new slice with `"Diana"` added to the end. You assign it back to `names` to keep the updated list.

## Expected Output

```
First name: Alice
All names: [Alice Bob Charlie]
After append: [Alice Bob Charlie Diana]
```

## Challenge

Create a slice of 3 integers and print each of the following:

* The full slice
* The first value
* The length of the slice

Then, append a fourth number to the slice and print the result.

You can use `len()` to get the length of a slice.

## Solution

Here's one way to solve the challenge:

```go
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}

    fmt.Println("Original slice:", numbers)
    fmt.Println("First number:", numbers[0])
    fmt.Println("Length:", len(numbers))

    numbers = append(numbers, 40)
    fmt.Println("After append:", numbers)
}
```

If your numbers or messages differ, that's fine. The goal is to create and use a slice and see how it behaves when you append a new value.

## Summary

You've learned:

* Slices are flexible, resizable lists in Go
* They're created using square brackets with no size: `[]type{...}`
* You can access values using index notation: `slice[0]`
* The `append()` function adds new values to a slice
* The `len()` function returns the number of elements

Slices are the preferred way to work with collections in Go. You'll see them used often when looping through items or passing data between functions.

This sets the stage for working with loops and building more complex data processing logic.

