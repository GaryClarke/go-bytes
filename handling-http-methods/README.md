# Handling HTTP Methods in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=8SMB5aD89Y8)**

## Introduction

So far, your handlers have responded to requests based only on the URL path.

However, HTTP requests also include a method. The most common methods are `GET` and `POST`.

In this lesson, you will learn how to inspect the request method using `r.Method`, branch your logic based on it, and return appropriate HTTP status codes.

## Uses / Use Cases

* Responding differently to `GET` and `POST` requests.
* Building basic form handlers.
* Returning proper status codes when a method is not allowed.
* Laying the foundation for simple APIs.

## Understanding HTTP Methods

Every HTTP request includes a method. You can access it from the request:

```go
r.Method
```

This is a string such as:

* `GET`
* `POST`
* `PUT`
* `DELETE`

In many simple servers:

* `GET` is used to retrieve data.
* `POST` is used to send data to the server.

## Branching on the Request Method

You can use a simple `if` statement to branch based on the method.

Here is an example.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            fmt.Fprintln(w, "GET request received")
            return
        }

        if r.Method == http.MethodPost {
            fmt.Fprintln(w, "POST request received")
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    http.ListenAndServe(":8080", nil)
}
```

Explanation:

* `r.Method` contains the request method.
* `http.MethodGet` and `http.MethodPost` are constants provided by Go.
* If the method is neither `GET` nor `POST`, a 405 status is returned.

In small examples, branching inside a single handler is completely normal. It keeps everything in one place and makes the flow easy to see.

In larger applications, it is common to separate behaviour more clearly, either by using different handlers or a routing library. For now, focusing on `r.Method` helps you understand how HTTP requests are handled at a basic level.

## Returning Status Codes

So far, you have relied on the default status code, which is `200 OK`.

You can explicitly set a status code using:

```go
w.WriteHeader(http.StatusCreated)
```

For example:

```go
w.WriteHeader(http.StatusCreated)
fmt.Fprintln(w, "Resource created")
```

Important:

* If you do not call `WriteHeader`, Go automatically sends `200 OK`.
* You must call `WriteHeader` before writing the response body.
* Once the body is written, the status code cannot be changed.

The helper function `http.Error` sets both a status code and a response body.

## Testing Different Methods

Run the program:

```
go run main.go
```

Open your browser and visit:

```
http://localhost:8080
```

You will see the `GET` response.

To test a `POST` request, you can use curl:

```
curl -X POST http://localhost:8080
```

You should see:

```
POST request received
```

If you try another method, such as `PUT`, you should receive a 405 response.

## Challenge

Write a program that:

1. Starts an HTTP server on port 8080.
2. Registers a handler for `/submit`.
3. Uses a `switch` statement on `r.Method`.

If you have learned switch statements, use one here. It fits this pattern well.

4. Returns:

   * A message for `GET`.
   * A different message for `POST`.
   * A 405 status code for all other methods.

Test your server using both a browser and curl.

## Solution

Great job if you attempted this challenge! Here is one possible solution:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            fmt.Fprintln(w, "Submit page")
        case http.MethodPost:
            w.WriteHeader(http.StatusCreated)
            fmt.Fprintln(w, "Submission received")
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.ListenAndServe(":8080", nil)
}
```

The key idea is that the same route can behave differently depending on the HTTP method.

## Summary

You have now handled HTTP methods in a Go server.

* `r.Method` contains the HTTP method.
* You can branch logic using `if` or `switch`.
* Go provides constants like `http.MethodGet`.
* `w.WriteHeader` sets the response status code.
* If no status is set, `200 OK` is used by default.
* `http.Error` sends both a status code and a message.

This allows your server to respond more precisely and behave more like a real web application or API.
