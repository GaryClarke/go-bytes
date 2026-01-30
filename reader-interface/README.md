# The Reader Interface in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=wx9lOuF6UrI)**

## Introduction

In Go, many different things can be read from. Files, network connections, HTTP request bodies, and even strings can all provide data. Rather than having separate APIs for each of these, Go uses a common interface called `io.Reader`.

The `io.Reader` interface describes *where data comes from*, not how it is stored. If a type knows how to provide bytes when asked, it can implement this interface and work with a huge amount of Go's standard library.

In this lesson, you learn what `io.Reader` is, how it works, and why it appears everywhere in Go programs.

## Uses / Use Cases

* Reading data from files.
* Reading HTTP request bodies.
* Reading from strings or byte slices.
* Writing reusable input processing code.
* Decoupling logic from the data source.

## The io.Reader Interface

The `io.Reader` interface is defined like this:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Explanation:

* Any type with a `Read` method matching this signature implements `io.Reader`.
* The method reads data into the provided byte slice.
* It returns how many bytes were read and an error.
* When there is no more data, `Read` returns `io.EOF`.

This small interface is one of the foundations of Go's IO system.

## Reading from an io.Reader

Because many types implement `io.Reader`, you can write functions that read data without caring where it comes from.

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

func readAll(r io.Reader) {
    data, err := io.ReadAll(r)
    if err != nil {
        return
    }
    fmt.Println(string(data))
}

func main() {
    reader := strings.NewReader("Hello from io.Reader")
    readAll(reader)
}
```

Expected output:

```
Hello from io.Reader
```

Explanation:

* `strings.NewReader` creates a reader from a string.
* `readAll` works with any `io.Reader`.
* The function does not know or care where the data came from.

## Files Implement io.Reader

Files opened with `os.Open` implement `io.Reader`.

```go
package main

import (
    "io"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return
    }

    println(string(data))
}
```

Explanation:

* The file provides data through the `Read` method.
* `io.ReadAll` reads until `io.EOF`.
* This is a simple and common way to read small files.

## The Power of Interface-Based Design

The real value of `io.Reader` is flexibility.

You can write code that:

* Reads from a file today.
* Reads from a network connection tomorrow.
* Reads from a string in tests.

Your code depends on behaviour, not concrete types.

Many parts of Go work by connecting readers and writers. For example, `io.Copy` takes an `io.Reader` and an `io.Writer` and streams data from one to the other:

```go
io.Copy(dst, src)
```

This pattern enables streaming without loading everything into memory, and it works because both readers and writers are defined by interfaces, not concrete types.

## Challenge

Write a function called `printContents` that:

1. Accepts an `io.Reader`.
2. Reads all data from it.
3. Prints the contents as a string.
4. Call it once with a string reader.
5. Call it once with a file opened using `os.Open`.

Run it using:

```
go run main.go
```

**Note:** Make sure you have a file named `example.txt` in your directory before running the program, or create one with some text content.

## Solution

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

func printContents(r io.Reader) {
    data, err := io.ReadAll(r)
    if err != nil {
        return
    }
    fmt.Println(string(data))
}

func main() {
    textReader := strings.NewReader("Hello from a string reader")
    printContents(textReader)

    file, err := os.Open("example.txt")
    if err != nil {
        return
    }
    defer file.Close()

    printContents(file)
}
```

After running this program:

* The string message is printed to the terminal.
* The contents of `example.txt` are printed to the terminal.

If you didn't get this right the first time, that's completely normal. The key idea is that the function works the same way, regardless of where the data comes from.

## Summary

You have now learned how the `io.Reader` interface works in Go. Here is what you covered:

* `io.Reader` describes the ability to read bytes.
* Any type with a `Read([]byte)` method implements it.
* Files, request bodies, and strings all implement `io.Reader`.
* Functions can depend on interfaces instead of concrete sources.
* `io.Reader` and `io.Writer` work together to enable streaming.

Understanding `io.Reader` completes the core IO model in Go and prepares you for more advanced topics like buffering, streaming, and network programming.
