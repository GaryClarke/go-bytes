# Reading from Files in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

After writing data to files, the next natural step is reading it back. Many programs need to load configuration files, read saved data, or process text from disk.

Go provides the `os` package for opening files and the `io` package for reading their contents. Together, they let you read files safely and predictably.

In this lesson, you will learn how to open a file, read its contents into memory, handle errors correctly, and close the file when you are done.

## Uses / Use Cases

* Reading configuration or data files.
* Loading text written by another program.
* Processing exported reports or logs.
* Verifying file contents during development.
* Building simple file-based workflows.

## Example: Opening a File with os.Open

To read a file, you first need to open it using `os.Open`.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    fmt.Println("File opened:", file.Name())
}
```

Explanation:

* `os.Open` opens an existing file for reading.
* It returns a file and an error.
* If the file does not exist, an error is returned.
* You must always check the error before continuing.

At this point, the file is open but no data has been read yet.

## Example: Reading the Entire File

A simple way to read all the contents of a file is with `io.ReadAll`.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println(string(data))
}
```

Explanation:

* `io.ReadAll` reads everything from the file.
* It returns a slice of bytes.
* Converting it to `string` makes it printable.
* `defer file.Close()` ensures the file is closed properly.

This approach is ideal for small files that fit comfortably in memory.

## Expected Output

If `example.txt` contains:

```
Hello
From a file
```

The program output will be:

```
Hello
From a file
```

## Challenge

Write a program that:

1. Opens a file called `notes.txt`.
2. Reads its entire contents.
3. Prints the contents to the terminal.
4. Uses `defer` to close the file.

You can reuse a file created in a previous lesson or create one manually.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("notes.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println(string(content))
}
```

Expected output:

```
Lesson notes
Writing files in Go
Using os.Create
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!  
The important thing is understanding how to open files, read their contents, and close them properly.

## Summary

You have now learned how to read from files in Go. Here is what you covered:

* `os.Open` opens a file for reading.
* Always check the error returned when opening files.
* `io.ReadAll` reads the entire file into memory.
* File contents are returned as bytes and can be converted to strings.
* Files should always be closed using `defer`.

This lesson completes the basics of file input and output. From here, you can explore reading files line by line, appending data, or working with larger files efficiently.


