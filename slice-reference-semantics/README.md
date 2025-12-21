# Slices and Reference Semantics in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=YAYvCD-lt3Q)**

## Introduction

Slices let you work with sequences of values in a flexible way. Now it is time to explore how slices behave when you pass them to functions or assign them to new variables.

Slices look like ordinary values, but they do not behave like integers or structs. A slice contains a small header that points to a backing array where the actual data lives. When you pass a slice to a function, Go copies only the header. The backing array is shared.

This means that modifying a slice inside a function usually modifies the original. However, appending to a slice can behave differently depending on capacity. Understanding this behaviour is essential for writing predictable Go programs.

## Uses / Use Cases

* Sharing data across functions without copying.
* Allowing helper functions to update slices.
* Working with large sequences efficiently.
* Understanding when append affects the original slice and when it does not.

## Example: Modifying Elements Through a Function

Here is a simple slice, and a function that doubles the first element.

```go
package main

import "fmt"

func doubleFirst(values []int) {
    values[0] = values[0] * 2
}

func main() {
    nums := []int{10, 20, 30}

    doubleFirst(nums)

    fmt.Println("After:", nums)
}
```

Expected output:

```
After: [20 20 30]
```

Explanation:

* The function receives a copy of the slice header.
* The header still points to the same backing array.
* Updating `values[0]` updates the original slice in `main`.

This is reference behaviour, meaning both the original slice and the function parameter share the same backing array.

## Example: Assigning Slices Creates Shared References

Assigning one slice to another variable does not create a copy of the array.

```go
package main

import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := a

    b[1] = 99

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

Expected output:

```
a: [1 99 3]
b: [1 99 3]
```

Both slices point to the same backing array, so changes through either name affect both.

## Example: When Append Creates a New Array

If there is not enough capacity, `append` creates a new backing array and copies the data into it. After that, the two slices no longer share data.

```go
package main

import "fmt"

func main() {
    a := []int{1, 2} // capacity is 2
    b := a

    b = append(b, 3) // triggers new array
    b[0] = 999

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

Expected output:

```
a: [1 2]
b: [999 2 3]
```

Explanation:

* `a` and `b` originally pointed to the same array.
* `append` needed more space, so Go allocated a new array for `b`.
* After that, modifying `b` does not affect `a`.

This is the key subtlety when working with slices.

## Challenge

Write a program that:

1. Creates a slice of integers with length 2 and capacity 2 using `make([]int, 2, 2)`. Then assign values to the first two elements (e.g., 10 and 20).
2. Passes the slice to a function that modifies one of the elements (e.g., change the first element to 100).
3. Passes the slice to another function that appends a value (e.g., 50) and returns the new slice.
4. Print the original slice and the new slice after the append so you can see whether the original changed.

Because your slice starts with capacity 2, appending will trigger a new array allocation, so the original slice will remain unchanged after the append.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func modify(values []int) {
    values[0] = 100
}

func addValue(values []int) []int {
    return append(values, 50)
}

func main() {
    nums := make([]int, 2, 2)
    nums[0], nums[1] = 10, 20

    modify(nums)
    fmt.Println("After modify:", nums)

    nums2 := addValue(nums)
    fmt.Println("Original after append:", nums)
    fmt.Println("New slice after append:", nums2)
}
```

Possible output:

```
After modify: [100 20]
Original after append: [100 20]
New slice after append: [100 20 50]
```

The appended slice has a new backing array, so further changes to `nums2` will not affect `nums`.

If you did not get it exactly right the first time, that is completely normal. Slices and capacity can be tricky at first. The important thing is understanding how the backing array is shared until capacity runs out.

## Summary

You have now learned how slices behave with reference semantics in Go. Here is what you covered:

* Slices contain a header that points to a backing array.
* Passing a slice to a function passes a copy of the header, not the data.
* Modifying elements through a slice updates the shared backing array.
* Assigning a slice to another variable creates another reference to the same data.
* Appending to a slice may or may not create a new array, depending on capacity.
* When a new array is created, the slices no longer share data.

Understanding these behaviours helps you write efficient, predictable Go code and prepares you for more complex slice operations later.

