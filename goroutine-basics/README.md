# Goroutines Basics in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=vIzQ1NgC2-Q)**

## Introduction

Go makes it easy to run code concurrently. Instead of dealing with threads directly, Go provides *goroutines*, which are lightweight units of execution managed by the Go runtime.

A goroutine runs a function at the same time as other code. Starting a goroutine is simple, but understanding when and why to use one is important.

In this lesson, you learn what goroutines are, how to start them using the `go` keyword, and what happens when code runs concurrently.

## Uses / Use Cases

* Running work in the background.
* Performing multiple tasks at the same time.
* Keeping programs responsive.
* Handling concurrent operations such as IO or waiting.
* Preparing for communication with channels.

## What Is a Goroutine?

A goroutine is a function that runs concurrently with other code in the same program.

Goroutines are:

* Lightweight compared to OS threads.
* Managed by the Go runtime.
* Cheap to create and destroy.
* Designed to scale to many concurrent tasks.

You do not create or manage threads directly. Go handles that for you.

## Starting a Goroutine with go

You start a goroutine by placing the `go` keyword before a function call.

```go
package main

import "fmt"

func sayHello() {
    fmt.Println("Hello from a goroutine")
}

func main() {
    go sayHello()

    for i := 0; i < 1000000; i++ {
        // Do some work
    }
    fmt.Println("Hello from main")
}
```

Expected output (order may vary):

```
Hello from a goroutine
Hello from main
```

Or:

```
Hello from main
Hello from a goroutine
```

Explanation:

* `sayHello` is started as a goroutine.
* When you start a goroutine, `main` continues immediately without waiting.
* The loop in `main` takes time to complete, giving the goroutine a chance to run.
* When `main` returns, the program exits, even if goroutines are still running.
* Without the loop, the program might exit before the goroutine prints its message.
* The order of output is not guaranteed because both run concurrently.

You can also start an anonymous function as a goroutine:

```go
go func() {
    fmt.Println("Running in the background")
}()
```

## Goroutines Run Independently

When a goroutine is started, it runs independently of the function that started it. This means the program may exit before the goroutine finishes running.

```go
package main

import "fmt"

func main() {
    go func() {
        fmt.Println("This may not print")
    }()
}
```

Explanation:

* The goroutine starts, but `main` exits immediately.
* The program ends before the goroutine has a chance to run.
* This is why coordination is important.

Channels and other synchronisation tools help solve this problem.

## Goroutines vs Function Calls

A normal function call:

```go
sayHello()
```

Runs synchronously. The caller waits until it finishes.

A goroutine call:

```go
go sayHello()
```

Runs concurrently. The caller continues immediately.

This single keyword makes a big difference in how your program behaves.

## When Should You Use a Goroutine?

Use a goroutine when:

* You want work to run concurrently.
* The work can happen independently.
* You do not need the result immediately.
* The task might block or wait.

Do not use goroutines just to make code "faster". Concurrency is about structure, not speed.

## Challenge

Write a program that:

1. Defines a function that prints a message.
2. Starts that function as a goroutine.
3. Does some work in `main` (like a loop) before printing another message.

Run the program multiple times and observe the output order.

Run it using:

```
go run main.go
```

## Solution

```go
package main

import "fmt"

func worker() {
    fmt.Println("Worker running")
}

func main() {
    go worker()

    for i := 0; i < 1000000; i++ {
        // Do some work
    }
    fmt.Println("Main running")
}
```

Expected output (order may vary):

```
Worker running
Main running
```

Or:

```
Main running
Worker running
```

You may see different output orders each time you run the program. This is expected because goroutines run concurrently. The loop gives the goroutine time to run while `main` is busy, ensuring both messages appear.

## Summary

You have now learned the basics of goroutines in Go. Here is what you covered:

* Goroutines are lightweight concurrent functions.
* They are started using the `go` keyword.
* Goroutines run independently of the caller.
* The program exits when `main` finishes.
* Output order is not guaranteed.

This lesson sets the stage for channels, which allow goroutines to communicate and coordinate safely.
