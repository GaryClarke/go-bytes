# Channels in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

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
    messages := make(chan string)

    go func() {
        messages <- "Hello from a goroutine"
    }()

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
* You send a value into a channel using `<-`: `messages <- "Hello from a goroutine"`
* You receive a value from a channel using `<-`: `msg := <-messages`
* The direction depends on which side of the operator the channel appears on.
* The goroutine sends a message, and `main` receives it.
* The program waits until the value is received.

## Channels Are Blocking

By default, channels block until both sides are ready. This means sending blocks until a receiver is ready, and receiving blocks until a sender is ready.

```go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42
        fmt.Println("Value sent")
    }()

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

* Sending blocks until a receiver is ready.
* Receiving blocks until a sender is ready.
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
    numbers := make(chan int)

    go func() {
        numbers <- 100
    }()

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
