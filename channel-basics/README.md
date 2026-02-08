# Channels in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=WFugE45bUZg)**

## Introduction

Go has built-in support for concurrency. One of the core tools for communicating between concurrent parts of a program is the *channel*.

A channel allows one part of a program to send data to another. Rather than sharing memory directly, Go encourages sharing data by communicating through channels.

In this lesson, you will learn how to create channels, send values into them, and receive values from them.

## Uses / Use Cases

* Sending data between goroutines.
* Coordinating work between concurrent tasks.
* Passing results from one part of a program to another.
* Avoiding shared state and locks.
* Writing clearer concurrent code.

## Creating and Using Channels

Channels must be created before they can be used. This is done with the built-in `make()` function.

```go
package main

import "fmt"

func main() {
    // Create a channel that can carry string values
    messages := make(chan string)

    // Start a goroutine that sends a message
    go func() {
        // Send operation: send "Hello from a goroutine" into the channel
        messages <- "Hello from a goroutine"
    }()

    // Receive operation: receive a value from the channel and store it in msg
    // This blocks until a value is sent
    msg := <-messages
    fmt.Println(msg)
}
```

Expected output:

```
Hello from a goroutine
```

Explanation:

* `make(chan string)` creates a channel that can carry `string` values.
* You send a value into a channel using `<-`: `messages <- "Hello from a goroutine"` (this is the sender)
* You receive a value from a channel using `<-`: `msg := <-messages` (this entire statement is the receiver)
* The direction depends on which side of the operator the channel appears on.
* The goroutine sends a message, and `main` receives it.
* The receive operation blocks until a value is sent, so the program waits until the value is received.

## Channels Are Blocking

By default, channels block until both sides are ready. This means sending blocks until a receiver is ready, and receiving blocks until a sender is ready.

```go
package main

import "fmt"

func main() {
    // Create an unbuffered channel that can carry int values
    ch := make(chan int)

    // Start a goroutine that sends a value
    go func() {
        // Send operation: send 42 into the channel
        // This blocks until someone is ready to receive
        ch <- 42
        fmt.Println("Value sent")
    }()

    // Receive operation: receive a value from the channel and store it in value
    // This blocks until someone sends a value
    // The entire "value := <-ch" is the receiver
    value := <-ch
    fmt.Println("Value received:", value)
}
```

Expected output:

```
Value received: 42
Value sent
```

Explanation:

* The send operation `ch <- 42` blocks until a receiver is ready.
* The receive operation `value := <-ch` blocks until a sender is ready.
* The entire `value := <-ch` statement is the receiver - it receives a value from the channel and stores it in the variable `value`.
* Both operations block until they can complete, which synchronizes the goroutines.
* This blocking behaviour makes channels useful for synchronisation.
* The ordering here is guaranteed by the channel, not by timing.

## Challenge

Write a program that:

1. Creates a channel of integers.
2. Starts a goroutine that sends a number into the channel.
3. Receives the number in `main`.
4. Prints the received value.

Run it using:

```
go run main.go
```

## Solution

```go
package main

import "fmt"

func main() {
    // Create a channel for integers
    numbers := make(chan int)

    // Start a goroutine that sends a number
    go func() {
        // Send 100 into the channel
        numbers <- 100
    }()

    // Receive the number from the channel
    value := <-numbers
    fmt.Println("Received:", value)
}
```

Expected output:

```
Received: 100
```

If you didn't get this working on the first attempt, that's completely normal. Channels can feel unfamiliar at first, but the send and receive pattern becomes clearer with practice.

The key idea is that channels allow goroutines to communicate safely. The blocking behaviour ensures that data is passed correctly between concurrent parts of your program.

## Summary

You have now learned the basics of channels in Go. Here is what you covered:

* Channels are used to pass data between goroutines.
* Channels are created using `make()`.
* Values are sent and received using the `<-` operator.
* Channels block until both sender and receiver are ready.
* This blocking behaviour helps with coordination and safety.

This lesson lays the foundation for more advanced topics like buffered channels, channel directions, and selecting between multiple channels.
