# The Writer Interface in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In Go, many different things can be written to. Files, network connections, HTTP responses, and even the terminal can all receive data. Rather than having separate APIs for each of these, Go uses a common interface called `io.Writer`.

The power of `io.Writer` is that it describes *behaviour*, not implementation. If a type knows how to write bytes somewhere, it can implement this interface and immediately work with a huge range of Go code.

In this lesson, you learn what the `io.Writer` interface is, how it works, and why it is one of the most important interfaces in the Go standard library.

## Uses / Use Cases

* Writing data to files.
* Sending data over the network.
* Writing HTTP responses.
* Writing logs or formatted output.
* Decoupling code from where data is written.

## The io.Writer Interface

The `io.Writer` interface is defined like this:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Explanation:

* Any type that has a `Write` method with this signature implements `io.Writer`.
* The method receives a slice of bytes.
* It returns how many bytes were written and an error.

That's it. There is no keyword or special syntax for implementing interfaces in Go.

## Writing to an io.Writer

Because many types implement `io.Writer`, you can write functions that work with any of them.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func writeMessage(w io.Writer) {
    w.Write([]byte("Hello from io.Writer\n"))
}

func main() {
    writeMessage(os.Stdout)
}
```

Expected output:

```
Hello from io.Writer
```

Explanation:

* `os.Stdout` represents standard output.
* It implements `io.Writer`.
* The same function can write to many different destinations.

## Files Implement io.Writer

Files returned by `os.Create` or `os.OpenFile` also implement `io.Writer`.

```go
package main

import (
    "io"
    "os"
)

func writeMessage(w io.Writer) {
    w.Write([]byte("Written to a file\n"))
}

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        return
    }
    defer file.Close()

    writeMessage(file)
}
```

Explanation:

* `file` implements `io.Writer`.
* The same function writes to a file without knowing it is a file.
* This is interface based design in action.

## The Power of Interface-Based Design

The real value of `io.Writer` is flexibility.

You can write code once and reuse it everywhere:

* Write to the terminal during development.
* Write to a file in production.
* Write to an HTTP response in a web server.
* Write to a buffer in tests.

Your code depends on behaviour, not concrete types.

Many standard library functions accept an `io.Writer`. For example, `fmt.Fprintln` works with files, stdout, and HTTP responses because they all implement `io.Writer`:

```go
fmt.Fprintln(os.Stdout, "Hello")
fmt.Fprintln(file, "Hello")
```

The same pattern appears throughout the standard library, making code flexible and reusable.

## Challenge

Write a function called `writeGreeting` that:

1. Accepts an `io.Writer`.
2. Writes a greeting message.
3. Call it once with `os.Stdout`.
4. Call it once with a file created using `os.Create`.

Run it using:

```
go run main.go
```

## Solution

```go
package main

import (
    "io"
    "os"
)

func writeGreeting(w io.Writer) {
    w.Write([]byte("Hello from the writer interface\n"))
}

func main() {
    writeGreeting(os.Stdout)

    file, err := os.Create("greeting.txt")
    if err != nil {
        return
    }
    defer file.Close()

    writeGreeting(file)
}
```

After running this program:

* The message is printed to the terminal.
* The same message is written to `greeting.txt`.

You can verify the file was created by checking `greeting.txt` in your directory. The file should contain the greeting message.

If you didn't get this right the first time, that's completely normal. The key idea is that the function does not care *where* it is writing.

## Summary

You have now learned how the `io.Writer` interface works in Go. Here is what you covered:

* `io.Writer` describes the ability to write bytes.
* Any type with a `Write([]byte)` method implements it.
* Files, stdout, and HTTP responses all implement `io.Writer`.
* Functions can depend on interfaces instead of concrete types.
* This leads to flexible, reusable code.

Understanding `io.Writer` is a major step towards writing idiomatic Go. It shows how small interfaces enable powerful composition across the standard library.
