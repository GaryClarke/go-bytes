# Type Definitions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=kE1Edf6SvMM)**

## Introduction

Go allows you to create new types based on existing ones. This is called a *type definition*. Defining your own types helps make code clearer, safer, and easier to understand.

At first glance, type definitions can look similar to type aliases, but they behave very differently. Understanding the difference between the two is important, especially as your programs grow and you start modelling real-world concepts.

In this lesson, you will learn how to define new types, how they differ from aliases, and when each should be used.

## Uses / Use Cases

* Giving meaningful names to values.
* Preventing accidental misuse of similar values.
* Modelling domain concepts more clearly.
* Improving readability and maintainability.
* Creating types with their own behaviour later on.

## Defining a New Type

A type definition creates a *new, distinct type* based on an existing one.

```go
package main

import "fmt"

type UserID int

func main() {
    var id UserID = 10
    fmt.Println(id)
}
```

Expected output:

```
10
```

Explanation:

* `UserID` is a new type.
* It is based on `int`, but it is not the same as `int`.
* The underlying representation is still an integer.

This distinction is important.

## Type Safety with Definitions

Because `UserID` is a distinct type, it cannot be mixed freely with `int`.

```go
package main

type UserID int

func main() {
    var id UserID = 10
    var number int = 10

    // This will not compile:
    // id = number
}
```

Explanation:

* Even though both are based on `int`, Go treats them as different types.
* This prevents accidental misuse.
* Type definitions add safety, not just names.
* If you try to compile this code, Go will report a type mismatch error.

You can convert explicitly when needed.

```go
id = UserID(number)
```

## Why Type Definitions Are Useful

Type definitions help express intent.

```go
type Email string
type Username string
```

Both are strings, but they represent different concepts. By defining separate types, you reduce the risk of passing the wrong value to a function.

```go
func sendEmail(to Email) {
    // Function implementation would go here
}
```

This function cannot accidentally be called with a `Username`.

## Type Aliases

Go also supports *type aliases*. An alias does not create a new type. It creates another name for an existing type.

```go
package main

type ID = int

func main() {
    var a ID = 5
    var b int = 10

    a = b
}
```

Explanation:

* `ID` is an alias for `int`.
* There is no new type created.
* `ID` and `int` are interchangeable.

Aliases exist mainly for compatibility and refactoring.

## Key Difference: Definition vs Alias

The syntax looks similar, but the behaviour is not.

```go
type Age int    // new type
type Count = int // alias
```

* `Age` is a distinct type.
* `Count` is just another name for `int`.

This single `=` character changes everything.

## When to Use Each

Use a **type definition** when:

* You want stronger type safety.
* You are modelling a domain concept.
* You want to prevent mixing values accidentally.
* You plan to add methods later.

Use a **type alias** when:

* Renaming an existing type.
* Maintaining backward compatibility.
* Gradually refactoring code.

For most application code, type definitions are far more common.

## Challenge

Write a program that:

1. Defines a new type called `ProductID` based on `int`.
2. Defines a function that accepts a `ProductID`.
3. Tries to call the function with a plain `int` (this will not compile) and observe the error message.
4. Fixes the call using an explicit conversion.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

type ProductID int

func printProduct(id ProductID) {
    fmt.Println("Product ID:", id)
}

func main() {
    var rawID int = 42

    // printProduct(rawID) // does not compile

    printProduct(ProductID(rawID))
}
```

Expected output:

```
Product ID: 42
```

If you did not get this exactly right the first time, that is completely normal. The key takeaway is understanding that a type definition creates a *new* type, even if it is based on an existing one.

## Summary

You have now learned how type definitions work in Go. Here is what you covered:

* Type definitions create new, distinct types.
* They are based on existing types but are not interchangeable.
* This improves safety and expresses intent.
* Type aliases create alternate names, not new types.
* The `=` in a type alias makes a crucial difference.

Type definitions are a powerful tool in Go. They help you write clearer code today and make it easier to extend behaviour later on.
