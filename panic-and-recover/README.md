# Panic and Recover in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Go uses explicit error handling for most situations.

However, there are cases where a program cannot continue safely. In these situations, Go provides `panic`.

In this lesson, you will learn what `panic` does, how it propagates through your program, and how `recover()` can be used to handle it safely.

## Uses / Use Cases

* Stopping execution when something unrecoverable occurs.
* Catching unexpected failures at program boundaries.
* Preventing a program from crashing completely.

Most Go code uses errors, not panic.

## What is panic?

`panic` stops the normal flow of execution.

When a panic occurs:

* The current function stops immediately.
* Deferred functions still run.
* The panic moves up the call stack.
* If it is not recovered, the program exits.

Here is a simple example.

```go
package main

import "fmt"

func main() {
    fmt.Println("Starting")
    panic("something went wrong")
    fmt.Println("This will not run")
}
```

Output:

```
Starting
panic: something went wrong
```

## Panic Propagation

Panics do not stay in one function. They move up through the call stack.

```go
func main() {
    doWork()
}

func doWork() {
    panic("failure")
}
```

If no function handles the panic, the program will crash.

## Using recover

`recover()` allows a program to stop a panic and continue running.

However, it only works inside a deferred function.

Here is an example.

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()

    panic("something went wrong")

    fmt.Println("This will not run")
}
```

Output:

```
Recovered from: something went wrong
```

The program does not crash because the panic was recovered.

## Why recover must be deferred

`recover()` only works when called inside a deferred function.

This is because panics unwind the stack, and deferred functions run during that process.

Without `defer`, `recover()` will not catch the panic.

## When to Use panic

`panic` is not a replacement for normal error handling.

Use panic when:

* The program is in an invalid state.
* Continuing would cause incorrect behaviour.
* The issue cannot be handled locally.

For example:

* Invalid assumptions
* Corrupted state
* Programming errors

## Panic vs Errors in Servers

In most applications, especially servers, you should return errors rather than panic.

Errors:

* Are expected
* Can be handled cleanly
* Keep the program stable

Panic:

* Is for unexpected situations
* Should be rare

In server applications, panic is often caught at a higher level to prevent the entire server from crashing.

## Challenge

Write a program that:

1. Calls a function that triggers a panic.
2. Uses `defer` and `recover()` in `main` to handle it.
3. Prints a message when the panic is recovered.

Verify that the program does not crash.

## Solution

Great job if you attempted this challenge! Here is one possible solution.

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    doWork()

    fmt.Println("Program continues")
}

func doWork() {
    panic("unexpected error")
}
```

Output:

```
Recovered: unexpected error
Program continues
```

## Summary

You have learned how panic and recover work in Go.

* `panic` stops normal execution.
* Panics propagate up the call stack.
* Deferred functions still run during a panic.
* `recover()` can stop a panic when used inside `defer`.
* Most code should use errors instead of panic.
* Panic is reserved for unexpected or unrecoverable situations.

Understanding this behaviour helps you build more robust Go programs.
