# JSON Responses in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=BxXfcm6KA68)**

## Introduction

So far, your handlers have returned plain text responses.

Real APIs usually return JSON.

In this lesson, you will learn how to encode data as JSON inside a handler, set the correct response headers, and handle potential encoding errors.

## Uses / Use Cases

* Returning structured API responses.
* Sending data to frontend applications.
* Building simple JSON APIs.
* Making HTTP responses machine-readable.

## Defining the Data

JSON responses usually start with a struct.

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

This struct defines the shape of the JSON response. The tags control the JSON key names; see the json-tags lesson if you have not covered these yet.

## Encoding JSON in a Handler

Here is a complete example.

```go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        user := User{
            ID:    1,
            Name:  "Alice",
            Email: "alice@example.com",
        }

        w.Header().Set("Content-Type", "application/json")

        err := json.NewEncoder(w).Encode(user)
        if err != nil {
            http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
```

## Setting the Content-Type Header

Before sending JSON, you should set the `Content-Type` header:

```
application/json
```

This tells the client how to interpret the response body.

You set headers using:

```go
w.Header().Set("Content-Type", "application/json")
```

Headers must be set before writing the response body.

## Encoding Directly to the Response

Instead of calling `json.Marshal` and then writing the bytes manually, this lesson uses:

```go
json.NewEncoder(w).Encode(user)
```

This writes the JSON directly to the response.

It also handles formatting and writes a newline at the end.

## Handling Encoding Errors

Encoding can fail in rare cases, for example if the data contains unsupported types.

You should always check the error returned by `Encode`.

If encoding fails, return a 500 status code. The helper `http.Error` does both: it sets the status code and writes the message as the response body. You pass the response writer, the message string, and the status code. It is a convenient way to send an error response without calling `WriteHeader` and writing the body separately.

## Running and Testing the Server

Run the program:

```
go run main.go
```

Visit:

```
http://localhost:8080/user
```

You should see:

```
{"id":1,"name":"Alice","email":"alice@example.com"}
```

Your browser may also show that the response type is JSON.

## Challenge

Write a program that:

1. Defines a `Product` struct.
2. Registers a `/product` route.
3. Sets the `Content-Type` header to `application/json`.
4. Encodes the product using `json.NewEncoder(w).Encode`.
5. Returns a 500 status if encoding fails.

Verify the JSON output in your browser.

## Solution

Great job if you attempted this challenge! Here is one possible solution:

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
        product := Product{
            ID:    10,
            Name:  "Notebook",
            Price: 12.99,
        }

        w.Header().Set("Content-Type", "application/json")

        err := json.NewEncoder(w).Encode(product)
        if err != nil {
            http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
```

## Summary

You have now returned JSON from a Go HTTP handler.

* Structs define the shape of your JSON.
* `encoding/json` converts structs into JSON.
* `json.NewEncoder(w).Encode` writes JSON directly to the response.
* The `Content-Type` header should be set to `application/json`.
* Encoding errors should return a 500 status.

This is the foundation for building JSON-based APIs in Go.
