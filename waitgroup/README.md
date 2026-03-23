# Waiting for Goroutines with sync.WaitGroup

**[Watch this lesson on YouTube](https://youtube.com/watch?v=Ng4Nw0iLkfI)**

## Introduction

Goroutines allow Go programs to run work concurrently.

However, a common issue appears when the `main` function finishes before your goroutines do. When that happens, the program exits and the goroutines stop immediately.

To solve this, Go provides `sync.WaitGroup`.

A WaitGroup allows your program to wait until a set of goroutines has finished running.

## Uses / Use Cases

* Waiting for multiple goroutines to complete.
* Coordinating concurrent tasks.
* Ensuring background work finishes before the program exits.

WaitGroups are often used in systems that process work in parallel.

## The Problem

Consider this program:

```go
package main

import "fmt"

func main() {
    go func() {
        fmt.Println("Working...")
    }()
}
```

You might expect `"Working..."` to print.

But in many cases the program exits before the goroutine runs.

The `main` function finishes immediately, and the program stops.

We need a way to tell Go:

> Wait until this work is finished before exiting.

## Introducing sync.WaitGroup

The `sync` package provides a type called `WaitGroup`.

A WaitGroup keeps track of how many goroutines are still running and allows the program to wait until they complete.

The basic workflow is:

1. Tell the WaitGroup how many goroutines to wait for.
2. Each goroutine signals when it finishes.
3. The program waits until all work is done.

## Basic Example

Here is a simple example using a WaitGroup.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(1)

    go func() {
        defer wg.Done()
        fmt.Println("Working...")
    }()

    wg.Wait()
}
```

Now the program waits until the goroutine finishes.

## Understanding Add, Done, and Wait

A WaitGroup works through three methods.

### Add

`Add(n)` tells the WaitGroup how many goroutines to wait for.

```go
wg.Add(1)
```

If you start three goroutines, you would use:

```go
wg.Add(3)
```

### Done

`Done()` signals that a goroutine has finished.

```go
wg.Done()
```

It decreases the internal counter by one.

A common pattern is:

```go
defer wg.Done()
```

This ensures the goroutine always signals completion.

### Wait

`Wait()` blocks execution until the counter reaches zero.

```go
wg.Wait()
```

When all goroutines call `Done()`, the program continues.

## Waiting for Multiple Goroutines

WaitGroups become more useful when working with several goroutines.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    tasks := []string{"task1", "task2", "task3"}

    for _, task := range tasks {
        wg.Add(1)

        go func(t string) {
            defer wg.Done()
            fmt.Println("Processing", t)
        }(task)
    }

    wg.Wait()
}
```

The program waits until all three goroutines finish.

## Challenge

Write a program that:

1. Creates a slice of numbers.
2. Starts a goroutine for each number.
3. Prints the number inside the goroutine.
4. Uses a `sync.WaitGroup` so the program waits for all goroutines to finish.

Verify that all numbers are printed before the program exits.

## Solution

Great job if you attempted this challenge! Here is one possible solution.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    numbers := []int{1, 2, 3, 4}

    var wg sync.WaitGroup

    for _, n := range numbers {
        wg.Add(1)

        go func(num int) {
            defer wg.Done()
            fmt.Println(num)
        }(n)
    }

    wg.Wait()
}
```

Each goroutine signals completion using `Done()`, and the program waits until all of them finish.

## Summary

You have learned how to wait for goroutines using `sync.WaitGroup`.

* Goroutines may finish after `main` exits.
* `sync.WaitGroup` allows a program to wait for concurrent work.
* `Add` tells the WaitGroup how many tasks to expect.
* `Done` signals that a task has finished.
* `Wait` blocks until all tasks complete.

This pattern is widely used when coordinating concurrent work in Go programs.
