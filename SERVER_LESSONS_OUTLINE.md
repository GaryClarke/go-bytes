# HTTP Server Lessons Outline

Split the original "servers-and-handlers" content into two focused lessons.

---

## Lesson 1: Starting an HTTP Server

**Folder:** `starting-http-server` (or `http-server-basics`)

**Single concept:** How do I run an HTTP server and respond to a request?

### Content Scope

**Include:**
- Go's `net/http` package provides HTTP server support
- A server listens on a port for incoming requests
- One minimal example: start server, register a single handler for `/`, write a response
- Explain `http.ListenAndServe(addr, nil)` - starts the server, blocks
- Explain the handler function: `func(w http.ResponseWriter, r *http.Request)` - what it receives
- `fmt.Fprintln(w, "Hello")` writes the response body
- How to run and test (visit `http://localhost:8080`)

**Do NOT include:**
- The `Handler` interface
- Multiple routes
- ServeMux (skip the term entirely, or one line: "Go uses a default router; you'll learn more later")
- HTTP method checking
- Status codes

### Example Code

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

### Challenge

Start an HTTP server on port 8080 that returns a single message when you visit `/`. Use a message of your choice.

### Key Takeaway

You can build a working HTTP server in Go with just a few lines. No framework required.

---

## Lesson 2: Handlers and Routes

**Folder:** `http-handlers-and-routes`

**Single concept:** How do I respond to different URL paths?

### Prerequisites

- Lesson 1 (Starting an HTTP Server)

### Content Scope

**Include:**
- Multiple routes with `http.HandleFunc("/path", handler)`
- Each path gets its own handler
- The default ServeMux matches incoming paths to handlers (brief, plain-language explanation)
- Optional: Handler function signature recap (w, r)

**Consider for this lesson or a later one:**
- HTTP method checking (`r.Method`, `http.MethodGet`, `w.WriteHeader`) - could be "Extra Credit" or "Handling HTTP Methods" lesson
- The `Handler` interface - only if you want to explain `HandleFunc` under the hood; otherwise defer

### Example Code

```go
http.HandleFunc("/", ...)
http.HandleFunc("/about", ...)
http.ListenAndServe(":8080", nil)
```

### Challenge

Create a server with handlers for `/` and `/hello`, each returning a different message.

### Key Takeaway

You register one handler per route. Go's default multiplexer sends each request to the matching handler.

---

## Possible Lesson 3 (Future): Handling HTTP Methods

**Folder:** `http-method-handling`

- Check `r.Method`
- Use `http.MethodGet`, `http.MethodPost`, etc.
- Return 405 Method Not Allowed when appropriate
- `w.WriteHeader(http.StatusMethodNotAllowed)`

---

## Summary

| Lesson | Focus | Core question |
|--------|-------|---------------|
| 1. Starting an HTTP Server | One server, one route | How do I run a server and respond? |
| 2. Handlers and Routes | Multiple paths | How do I respond to different URLs? |
| 3. HTTP Method Handling (optional) | GET vs POST etc. | How do I respond based on HTTP method? |
