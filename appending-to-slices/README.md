# Appending to Slices in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=7yWNfST7zMQ)**

## Introduction

Slices are one of the most important data structures in Go, and appending values to slices is something you will do constantly. Go provides the built-in `append` function to add elements to a slice.

While `append` looks simple on the surface, it has important behaviour around capacity and memory that is worth understanding. Knowing how `append` works will help you avoid subtle bugs and write more predictable code.

In this lesson, you will learn how to append values to slices, how capacity affects append, and when Go creates a new backing array.

## Uses / Use Cases

* Growing slices dynamically.
* Building lists of values over time.
* Collecting results in loops.
* Understanding slice behaviour when passed between functions.
* Avoiding unexpected side effects.

## Basic Usage of append

The simplest use of `append` is adding a single element.

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}

    numbers = append(numbers, 4)

    fmt.Println(numbers)
}
```

Expected output:

```
[1 2 3 4]
```

Explanation:

* `append` returns the updated slice.
* You must assign the result back to a variable.
* The original slice value is not modified unless you capture the return value.

This assignment step is important and easy to forget, which is a very common beginner mistake.

## Appending Multiple Values

You can append more than one value at a time.

```go
numbers = append(numbers, 5, 6, 7)
```

All values are added in order to the slice.

## Length and Capacity

Every slice has a length and a capacity.

```go
package main

import "fmt"

func main() {
    nums := make([]int, 0, 3)

    fmt.Println("len:", len(nums), "cap:", cap(nums))

    nums = append(nums, 1)
    nums = append(nums, 2)
    nums = append(nums, 3)

    fmt.Println("len:", len(nums), "cap:", cap(nums))
}
```

Expected output:

```
len: 0 cap: 3
len: 3 cap: 3
```

Explanation:

* Length is how many elements are in the slice.
* Capacity is how many elements it can hold before growing.
* As long as capacity is available, `append` reuses the same backing array.

## When append Creates a New Backing Array

When a slice runs out of capacity, `append` allocates a new backing array and copies the existing elements into it.

```go
package main

import "fmt"

func main() {
    nums := make([]int, 0, 2)

    nums = append(nums, 1, 2)
    fmt.Println("Before:", nums, "cap:", cap(nums))

    nums = append(nums, 3)
    fmt.Println("After:", nums, "cap:", cap(nums))
}
```

Expected output (capacity may vary):

```
Before: [1 2] cap: 2
After: [1 2 3] cap: 4
```

Explanation:

* The slice was full when appending `3`.
* Go allocated a new backing array with more capacity.
* The slice now points to a different array.

The exact growth strategy is handled by Go and should not be relied on. The important thing to understand is that a new array is created when capacity is exceeded, not the specific capacity value.

This behaviour becomes especially important when slices are passed to functions or shared between variables, which we cover in detail in the "Slices and Reference Semantics" lesson.

## Challenge

Write a program that:

1. Creates a slice with an initial capacity of 2 using `make([]int, 0, 2)`.
2. Uses a loop to append values 1 through 5, one at a time.
3. Prints the slice, length, and capacity after each append.

Observe when the capacity changes and notice how it grows.

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
    values := make([]int, 0, 2)

    for i := 1; i <= 5; i++ {
        values = append(values, i)
        fmt.Println("values:", values, "len:", len(values), "cap:", cap(values))
    }
}
```

Example output:

```
values: [1] len: 1 cap: 2
values: [1 2] len: 2 cap: 2
values: [1 2 3] len: 3 cap: 4
values: [1 2 3 4] len: 4 cap: 4
values: [1 2 3 4 5] len: 5 cap: 8
```

If your capacity values differ slightly, that is fine. The important part is noticing when growth happens and understanding that a new backing array was created.

If you did not get it exactly right the first time, that is completely normal. The important thing is understanding how `append` works with capacity.

## Summary

You have now learned how appending to slices works in Go. Here is what you covered:

* `append` adds elements to a slice and returns the updated slice.
* You must always assign the result of `append`.
* Slices have both length and capacity.
* `append` reuses the backing array when capacity allows.
* When capacity is exceeded, a new backing array is created.
* Understanding this helps avoid unexpected behaviour.
