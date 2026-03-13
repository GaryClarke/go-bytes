# The `init()` Function in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Go allows packages to run initialization code automatically using a special function called `init()`.

You do not call this function yourself. The Go runtime calls it during program startup.

In this lesson, you will learn when `init()` runs, what it is commonly used for, and why it should be used carefully.

## Uses / Use Cases

* Registering components with a package.
* Preparing package-level state.
* Setting up behaviour during package initialization.

Most Go programs use `init()` sparingly.

## What is `init()`?

`init()` is a function that runs automatically when a package is initialized.

Its signature looks like this:

```go
func init() {
    // initialization code
}
```

Unlike normal functions:

* It does not take parameters.
* It does not return values.
* You never call it directly.

Go executes `init()` during program startup.

## When `init()` Runs

Package initialization follows a predictable order.

1. Package-level variables are initialized.
2. The package's `init()` function runs.
3. The program's `main()` function runs.

If multiple packages are imported, Go initializes them before `main()` executes.

This ensures each package is ready to use before the program starts running.

## Example: Registering a Component

One common use of `init()` is registering behaviour with another package.

Here is a simple example.

```go
package main

import "fmt"

var formats = map[string]string{}

func registerFormat(name string, description string) {
    formats[name] = description
}

func init() {
    registerFormat("json", "JavaScript Object Notation")
}

func main() {
    fmt.Println(formats)
}
```

Output:

```
map[json:JavaScript Object Notation]
```

In this example, the `init()` function registers a format before `main()` runs.

This pattern is commonly used in packages that support pluggable behaviour.

## Real World Examples

Several standard library packages use this pattern.

For example, database drivers often register themselves automatically when imported.

This allows applications to enable drivers simply by importing a package.

## When to Avoid `init()`

Although `init()` can be useful, it should be used carefully.

Because it runs automatically, it can make program behaviour harder to follow.

Many Go developers prefer explicit setup inside `main()` instead.

Good guidelines:

* Use `init()` for simple registration.
* Avoid complex logic.
* Avoid hidden side effects.

Clear, explicit initialization is usually easier to understand.

## Challenge

Create a program that:

1. Defines a map that stores supported file types.
2. Implements a `registerType` function.
3. Uses `init()` to register `"csv"` with a short description.
4. Prints the registered types in `main()`.

Verify that the value is available when the program starts.

## Solution

Great job if you attempted this challenge! Here is one possible solution.

```go
package main

import "fmt"

var fileTypes = map[string]string{}

func registerType(name string, description string) {
    fileTypes[name] = description
}

func init() {
    registerType("csv", "Comma-separated values")
}

func main() {
    fmt.Println(fileTypes)
}
```

Because `init()` runs before `main()`, the map already contains the registered type.

## Summary

You have learned how the `init()` function works in Go.

* `init()` runs automatically during package initialization.
* It executes before `main()`.
* It does not take parameters or return values.
* It is commonly used for registration patterns.
* It should be used sparingly.

In most cases, explicit initialization inside `main()` leads to clearer programs.
