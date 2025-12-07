# Understanding the `defer` Keyword in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=iqXVKW0r8NI)**

## Introduction

Go provides a keyword called `defer` that lets you schedule a function call to run later. When you defer a function, Go delays running it until the surrounding function completes. This is especially useful for cleanup tasks such as closing files, releasing resources, or printing final messages.

`defer` helps you keep cleanup code close to the resource it belongs to. It also reduces mistakes that happen when cleanup is far away from setup code.

In this lesson, you will learn how `defer` works, when it runs, why the evaluation of arguments happens immediately, and how to use it effectively.

## Uses / Use Cases

- Closing files after you open them.
- Releasing database connections or network resources.
- Unlocking mutexes in concurrent code.
- Ensuring cleanup happens even when a function returns early.
- Improving readability by keeping setup and cleanup code together.

## Example

Here is a simple example that shows when deferred functions run.

```go
package main

import "fmt"

func process() {
    fmt.Println("Starting work")
    defer fmt.Println("Cleaning up")
    fmt.Println("Doing work")
}

func main() {
    process()
}
```

Let's break this down:

- **Deferring a call**
  `defer fmt.Println("Cleaning up")` schedules the print to run later.

- **When deferred calls run**
  The deferred call runs after the function `process` completes, not when it is defined.

- **Output order**
  The order of printed messages shows exactly when the deferred function executes.

## Expected Output

```
Starting work
Doing work
Cleaning up
```

## Example: Arguments Are Evaluated Immediately

A common point of confusion is that deferred function *arguments* are evaluated right away, even though the function is not executed yet.

```go
package main

import "fmt"

func demo() {
    x := 10
    defer fmt.Println("Deferred value:", x)
    x = 20
    fmt.Println("Updated value:", x)
}

func main() {
    demo()
}
```

Output:

```
Updated value: 20
Deferred value: 10
```

The deferred call sees the value of `x` at the moment `defer` is executed, not when the function runs later.

## Example: Multiple Deferred Calls

If you defer several calls, they run in **last in, first out** order.

```go
package main

import "fmt"

func show() {
    for i := 1; i <= 3; i++ {
        defer fmt.Println("Deferred:", i)
    }
    fmt.Println("End of function")
}

func main() {
    show()
}
```

Output:

```
End of function
Deferred: 3
Deferred: 2
Deferred: 1
```

This stack like behaviour is important when managing complex cleanup sequences.

## Challenge

Write a function `saveReport()` that:

1. Prints `"Saving report"` at the start.
2. Defers two cleanup messages: `"Closing resources"` and `"Cleaning temporary data"`.
3. Prints `"Writing data"`.

Then call `saveReport()` inside `main()` and confirm the deferred messages run in reverse order.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func saveReport() {
    fmt.Println("Saving report")
    defer fmt.Println("Closing resources")
    defer fmt.Println("Cleaning temporary data")
    fmt.Println("Writing data")
}

func main() {
    saveReport()
}
```

Expected output:

```
Saving report
Writing data
Cleaning temporary data
Closing resources
```

## Summary

You have now seen how Go's `defer` keyword helps you manage cleanup tasks in a clear and reliable way. Here is what you learned:

- `defer` schedules a function call to run when the surrounding function completes.
- Deferred calls help you keep setup and cleanup code close together.
- Arguments to deferred calls are evaluated immediately.
- When multiple calls are deferred, they run in reverse order.
- `defer` is used heavily in real programs, especially when working with files, networks, and concurrency.

This simple but powerful feature plays an important role in writing safe and readable Go code.

