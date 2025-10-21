# The main() Function in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=iH4rCuwIxo4)**

## Introduction

Every Go program starts at a function named `main()`. This is the entry point where execution begins.

Go code is organized into packages, functions, and statements:
- A package groups related code.
- A function is a reusable block of code that does something.
- A statement is a single instruction for the computer to run.

In this lesson, you write a minimal program with `main()` and run it to confirm your setup works.

## Uses / Use Cases

- Understand the minimal structure of a Go program
- Learn where execution starts in Go
- Run a Go file from the terminal
- Build confidence running Go code locally

## Example

Here is a complete Go program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go programmer!")
}
```

Let’s break it down:

- `package main`  
  Declares the package for this file. The special name `main` makes the program executable. Other package names are used for libraries.

- `import "fmt"`  
  Brings in Go’s formatting package so you can print text.

- `func main()`  
  Defines the function where the program starts. Go looks for `main()` and runs the code inside it.

- `fmt.Println("Hello, Go programmer!")`  
  Calls a function that prints text followed by a new line. The text is a string, which means characters inside quotation marks.

To run this file:

1. Save it as `main.go`
2. Open your terminal in the same folder
3. Run:

```
go run main.go
```

## Expected Output

```
Hello, Go programmer!
```

If you see that, your Go setup works and you have written your first Go program.

## Challenge

Write a program that prints two lines:

1. Your name
2. A short message of your choice

Use one `fmt.Println()` call for each line, then run:

```
go run main.go
```

## Solution

Great job if you attempted this challenge. Here is one possible solution:

```go
package main

import "fmt"

func main() {
    fmt.Println("Gary Clarke")
    fmt.Println("I'm learning Go!")
}
```

If your solution looks different, that is fine. The goal is to practice writing and running Go code on your machine.

## Summary

You learned:

- `package main` marks the file as an executable program
- `import "fmt"` lets you print text
- `main()` is the program entry point
- `fmt.Println()` prints strings to the terminal
- You can run a Go file with `go run main.go`

This foundation prepares you for variables, loops, and full applications later.


