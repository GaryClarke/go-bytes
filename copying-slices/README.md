# The copy() Function for Slices in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=0ZR85YbZ36E)**

## Introduction

Slices in Go often share the same underlying data. This makes them efficient, but it can also lead to unexpected behaviour when changes in one slice affect another.

When you want a *separate copy* of slice data, Go provides the built-in `copy()` function. It allows you to copy elements from one slice into another, giving you control over when data is shared and when it is independent.

In this lesson, you will learn how `copy()` works, when to use it, and how it helps avoid subtle slice related bugs.

## Uses / Use Cases

* Creating independent copies of slices.
* Preventing changes to one slice from affecting another.
* Preparing slices for safe modification.
* Working with slices that share backing arrays.
* Writing clearer and safer slice manipulation code.

## The Problem: Assigning Slices Shares Data

When you assign a slice to another variable, both variables point to the same backing array. This means changes to one slice affect the other.

```go
package main

import "fmt"

func main() {
    original := []int{10, 20, 30}
    assigned := original

    assigned[0] = 99

    fmt.Println("original:", original)
    fmt.Println("assigned:", assigned)
}
```

Expected output:

```
original: [99 20 30]
assigned: [99 20 30]
```

Explanation:

* Assigning `original` to `assigned` does not create a copy.
* Both variables point to the same backing array.
* Modifying `assigned[0]` also changes `original[0]`.

This is the problem that `copy()` solves. When you need an independent copy that you can modify without affecting the original, you must use `copy()`.

## Basic Usage of copy()

The `copy()` function copies elements from one slice into another.

```go
package main

import "fmt"

func main() {
    source := []int{1, 2, 3}
    target := make([]int, len(source))

    copy(target, source)

    fmt.Println("source:", source)
    fmt.Println("target:", target)
}
```

Expected output:

```
source: [1 2 3]
target: [1 2 3]
```

Explanation:

* The first argument is the destination slice.
* The second argument is the source slice.
* `copy()` copies elements in order.
* The destination slice must already exist.

After copying, `source` and `target` contain the same values, but they are separate slices.

## Modifying the Copied Slice

Once a slice has been copied, modifying one does not affect the other.

```go
package main

import "fmt"

func main() {
    original := []int{10, 20, 30}
    duplicate := make([]int, len(original))

    copy(duplicate, original)

    duplicate[0] = 99

    fmt.Println("original:", original)
    fmt.Println("duplicate:", duplicate)
}
```

Expected output:

```
original: [10 20 30]
duplicate: [99 20 30]
```

Explanation:

* The slices no longer share the same backing array.
* Changes to `duplicate` do not affect `original`.

## How Many Elements Are Copied

The number of elements copied is the minimum of the two slice lengths.

```go
package main

import "fmt"

func main() {
    src := []int{1, 2, 3, 4}
    dst := make([]int, 2)

    n := copy(dst, src)

    fmt.Println("copied:", n)
    fmt.Println("dst:", dst)
}
```

Expected output:

```
copied: 2
dst: [1 2]
```

Explanation:

* `copy()` returns the number of elements copied.
* Only as many elements as the destination can hold are copied.
* Extra elements in the source are ignored.
* This behaviour is useful when you want to copy only part of a larger slice.

## copy() with Overlapping Slices

The `copy()` function also works safely with overlapping slices.

```go
package main

import "fmt"

func main() {
    data := []int{1, 2, 3, 4}

    copy(data[1:], data)

    fmt.Println(data)
}
```

Expected output:

```
[1 1 2 3]
```

Explanation:

* Go handles overlapping memory correctly.
* The behaviour is well defined and safe.
* This is useful for shifting elements inside a slice.
* In this example, copying `data` into `data[1:]` shifts the elements to the right, with the first element duplicated.

## Common Pattern: Copying Before Modifying

A very common pattern is copying a slice before making changes to ensure the original remains unchanged.

```go
package main

import "fmt"

func cloneSlice(input []int) []int {
    result := make([]int, len(input))
    copy(result, input)
    return result
}

func main() {
    original := []int{1, 2, 3}
    cloned := cloneSlice(original)

    cloned[0] = 99

    fmt.Println("original:", original)
    fmt.Println("cloned:", cloned)
}
```

Expected output:

```
original: [1 2 3]
cloned: [99 2 3]
```

This pattern ensures the caller's slice remains unchanged.

## Challenge

Write a program that:

1. Creates a slice of at least 3 integers.
2. Creates a copy of the slice using `copy()`.
3. Modifies the copied slice.
4. Prints both slices to show they are independent.

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
    numbers := []int{5, 10, 15}
    copied := make([]int, len(numbers))

    copy(copied, numbers)

    copied[1] = 99

    fmt.Println("numbers:", numbers)
    fmt.Println("copied:", copied)
}
```

Expected output:

```
numbers: [5 10 15]
copied: [5 99 15]
```

If you did not get this exactly right the first time, that is completely normal. The key idea is understanding when slices share data and when they do not.

## Summary

You have now learned how to use the `copy()` function with slices. Here is what you covered:

* `copy()` copies elements from one slice to another.
* The destination slice must be created first.
* Only as many elements as fit are copied.
* Copied slices do not share backing arrays.
* `copy()` helps avoid unexpected side effects when modifying slices.

Understanding `copy()` completes the core picture of how slices behave in Go and prepares you for more advanced slice and memory topics.
