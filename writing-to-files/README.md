# Writing to Files in Go

**[Watch this lesson on YouTube](https://youtu.be/WTVerBW1F34)**

## Introduction

So far, most of your programs have printed output to the terminal. In real programs, you often want to write data to a file instead. This could be logs, reports, exported data, or anything else that needs to be saved.

Go provides the `os` package to work with files and the file system. One of the simplest and most common ways to create and write to a file is using `os.Create`. This function creates a file and gives you a value you can write to.

In this lesson, you will learn how to create a file, write text to it, handle errors properly, and make sure the file is closed correctly.

## Uses / Use Cases

* Saving program output to a file.
* Writing logs or reports.
* Exporting data for later use.
* Creating simple text files from Go programs.
* Learning the basics of file handling.

## Example: Creating a File with os.Create

The `os.Create` function creates a file or truncates it if it already exists.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    fmt.Println("File created:", file.Name())
}
```

Explanation:

* `os.Create` creates a file in the current directory.
* If the file already exists, it is emptied.
* It returns a file and an error.
* Always check the error before continuing.

At this point, the file exists but contains no data.

## Example: Writing Text to a File

Once you have a file, you can write to it using `WriteString`.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    file.WriteString("Hello, file!\n")
    file.WriteString("This text was written from Go.\n")

    fmt.Println("Data written to file")
}
```

Explanation:

* `WriteString` writes text to the file.
* Writing happens in order, just like printing lines.
* Newlines must be added manually.

At this stage, the file contains the written text.

## Example: Closing the File with defer

Files should always be closed when you are finished with them. This ensures that data is written properly and system resources are released.

The idiomatic way to do this in Go is with `defer`.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    file.WriteString("First line\n")
    file.WriteString("Second line\n")

    fmt.Println("File written and closed")
}
```

Explanation:

* `defer file.Close()` ensures the file is closed when `main` finishes.
* This works even if the function returns early.
* Using `defer` for cleanup is a core Go pattern.

## Expected Output

After running the program, the terminal output might look like:

```
File written and closed
```

The file `example.txt` will contain:

```
First line
Second line
```

## Aside: Writing Files with os.WriteFile

For simple cases where you want to write all data at once, Go also provides a helper function called `os.WriteFile`.

```go
err := os.WriteFile("example.txt", []byte("Hello from Go\n"), 0644)
if err != nil {
    fmt.Println("Error writing file:", err)
}
```

This function:

* Creates or opens the file.
* Writes all the data.
* Closes the file internally.

You do not need to call `Close` or use `defer`.
It is useful for small, one-off writes, but using `os.Create` is better when you want to write incrementally or learn how file handling works.

## Challenge

Write a program that:

1. Creates a file called `notes.txt`.
2. Writes three lines of text to the file.
3. Uses `defer` to close the file.
4. Prints a message when the file is successfully written.

Run it using:

```
go run main.go
```

Then open the file to confirm its contents.

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("notes.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    file.WriteString("Lesson notes\n")
    file.WriteString("Writing files in Go\n")
    file.WriteString("Using os.Create\n")

    fmt.Println("notes.txt written successfully")
}
```

Expected output:

```
notes.txt written successfully
```

The file `notes.txt` should contain:

```
Lesson notes
Writing files in Go
Using os.Create
```

## Summary

You have now learned how to write to files in Go. Here is what you covered:

* `os.Create` creates or truncates a file.
* It returns a file and an error that must be checked.
* You can write text using `WriteString`.
* Files should always be closed using `defer`.
* `os.WriteFile` is a convenience function for simple cases.
* Writing to files is a core skill for real programs.

This lesson lays the foundation for reading files, appending data, and working with more advanced file operations later on.


