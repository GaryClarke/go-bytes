# Type Assertions in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=UOrZGoOJTMw)**

## Introduction

Sometimes you have a value and you need to check what type it actually is. For example, you might have a variable of type `any` that could hold a string, an integer, or something else, and you need to find out which one it is.

This is where **type assertions** come in. A type assertion lets you check if a value is a specific type and extract it if it is. You can think of it as asking, "Is this value actually a string?" or "Is this value actually an integer?" and then using it as that type if the answer is yes.

In this lesson, you will learn how type assertions work, how to use them safely, and when they are useful.

## Uses / Use Cases

* Extracting concrete values from interface parameters.
* Handling different underlying types stored in an interface.
* Working with values returned as `any` or other interface types.
* Writing flexible functions that accept interfaces but still need type specific behaviour.

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

Expected output:

```
hello
```

Explanation:

* `value` is an interface value of type `any`.
* `value.(string)` is a type assertion.
* It tells Go, "I expect this value to be a string."
* If the value really is a string, the assertion succeeds and returns it.

This works only if the underlying type matches. If the type does not match, the program will panic at runtime. That is why you should use safe type assertions instead.

## Example: Safe Type Assertions

You can use the "comma ok" pattern to check whether the assertion succeeded before using the value.

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

Expected output:

```
Value is not a string
```

Explanation:

* The assertion returns two values.
* `text` is the extracted value, if successful.
* `ok` is a boolean that tells you whether the assertion worked.
* Check `ok` before using the extracted value to avoid panics.

This is the recommended way to use type assertions when the type is not guaranteed.

## Example: Using Safe Type Assertions with Multiple Types

You can check for multiple types using a chain of if statements.

```go
package main

import "fmt"

func process(value any) {
    if s, ok := value.(string); ok {
        fmt.Println("Got a string:", s)
    } else if i, ok := value.(int); ok {
        fmt.Println("Got an integer:", i)
    } else {
        fmt.Println("Unknown type")
    }
}

func main() {
    process("hello")
    process(42)
    process(true)
}
```

Expected output:

```
Got a string: hello
Got an integer: 42
Unknown type
```

This pattern lets you handle different types stored in an interface value.

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
