# Understanding Packages in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=Slh9sTahWV0)**

## Introduction

Packages are Go's way of organising code. A package is a group of related files that work together to provide useful functionality. Every Go file belongs to a package, and every Go program is built from one or more packages.

Most of the time, you will use Go's standard library packages, which give you ready-made tools for printing text, working with maths, dealing with strings, handling time, and much more.

In this lesson, you will learn what a package is, how to import one, and how to call functions inside it.

## Uses / Use Cases

* Organising Go's built in functionality into clear groups.

* Reusing code that lives in the standard library.

* Keeping your own programs tidy by separating concerns.

* Calling helper functions from the packages you import.

## Example

Here is a simple program that uses two standard library packages, `fmt` and `math`.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    result := math.Sqrt(49)
    fmt.Println("Square root:", result)
}
```

Let's break this down:

* **Package declaration**

  The file begins with `package main`. Every Go file starts by telling Go which package it belongs to.

  The special package named `main` tells Go this file is part of an executable program.

* **Importing packages**

  The `import` block lists the packages we want to use.

  `fmt` provides printing functions.

  `math` provides mathematical tools like `Sqrt`.

* **Calling a function from a package**

  `math.Sqrt(49)` means: call the `Sqrt` function that lives inside the `math` package.

  Go uses dot notation to access functions inside a package.

This is the core pattern for using packages in Go.

## Expected Output

```
Square root: 7
```

## Challenge

Write a program that:

1. Imports the `strings` package from the standard library

2. Uses `strings.ToUpper` to convert a message to uppercase

3. Prints the result using `fmt.Println`

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    msg := "go makes packages easy"
    upper := strings.ToUpper(msg)
    fmt.Println(upper)
}
```

## Summary

You have learned the basics of how Go uses packages. Here is what you covered:

* Every Go file begins with a package name.

* The `main` package is used for programs you can run.

* You import standard library packages with the `import` keyword.

* Package functions are called with dot notation.

* Packages help keep code organised and easy to understand.

In a later lesson, you will learn how to create your own packages and structure your own projects.
