# Writing HTTP Responses in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=7pEKQNECJ4U)**

## Introduction

So far, your handlers have returned text using `fmt.Fprintln`.

That works well for simple examples, but under the hood, HTTP responses are just bytes written back to the client.

In this lesson, you will learn how Go sends HTTP response data using `http.ResponseWriter`, and how to write raw bytes directly to the response.

## Uses / Use Cases

* Sending plain text or structured data in responses.
* Returning data generated at runtime.
* Writing responses without using formatting helpers.
* Understanding how all HTTP responses are ultimately sent.

## The http.ResponseWriter

Every HTTP handler receives a value of type `http.ResponseWriter`.

```go
func(w http.ResponseWriter, r *http.Request)
```

The `ResponseWriter` represents the outgoing HTTP response. Anything written to it is sent back to the client.

In earlier lessons, you used it indirectly with `fmt.Fprintln`. That function simply writes formatted text to the writer.

## Writing Directly to the Response

You can write directly to the response using the `Write` method.

Here is a simple example.

```go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from Write"))
    })

    http.ListenAndServe(":8080", nil)
}
```

Explanation:

* `w.Write` sends data directly to the response.
* It expects a slice of bytes.
* The data is written to the HTTP response body.

## Understanding the Byte Slice Argument

HTTP responses are sent as raw bytes over the network.

Because of this, `Write` accepts a `[]byte` rather than a string. Converting a string to bytes is straightforward.

```go
[]byte("Hello")
```

This conversion creates a byte slice containing the string data, which can then be written to the response.

## Comparing fmt.Fprintln and w.Write

Both of these approaches write data to the response:

```go
fmt.Fprintln(w, "Hello")
```

```go
w.Write([]byte("Hello"))
```

The key difference is that `fmt.Fprintln` formats values before writing, while `w.Write` sends bytes exactly as provided.

For simple responses, either approach is fine. Writing bytes gives you more direct control over what is sent.

## Running and Testing the Server

Run the program:

```
go run main.go
```

Visit:

```
http://localhost:8080
```

You should see:

```
Hello from Write
```

The browser receives the response body exactly as written by the handler.

## Challenge

Write a program that:

1. Starts an HTTP server on port 8080.
2. Registers a handler for `/`.
3. Uses `w.Write` to send a response made up of two separate writes.

For example, write one part of the message, then write another.

## Solution

Great job if you attempted this challenge! Here is one possible solution:

```go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello "))
        w.Write([]byte("from Go"))
    })

    http.ListenAndServe(":8080", nil)
}
```

Both writes are sent to the client as part of the same HTTP response body.

## Summary

You have learned how HTTP responses are written in Go.

* `http.ResponseWriter` represents the outgoing response.
* Writing to the response sends data to the client.
* `w.Write` sends raw bytes.
* Strings can be converted to byte slices using `[]byte(...)`.
* Helper functions like `fmt.Fprintln` write to the same underlying response.
* `http.ResponseWriter` implements the `io.Writer` interface, so any function that accepts `io.Writer` can write to an HTTP response. That is why `fmt.Fprintln` works with it, and why you can reuse the same writing logic for files, network connections, and HTTP responses.

This understanding will be useful as you start sending different kinds of data in HTTP responses.
