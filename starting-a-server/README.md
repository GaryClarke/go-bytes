# Starting an HTTP Server in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=wKzX6qYA7sM)**

## Introduction

Go has built-in support for creating HTTP servers through its standard library. You do not need a framework or any third-party packages to get started.

With just a few lines of code, you can start a server, listen for incoming requests, and send responses back to a browser.

In this lesson, you will learn how to start a basic HTTP server in Go and respond to a single HTTP request.

## Uses / Use Cases

* Building simple web servers.
* Creating APIs and internal services.
* Prototyping HTTP based tools.
* Understanding how Go web frameworks work under the hood.

## The net/http Package

Go's HTTP functionality lives in the `net/http` package. This package provides everything needed to build both HTTP clients and servers.

To start a server, you will:

1. Register a handler function.
2. Start listening on a port.

## Creating Your First HTTP Server

Here is a minimal but complete HTTP server.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, HTTP")
    })

    http.ListenAndServe(":8080", nil)
}
```

Explanation:

* `http.HandleFunc` registers a function to handle requests for `/`.
* The handler function is called whenever a request matches that path.
* `http.ResponseWriter` is used to write the response.
* `http.Request` contains information about the incoming request.
* `http.ListenAndServe(":8080", nil)` starts the server on port 8080.

## Understanding the Handler Function

The handler function has this shape:

```go
func(w http.ResponseWriter, r *http.Request)
```

Explanation:

* `w` is used to write data back to the client.
* `r` gives access to the incoming request.
* Writing to `w` sends data in the HTTP response body.

In this lesson, the handler simply writes a short message.

## Running and Testing the Server

Run the program using:

```
go run main.go
```

You should see no output in the terminal. This is expected.

The server is now running and listening for requests.

Open a browser and visit:

```
http://localhost:8080
```

You should see:

```
Hello, HTTP
```

The program will keep running until you stop it using Ctrl+C.

## Challenge

Write a program that:

1. Starts an HTTP server on port 8080.
2. Registers a handler for `/home`.
3. Returns a message of your choice when the route is visited.

Run the program and visit the route in your browser to verify it works.

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to my server")
    })

    http.ListenAndServe(":8080", nil)
}
```

If you didn't get this working on the first attempt, that's completely normal. The key goal is seeing a response in your browser and knowing that Go is running your server.

## Summary

You have now started a basic HTTP server in Go. Here is what you learned:

* Go provides HTTP server support in the standard library.
* An HTTP server listens on a port for incoming requests.
* Handlers are functions that respond to requests.
* Writing to `http.ResponseWriter` sends data to the client.
* A working HTTP server can be created with just a few lines of code.

This lesson gives you the foundation needed to build more complex servers. Next, you will learn how to handle multiple routes and respond to different URLs.
