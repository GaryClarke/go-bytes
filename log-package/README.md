# Using the log Package

**[Watch this lesson on YouTube](https://youtu.be/ggjSwiLZ_6A)**

## Introduction

As your programs grow, printing messages with `fmt.Println` is no longer enough. You will want a way to record information about what your program is doing, especially when something unexpected happens. Go includes a built in package called `log` that provides simple and reliable logging tools.

The `log` package prints messages with timestamps and supports helpful functions for ordinary messages and fatal errors. It is a small package, but powerful enough for many applications and ideal for beginners.

In this lesson, you will learn how to use the `log` package to write log messages, include timestamps, and understand when to use each logging function.

## Uses / Use Cases

- Recording information about program behaviour.
- Debugging unexpected behaviour in development.
- Logging errors in a clear and timestamped format.
- Replacing `fmt.Println` for more structured output.
- Providing helpful runtime information for future troubleshooting.

## Example: Basic Logging

Here is a basic example using the `log` package:

```go
package main

import (
    "log"
)

func main() {
    log.Println("Starting application")
    log.Println("Loading configuration")
    log.Println("Application running")
}
```

Output will look something like this:

```
2025/01/01 12:00:00 Starting application
2025/01/01 12:00:00 Loading configuration
2025/01/01 12:00:00 Application running
```

Let's break this down:

- `log.Println` works like `fmt.Println`, but includes a timestamp.
- The timestamp format is fixed and easy to read.
- This is the simplest and most common way to log events in Go.

## Example: Logging Errors

Logging becomes especially useful when something goes wrong.

```go
package main

import (
    "errors"
    "log"
)

func loadConfig() error {
    return errors.New("config file not found")
}

func main() {
    err := loadConfig()
    if err != nil {
        log.Println("Error:", err)
        return
    }

    log.Println("Config loaded successfully")
}
```

Explanation:

- When `loadConfig` returns an error, we log it with `log.Println`.
- Logging with a timestamp makes it easier to trace when the issue occurred.

## Example: log.Fatal

The `log` package also provides `log.Fatal` for messages that should stop the program immediately.

```go
package main

import (
    "errors"
    "log"
)

func connect() error {
    return errors.New("database not reachable")
}

func main() {
    err := connect()
    if err != nil {
        log.Fatal("Connection failed:", err)
    }

    log.Println("Connected successfully")
}
```

Output:

```
2025/01/01 12:00:00 Connection failed: database not reachable
```

After printing the message, `log.Fatal` terminates the program.
This makes it useful for unrecoverable errors such as missing configuration or failed connections.

## Example: Customising Flags (Optional but Useful)

You can change what appears in log messages by setting flags.

```go
package main

import (
    "log"
)

func main() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.Println("Customised logging")
}
```

This adds the date, time, and file/line number to each message.
Although optional, it is a nice demonstration of how customisable the `log` package is.

## Expected Output

Depending on your system, you may see something like:

```
2025/01/01 12:00:00 Customised logging main.go:9
```

The exact timestamp and file path will vary.

## Challenge

Write a program that:

1. Logs a startup message using `log.Println`.
2. Calls a function called `loadUser` that always returns an error.
3. If the error is not nil, use `log.Fatal` to log the error and exit.
4. Otherwise, log a success message.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import (
    "errors"
    "log"
)

func loadUser() error {
    return errors.New("user data missing")
}

func main() {
    log.Println("Starting service")

    err := loadUser()
    if err != nil {
        log.Fatal("Critical failure, shutting down:", err)
    }

    log.Println("Service started successfully")
}
```

## Summary

You have now learned how to use Go's built in `log` package. Here is what you covered:

- `log.Println` prints messages with timestamps.
- Logging is better than printing for information you want to keep track of.
- `log.Fatal` logs a message and exits the program immediately.
- You can customise log behaviour with `log.SetFlags`.
- Logging helps you debug and monitor your programs effectively.

The `log` package is small, simple, and reliable, and it prepares you for more advanced logging tools you may use later.

