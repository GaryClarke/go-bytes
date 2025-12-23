# Type Assertions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Interfaces in Go describe behaviour, not concrete types. When a value is stored in an interface, Go remembers the concrete type inside it, but you can only access it through the interface's methods.

Sometimes, you need to get the concrete value back out of an interface. This is where **type assertions** come in. A type assertion lets you ask, "Is this interface value actually a specific concrete type?" If it is, you can extract and use it.

In this lesson, you will learn how type assertions work, how to use them safely, and when they make sense.

## Uses / Use Cases

* Extracting concrete values from interface parameters.
* Handling different underlying types stored in an interface.
* Writing flexible functions that accept interfaces but still need type specific behaviour.
* Working with values returned as `any` or other interface types.

## Example: A Simple Type Assertion

Here is a basic example using `any`, which can hold a value of any type.

```go
package main

import "fmt"

func main() {
    var value any = "hello"

    text := value.(string)

    fmt.Println(text)
}
```

Explanation:

* `value` is an interface value of type `any`.
* `value.(string)` is a type assertion.
* It tells Go, "I expect this value to be a string."
* If the value really is a string, the assertion succeeds and returns it.

This works only if the underlying type matches.

## Example: What Happens When the Type Is Wrong

If the underlying type does not match, a type assertion without protection will cause a panic.

```go
package main

import "fmt"

func main() {
    var value any = 42

    text := value.(string)

    fmt.Println(text)
}
```

This program will crash at runtime, because the value is not a string.

To avoid this, Go provides a safe form of type assertion.

## Example: Safe Type Assertions

You can use the "comma ok" pattern to check whether the assertion succeeded.

```go
package main

import "fmt"

func main() {
    var value any = 42

    text, ok := value.(string)

    if !ok {
        fmt.Println("Value is not a string")
        return
    }

    fmt.Println("String value:", text)
}
```

Explanation:

* The assertion returns two values.
* `text` is the extracted value, if successful.
* `ok` is a boolean that tells you whether the assertion worked.
* This pattern prevents runtime panics.

This is the recommended way to use type assertions when the type is not guaranteed.

## Example: Type Assertions with Interfaces

Type assertions are especially useful when working with interfaces.

```go
package main

import "fmt"

type Printer interface {
    Print()
}

type ConsolePrinter struct{}

func (ConsolePrinter) Print() {
    fmt.Println("Printing to console")
}

func debug(p Printer) {
    if cp, ok := p.(ConsolePrinter); ok {
        fmt.Println("Concrete type is ConsolePrinter:", cp)
    } else {
        fmt.Println("Unknown printer type")
    }
}

func main() {
    var p Printer = ConsolePrinter{}
    debug(p)
}
```

Explanation:

* `p` is an interface value.
* Inside `debug`, we check whether `p` is actually a `ConsolePrinter`.
* The function still accepts the interface, but can branch based on the concrete type.

This pattern should be used sparingly, but it is sometimes necessary.

## Expected Output

```
Printing to console
Concrete type is ConsolePrinter: {}
```

## When to Use Type Assertions

Type assertions are useful, but they should not be your first tool. Prefer:

* Interfaces to describe behaviour.
* Polymorphism through method calls.

Use type assertions when:

* You genuinely need type specific behaviour.
* You are working with values stored as `any`.
* You are handling data from flexible or external sources.

If you find yourself using many type assertions, it may be a sign that your interfaces need adjusting.

## Challenge

Write a function called `describe` that accepts a parameter of type `any`.

Inside the function:

1. If the value is a string, print `"String:"` followed by the value.
2. If the value is an integer, print `"Int:"` followed by the value.
3. Otherwise, print `"Unknown type"`.

Use safe type assertions to check each case.

Call the function with a string, an integer, and a different type.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func describe(value any) {
    if s, ok := value.(string); ok {
        fmt.Println("String:", s)
        return
    }

    if i, ok := value.(int); ok {
        fmt.Println("Int:", i)
        return
    }

    fmt.Println("Unknown type")
}

func main() {
    describe("hello")
    describe(42)
    describe(true)
}
```

Expected output:

```
String: hello
Int: 42
Unknown type
```

If you did not get it exactly right the first time, that is completely normal. Type assertions can be tricky when you first encounter them. The important thing is understanding how the comma ok pattern works and when to use safe assertions.

## Summary

You have now learned how type assertions work in Go. Here is what you covered:

* Interfaces can hold values of different concrete types.
* A type assertion extracts a concrete value from an interface.
* Unsafe assertions can cause runtime panics.
* Safe assertions use the comma ok pattern.
* Type assertions are useful when working with `any` or flexible interfaces.
* They should be used carefully and only when necessary.

Type assertions give you a way to inspect and work with concrete types while still benefiting from Go's interface based design.

