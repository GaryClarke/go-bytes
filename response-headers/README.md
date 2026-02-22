# HTTP Response Headers in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=hLCJNQNthJg)**

## Introduction

So far, you have written response bodies and set status codes.

HTTP responses also include headers. Headers contain metadata about the response.

In this lesson, you will learn how to set response headers, understand what they represent, and why the order of operations matters when writing a response.

## Uses / Use Cases

* Setting the `Content-Type` of a response.
* Controlling caching behaviour.
* Sending metadata to clients.
* Ensuring correct HTTP status handling.

## What Are HTTP Headers?

An HTTP response has three main parts:

1. Status code
2. Headers
3. Body

Headers provide additional information about the response.

For example:

* `Content-Type` tells the client how to interpret the body.
* `Cache-Control` can control caching behaviour.

Headers do not appear in the body. They are sent before the body.

## Setting Headers in Go

You set headers using the `Header()` method on `http.ResponseWriter`.

```go
w.Header().Set("Content-Type", "application/json")
```

This modifies the response headers before they are sent.

Here is a simple example.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintln(w, "Hello with headers")
    })

    http.ListenAndServe(":8080", nil)
}
```

The browser will now interpret the response as plain text.

## Status Code and Header Ordering

The order in which you write a response matters.

Important:

* If you do not call `w.WriteHeader`, Go automatically sends `200 OK`.
* The first call that writes to the response body sends the headers.
* After the headers are sent, they cannot be changed.

For example:

```go
w.WriteHeader(http.StatusCreated)
w.Write([]byte("Created"))
```

This works because the status code is set before writing the body.

But this will not behave as expected:

```go
w.Write([]byte("Created"))
w.WriteHeader(http.StatusCreated)
```

Once the body is written, the status code has already been sent as `200 OK`.

The same applies to headers. They must be set before writing the body.

## Putting It Together

Here is a correct example that sets both a header and a status code.

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)

        msg := Message{Text: "Created successfully"}
        if err := json.NewEncoder(w).Encode(msg); err != nil {
            http.Error(w, "encoding error", http.StatusInternalServerError)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
```

The sequence is:

1. Set headers.
2. Set status code.
3. Write the body.

That order ensures the response is sent correctly. If encoding the JSON fails, we respond with an error instead of sending a partial response.

## Challenge

Write a program that:

1. Starts an HTTP server on port 8080.
2. Registers a `/created` route.
3. Sets the `Content-Type` header to `text/plain`.
4. Sets the status code to 201.
5. Writes a short response body.

Verify that:

* The status code is 201.
* The response body appears correctly.

## Solution

Here is one possible solution.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintln(w, "Resource created")
    })

    http.ListenAndServe(":8080", nil)
}
```

## Summary

You have now learned how HTTP response headers work in Go.

* Headers store metadata about a response.
* Use `w.Header().Set` to modify headers.
* Use `w.WriteHeader` to set the status code.
* If no status is set, `200 OK` is used.
* Headers and status must be set before writing the body.

Understanding this order is essential for writing correct HTTP handlers.
