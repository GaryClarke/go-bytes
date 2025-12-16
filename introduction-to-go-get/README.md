# Introduction to go get

**[Watch this lesson on YouTube](https://youtube.com/watch?v=wHm5AgdBYMY)**

## Introduction

When you build Go programs, you will often want to use code written by other developers. This might be a library that generates unique IDs, sends HTTP requests, handles colours in the terminal, or anything else you might need.

Go manages external dependencies through modules, which you created using `go mod init` earlier in the course. To add a new dependency to your project, you use a command called `go get`.

In this lesson, you will learn what `go get` does, how it interacts with your module, and how to import and use third party packages in your code.

## Uses / Use Cases

- Adding external libraries to your project.
- Updating a dependency to a newer version.
- Making your code more powerful by using community packages.
- Managing dependencies in a clear and predictable way.

## Example: Installing a Package with go get

To demonstrate `go get`, we will install the `github.com/google/uuid` package. This package provides functions for generating unique identifiers.

Run this command inside your module folder:

```
go get github.com/google/uuid
```

After running this command:

- Your `go.mod` file will be updated with the new dependency.
- Your `go.sum` file will include checksums to verify the package. This file is automatically created or updated by Go, and you do not need to edit it manually.
- The package is now available for you to import and use in your code.

### Using the Package

Here is a simple program that generates a UUID and prints it:

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
}
```

Let's break this down:

- **Importing the package**
  The import path `github.com/google/uuid` matches the module path on GitHub.
  This is the same path you used in the `go get` command.

- **Calling its functions**
  `uuid.New()` creates a new UUID.
  You can print it like any other value.

- **Why it works**
  Because your module now depends on `github.com/google/uuid`, Go knows how to download and manage the code behind the scenes.

## Expected Output

Your output will look similar to:

```
Generated UUID: 3a91d5c0-7047-4d64-91a5-0fca1f235c79
```

The exact value will be different each time.

## Updating Dependencies

If the library publishes a new version, you can update it by running:

```
go get -u github.com/google/uuid
```

This tells Go to fetch the latest version allowed by your module's requirements.

## Challenge

Install the `github.com/google/uuid` package using `go get`, then write a program that:

1. Generates two UUIDs.
2. Prints both of them on separate lines.
3. Runs with:

```
go run main.go
```

This confirms that both the installation and the import are working correctly.

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    first := uuid.New()
    second := uuid.New()

    fmt.Println("First UUID:", first)
    fmt.Println("Second UUID:", second)
}
```

Running this program should print two different UUIDs.

## Summary

You have now learned the basics of adding third party code to your Go projects. Here is what you covered:

- `go get` installs external packages into your module.
- Your `go.mod` and `go.sum` files track the dependencies.
- You can import the installed package and use it in your code.
- `go get -u` updates a dependency to a newer version.
- External packages expand what your programs can do.

This prepares you for working with more complex libraries as your projects grow.

