# Handlers and Routes in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=uqQbNkK2c5c)**

## Introduction

In the previous lesson, you started an HTTP server and responded to a single request.

Real servers usually need to respond to more than one URL. Different paths often represent different pages or actions, even if they all run on the same server.

In this lesson, you will learn how to register multiple routes in Go and how incoming requests are matched to handler functions.

## Uses / Use Cases

* Serving different pages from the same server.
* Exposing multiple API endpoints.
* Separating behaviour by URL path.
* Building simple routing logic without a framework.

## Registering Multiple Routes

Go allows you to register more than one handler function by calling `http.HandleFunc` multiple times.

Each call associates a path with a function.

Here is a server with two routes.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Home page")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "About page")
    })

    http.ListenAndServe(":8080", nil)
}
```

Explanation:

* Each call to `http.HandleFunc` registers a path and a handler.
* The path determines which function runs for a given request.
* All handlers run on the same server and port.

## How Routes Are Matched

When a request comes in, Go looks at the request path and finds the most appropriate handler.

For example:

* Visiting `/` runs the handler registered for `/`.
* Visiting `/about` runs the handler registered for `/about`.

Paths must match exactly as written. A request to `/about` will not be handled by `/`.

## The Default ServeMux

You may have noticed that `http.ListenAndServe` is still called with `nil`.

```go
http.ListenAndServe(":8080", nil)
```

Passing `nil` tells Go to use the default request multiplexer, often called the default ServeMux.

The default ServeMux is responsible for:

* Storing the registered routes.
* Matching incoming requests to handlers.
* Calling the correct handler function.

When you use `http.HandleFunc`, you are registering routes on this default ServeMux.

You do not need to interact with it directly for basic routing.

## Running and Testing the Server

Run the program:

```
go run main.go
```

Open a browser and visit:

```
http://localhost:8080
```

You should see:

```
Home page
```

Now visit:

```
http://localhost:8080/about
```

You should see:

```
About page
```

Both routes are handled by the same server process.

## Challenge

Write a program that:

1. Starts an HTTP server on port 8080.
2. Registers three routes: `/`, `/contact`, and `/help`.
3. Returns a different message for each route.

Run the server and visit each path in your browser to confirm the correct response is shown.

## Solution

Nice work if you gave this a go. Here is one possible solution:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Home")
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Contact us")
    })

    http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Help page")
    })

    http.ListenAndServe(":8080", nil)
}
```

If your messages differ, that's fine. The important part is that each path triggers a different handler.

## Summary

You have now handled multiple routes in a Go HTTP server.

* Multiple handlers can be registered using `http.HandleFunc`.
* Each handler is matched to a specific URL path.
* The default ServeMux stores routes and dispatches requests.
* Passing `nil` to `http.ListenAndServe` uses the default ServeMux.
* One server can respond to many different paths.

This builds on the previous lesson and prepares you for handling more varied HTTP behaviour.
