# Using the `any` Type

**[Watch this lesson on YouTube](https://youtube.com/watch?v=lmadIUfnIvo)**

## Introduction

Go is a statically typed language, which means every variable has a clear and specific type. Most of the time this is exactly what you want, because it helps catch mistakes early and makes code easier to understand.

Sometimes, though, you need a variable or a function that can work with values of any type. Go provides a special built in type for this situation. It is called `any`.

The `any` type is an alias for `interface{}`. It represents a value of any type and lets you write flexible code when strict typing is not required.

In this lesson, you will learn what `any` is, how to use it, and why it should be used carefully.

## Uses / Use Cases

- Storing values of different types in the same slice or map.
- Accepting any type as a function argument.
- Returning different kinds of values from a single function.
- Working with data that is not strongly typed, such as decoded JSON.
- Creating simple, generic helper functions.

## Example

Here is a simple function that accepts a value of any type and prints both the value and its underlying type:

```go
package main

import "fmt"

func describe(value any) {
    fmt.Println("Value:", value)
    fmt.Println("Type:", fmt.Sprintf("%T", value))
}

func main() {
    describe(42)
    describe("hello")
    describe(3.14)
}
```

Let’s break this down:

- **Using `any` as a parameter type**
  The function `describe` accepts `value any`. This means it will accept an integer, a string, a float, or any other Go type.

- **Printing the type**
  `fmt.Sprintf("%T", value)` tells us the dynamic type stored inside the variable.

- **Calling the function with different types**
  Each call works because `any` can hold any value.

This flexibility is useful, but it should be applied only when a more specific type does not make sense.

### Using `any` in a Map

You can also use `any` as a map value type to store different types of values in the same map:

```go
package main

import "fmt"

func main() {
    data := map[string]any{
        "name":  "Alice",
        "age":   30,
        "price": 19.99,
    }

    fmt.Println("Name:", data["name"])
    fmt.Println("Age:", data["age"])
    fmt.Println("Price:", data["price"])
}
```

Here, the map can store strings, integers, and floats all together because the value type is `any`. This is useful when you need to work with data that has mixed types, such as JSON data or configuration settings.

## Expected Output

From the first example:
```
Value: 42
Type: int
Value: hello
Type: string
Value: 3.14
Type: float64
```

From the map example:
```
Name: Alice
Age: 30
Price: 19.99
```

## Challenge

Create a function called `logValue` that takes a label as a string and a second argument of type `any`.

The function should print the label, the value, and the type of the value.

In `main()`, call `logValue` three times with different types of values.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here’s one way to solve it:

```go
package main

import "fmt"

func logValue(label string, v any) {
    fmt.Println(label+":", v)
    fmt.Println("Type:", fmt.Sprintf("%T", v))
}

func main() {
    logValue("Age", 30)
    logValue("Name", "Alice")
    logValue("Price", 19.99)
}
```

## Summary

You have learned how to use the `any` type in Go and where it fits into Go’s type system. Here is what you covered:

- `any` is an alias for `interface{}`, Go’s most general type.
- It can store values of any type.
- Functions that use `any` can accept many different kinds of inputs.
- This flexibility is helpful in situations where strict types do not apply.
- It should be used thoughtfully so that your code stays clear and predictable.

You now have a tool for writing more flexible functions, while still understanding when strong typing is the better choice.
