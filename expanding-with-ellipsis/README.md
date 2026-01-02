# Expanding Slices with `...` in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VUq5h4GHJSw)**

## Introduction

In Go, variadic functions accept a variable number of arguments. You have already learned how to define and call these functions. Sometimes, however, you already have your values stored in a slice and want to pass them into a variadic function.

This is where the ellipsis operator, written as `...`, comes in.

The `...` operator takes a slice and spreads its elements out as individual arguments. Without it, Go treats the slice as a single value, which usually causes a compile error.

In this lesson, you'll learn how slice expansion works, when it is required, when it fails, and how to forward variadic arguments correctly.

## Uses / Use Cases

* Passing a slice into a variadic function.
* Forwarding variadic arguments from one function to another.
* Avoiding common compile errors when working with variadic parameters.
* Writing flexible helper functions that wrap other variadic functions.

## Example: When Slice Expansion Is Required

Here is a simple variadic function:

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    values := []int{1, 2, 3}

    fmt.Println(sum(values...))
}
```

Expected output:

```
6
```

Explanation:

* `sum` expects individual `int` arguments.
* `values` is a slice of `int`.
* `values...` expands the slice so each element is passed separately.

Without the `...`, this code would not compile.

## Example: What Happens Without `...`

If you forget to expand the slice, Go will reject the call.

```go
sum(values)
```

This fails because:

* `sum` expects `int, int, int`
* `values` is a single value of type `[]int`

Go does not automatically unpack slices. You must be explicit.

## Example: Forwarding Variadic Arguments

A very common pattern is forwarding variadic arguments from one function to another.

```go
package main

import "fmt"

func logMessages(msgs ...string) {
    for _, m := range msgs {
        fmt.Println("Log:", m)
    }
}

func logAll(msgs ...string) {
    logMessages(msgs...)
}

func main() {
    logAll("start", "processing", "done")
}
```

Expected output:

```
Log: start
Log: processing
Log: done
```

Explanation:

* `logAll` receives a variadic parameter.
* Inside the function, `msgs` is a slice.
* To pass those values on, you must expand the slice again using `msgs...`.

Without expansion, forwarding would fail to compile.

## Example: Expansion Only Works with Variadic Parameters

The ellipsis operator only works when calling a variadic function.

```go
func printTwo(a string, b string) {
    fmt.Println(a, b)
}
```

This will not work:

```go
values := []string{"hello", "world"}
printTwo(values...) // not allowed
```

Why?

* `printTwo` does not accept a variadic parameter.
* Slice expansion is only valid when the function expects a variadic argument.

This is an important rule to remember.

## Challenge

Write a variadic function called `printAll` that prints each string on a new line.

Then:

1. Create a slice of strings.
2. Call `printAll` using slice expansion.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func printAll(items ...string) {
    for _, item := range items {
        fmt.Println(item)
    }
}

func main() {
    words := []string{"one", "two", "three"}

    printAll(words...)
}
```

Expected output:

```
one
two
three
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!  
The important thing is understanding how `...` expands slices when calling variadic functions.

## Summary

You have now learned how to expand slices using the `...` operator. Here is what you covered:

* The ellipsis operator expands a slice into individual arguments.
* Expansion is required when passing a slice to a variadic function.
* Forwarding variadic arguments also requires expansion.
* Expansion only works with variadic parameters.
* Forgetting `...` is a common source of compile errors.

This small operator plays an important role in writing flexible and reusable Go functions.

