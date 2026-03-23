# Protecting Shared State with sync.Mutex

**[Watch this lesson on YouTube](https://youtube.com/watch?v=D4avInS_vME)**

## Introduction

Goroutines allow multiple pieces of code to run at the same time.

However, when those goroutines access the same data, problems can occur.

In this lesson, you will learn how to protect shared state using `sync.Mutex`, and ensure your program behaves correctly when running concurrently.

## Uses / Use Cases

* Protecting shared variables.
* Preventing race conditions.
* Coordinating access to shared data.
* Ensuring consistent results in concurrent programs.

## The Problem

Consider this example:

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    counter := 0

    for i := 0; i < 5; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            counter++
        }()
    }

    wg.Wait()

    fmt.Println(counter)
}
```

You might expect the output to always be:

```
5
```

But in reality, the result can be inconsistent.

This happens because multiple goroutines are updating the same variable at the same time.

## Understanding the Issue

The operation:

```go
counter++
```

is not a single step.

It involves:

1. Reading the current value
2. Incrementing it
3. Writing it back

If two goroutines do this at the same time, updates can be lost.

This is called a **race condition**.

## Introducing sync.Mutex

A mutex ensures that only one goroutine can access a piece of code at a time.

You use it to protect critical sections of your program.

## Using Lock and Unlock

Here is the corrected example.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex

    counter := 0

    for i := 0; i < 5; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()

    fmt.Println(counter)
}
```

Now the output will consistently be:

```
5
```

## Why This Works

The mutex ensures that only one goroutine can execute the protected section at a time.

```go
mu.Lock()
counter++
mu.Unlock()
```

Other goroutines must wait until the lock is released.

This prevents overlapping updates.

## Using defer for Unlock

It is common to use `defer` to ensure the mutex is always unlocked.

```go
mu.Lock()
defer mu.Unlock()

counter++
```

This pattern helps avoid bugs where a mutex is not released.

## When to Use a Mutex

Use a mutex when:

* Multiple goroutines share the same data.
* You need to ensure safe access to that data.
* You are protecting simple state like counters, maps, or structs.

Channels can also coordinate goroutines, but they are better suited for passing data between them.

A mutex is often simpler when you just need to protect shared state.

## Challenge

Write a program that:

1. Creates a shared counter.
2. Starts 10 goroutines.
3. Each goroutine increments the counter.
4. Uses a `sync.Mutex` to ensure the final result is correct.

Verify that the final value is always 10.

## Solution

Great job if you attempted this challenge! Here is one possible solution.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex

    counter := 0

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            mu.Lock()
            defer mu.Unlock()

            counter++
        }()
    }

    wg.Wait()

    fmt.Println(counter)
}
```

## Summary

You have learned how to protect shared data using `sync.Mutex`.

* Concurrent access to shared data can cause race conditions.
* `sync.Mutex` ensures only one goroutine accesses a section at a time.
* Use `Lock` and `Unlock` to protect critical sections.
* `defer mu.Unlock()` helps avoid errors.
* Mutexes are useful when protecting shared state.

This is an essential tool for writing safe concurrent programs in Go.
