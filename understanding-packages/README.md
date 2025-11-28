# Understanding Packages in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Go uses packages to organise code. A package is a collection of related files that belong together and provide a clear boundary for how code is grouped and reused. Every Go file belongs to a package, and every Go program is built from one or more packages.

Packages help you structure larger programs, avoid name collisions, and share code across different parts of your project. They also form the basis of Go's standard library, which is made up of many small, focused packages.

In this lesson, you will learn what a package is, how to import one, how to create your own, and how Go uses package names to find and use code.

## Uses / Use Cases

- Organising related code into separate folders.
- Reusing functions and types across different files.
- Splitting a project into clear and understandable parts.
- Importing functionality from Go's standard library.
- Preparing your code so it can grow as your project becomes larger.

## Example

Here is a simple program that uses the `math` package from Go's standard library:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    result := math.Sqrt(25)
    fmt.Println("Square root:", result)
}
```

Let's break this down:

- **Package declaration**
  Every Go file begins with a package name.
  `package main` tells Go that this file is part of the main package, which creates an executable program.

- **Importing packages**
  The import block lists the packages we want to use.
  `fmt` is used for printing.
  `math` gives us mathematical functions like `Sqrt`.

- **Using package functions**
  Functions inside a package are accessed with dot notation.
  `math.Sqrt(25)` calls the `Sqrt` function from the `math` package.

Packages keep the standard library organised and help you work with clean, focused tools.

## Creating Your Own Package

Let's imagine a folder structure like this:

```
understanding-packages/
    main.go
    helper/
        helper.go
```

Inside `helper/helper.go`:

```go
package helper

func Double(n int) int {
    return n * 2
}
```

Inside `main.go`:

```go
package main

import (
    "fmt"
    "understanding-packages/helper"
)

func main() {
    fmt.Println(helper.Double(10))
}
```

Here is what is happening:

- The `helper` folder contains a Go file with `package helper`.
- Because the folder name and package name match, Go knows how to organise it.
- In `main.go`, we import it using the full path.
- We can now call `helper.Double(10)` from the main package.

This is how you build larger programs in a clean and predictable way.

## Expected Output

```
Square root: 5
```

and for the custom helper example:

```
20
```

## Challenge

Create a folder named `shapes` next to `main.go`.

Inside it, create a file `circle.go` that contains a function called `Area` which takes a radius and returns the area of a circle using this formula:

```
area = 3.14 * radius * radius
```

Then in `main.go`:

1. Import your `shapes` package.
2. Call the `Area` function.
3. Print the result.

Run it using:

```
go run .
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

Folder structure:

```
understanding-packages/
    main.go
    shapes/
        circle.go
```

`shapes/circle.go`:

```go
package shapes

func Area(radius float64) float64 {
    return 3.14 * radius * radius
}
```

`main.go`:

```go
package main

import (
    "fmt"
    "understanding-packages/shapes"
)

func main() {
    fmt.Println(shapes.Area(5))
}
```

## Summary

You have learned how Go uses packages to structure code. Here is what you covered:

- Every Go file belongs to a package.
- The `main` package is used to create executable programs.
- You import packages with the `import` keyword.
- Package functions are accessed with dot syntax.
- You can create your own packages by placing files in their own folders.
- Packages let you organise code logically as your project grows.

Packages are a core part of Go's design and you will use them in every program you write.

