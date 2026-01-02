# Using `make()` in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=FcdTJbIz5p0)**

## Introduction

In Go, some types need to be initialised before you can use them. Slices and maps fall into this category. While you can declare these types using `var`, they are not usable until they have been properly created.

The built-in `make()` function is used to create and initialise slices and maps. It allocates the necessary underlying data structures so the value is ready to use immediately.

In this lesson, you'll learn when to use `make()`, how it works for different types, and how length and capacity affect slices.

## Uses / Use Cases

* Creating slices with a specific length or capacity.
* Initialising maps so you can add key value pairs.
* Avoiding nil values that cause runtime panics.
* Controlling memory usage and performance.

## Using make() with Slices

Slices are one of the most common uses of `make()`.

```go
package main

import "fmt"

func main() {
    numbers := make([]int, 3)

    fmt.Println(numbers)
    fmt.Println("Length:", len(numbers))
    fmt.Println("Capacity:", cap(numbers))
}
```

Expected output:

```
[0 0 0]
Length: 3
Capacity: 3
```

Explanation:

* `make([]int, 3)` creates a slice of length 3.
* All elements are initialised to the zero value.
* Length and capacity are both set to 3.

### Length and Capacity

You can also specify a different capacity.

```go
package main

import "fmt"

func main() {
    numbers := make([]int, 3, 5)

    fmt.Println(numbers)
    fmt.Println("Length:", len(numbers))
    fmt.Println("Capacity:", cap(numbers))
}
```

Expected output:

```
[0 0 0]
Length: 3
Capacity: 5
```

Here:

* The slice starts with 3 usable elements.
* It has room to grow without reallocating.
* Appending values up to the capacity will reuse the same backing array.

## Using make() with Maps

Maps must be initialised before you can add entries. Declaring a map without `make()` creates a nil map, which cannot be written to.

```go
package main

import "fmt"

func main() {
    scores := make(map[string]int)

    scores["Alice"] = 10
    scores["Bob"] = 7

    fmt.Println(scores)
}
```

Expected output:

```
map[Alice:10 Bob:7]
```

Explanation:

* `make(map[string]int)` allocates the internal data structure.
* Once created, you can safely add, update, and remove entries.

If you try to write to a map that has not been created with `make()`, your program will panic.

## make() vs var

It is important to understand the difference between declaring a variable and creating it.

```go
var nums []int
var data map[string]int
```

Both of these are nil values. You cannot safely use them until you call `make()`.

Using `var` creates nil values that will cause a panic if you try to use them. `make()` creates properly initialised values that are safe to use immediately.

Using `make()` creates a usable value:

```go
nums = make([]int, 3)
data = make(map[string]int)
```

## Challenge

Write a program that:

1. Creates a slice of strings with length 2 and capacity 4 using `make()`.
2. Adds two names to the slice (use the first two positions: `names[0]` and `names[1]`).
3. Appends two more names to the slice.
4. Creates a map from string to int using `make()` and adds two entries.
5. Prints the slice, its length, its capacity, and the map.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func main() {
    names := make([]string, 2, 4)
    names[0] = "Alice"
    names[1] = "Bob"

    names = append(names, "John", "Sarah")

    scores := make(map[string]int)
    scores["Alice"] = 5
    scores["Bob"] = 3

    fmt.Println("Names:", names)
    fmt.Println("Length:", len(names))
    fmt.Println("Capacity:", cap(names))
    fmt.Println("Scores:", scores)
}
```

Expected output:

```
Names: [Alice Bob John Sarah]
Length: 4
Capacity: 4
Scores: map[Alice:5 Bob:3]
```

Explanation:

* The slice was created with length 2 and capacity 4 using `make([]string, 2, 4)`.
* Two names were added to positions 0 and 1.
* When two more names were appended using `append()`, the original slice's length was updated from 2 to 4.
* Since the capacity was 4, there was enough room to hold all four names, so `append()` modified the original slice in place rather than creating a new one.
* The slice now has length 4 and capacity 4, all using the same backing array that was allocated by `make()`.

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!  
The important thing is understanding how `make()` creates usable slices and maps that you can work with immediately.

## Summary

You have now learned how to use `make()` in Go. Here is what you covered:

* `make()` is used to create slices and maps.
* It allocates and initialises internal data structures.
* Slices created with `make()` have a defined length and capacity.
* Maps must be created with `make()` before use.
* Declaring a variable with `var` is not the same as creating it with `make()`.

Understanding `make()` is essential for working confidently with Go's core data structures.

