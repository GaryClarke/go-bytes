# Arrays in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=2fqAOMNv5C4)**

## Introduction

An array in Go is a fixed-size collection of elements, all of the same type, stored in a specific order.

Think of it as a row of boxes, where each box holds a value, and each value has a number (called an index) starting from 0.

Arrays are useful when you want to store multiple values and access them by position.

In this lesson, you'll learn how to:

* Declare an array
* Initialize it with values
* Access and update elements
* Check the array's length

## Uses / Use Cases

* Grouping related values of the same type
* Accessing items by index
* Performing operations on a set of values
* Preparing for slices, loops, and collections

## Example

Let's declare an array of strings that stores three book titles:

```go
package main

import "fmt"

func main() {
    var books [3]string

    books[0] = "Go in Action"
    books[1] = "The Go Programming Language"
    books[2] = "Learn Go with Tests"

    fmt.Println("First book:", books[0])
    fmt.Println("All books:", books)
    fmt.Println("Array length:", len(books))
}
```

Explanation:

* `var books [3]string`
  Declares an array named `books` that can hold 3 strings. The size is fixed at 3.

* `books[0] = ...`
  Assigns a value to the first element. Arrays in Go are **zero-indexed**, so the first item is index 0.

* `fmt.Println(...)`
  Prints individual elements or the full array.

* `len(books)`
  Returns the number of elements in the array. In this case, it returns 3.

You can also declare and initialize in one line:

```go
books := [3]string{"Go in Action", "The Go Programming Language", "Learn Go with Tests"}
```

## Expected Output

```
First book: Go in Action
All books: [Go in Action The Go Programming Language Learn Go with Tests]
Array length: 3
```

## Challenge

Declare an array that holds three integers: 100, 200, and 300.
Then print:

* The second number in the array
* The entire array

Try doing this using both individual assignments **and** a shorthand declaration.

## Solution

Here's one way to solve it:

```go
package main

import "fmt"

func main() {
    numbers := [3]int{100, 200, 300}

    fmt.Println("Second number:", numbers[1])
    fmt.Println("All numbers:", numbers)
}
```

Remember: index 1 refers to the second item, because indexes start at 0.

## Summary

You've learned:

* How to declare and initialize arrays in Go
* That arrays have a fixed size and a consistent type
* How to access elements using their index
* That Go indexes start at 0
* How to print individual values and the entire array
* How to check the array's length using `len()`

Arrays are a great stepping stone to slices, which offer more flexibility and are used more often in real Go code.

