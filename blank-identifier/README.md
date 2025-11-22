# The Blank Identifier `_` in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=ymonJiPYgUQ)**

## Introduction

In Go, if you declare a variable and don't use it, the compiler throws an error. This is part of Go's philosophy of explicit, clear code. Every variable should have a purpose.

That's where the **blank identifier `_`** comes in.

The underscore allows you to **ignore values** you don't care about. It tells Go, *"I know this value is here, but I'm not using it on purpose."* This keeps your code clean while satisfying the compiler's requirement that all values be handled explicitly.

## Uses / Use Cases

* Avoiding compile errors for unused return values

* Ignoring loop indexes or values you don't need

* Handling functions that return multiple values when you only need some of them

## Examples

### 1. Ignoring a return value

```go
value, _ := someFunction()
```

If `someFunction()` returns two values but you only need the first one, use `_` to ignore the second.

### 2. Ignoring the index in a `for` loop

```go
for _, value := range []string{"Go", "Python", "Rust"} {
    fmt.Println("Language:", value)
}
```

When you only need the value from a range loop, use `_` to ignore the index. Go's `range` always returns both index and value, so you must handle both, even if you only use one.

## Challenge

1. Create a function that returns two values: a string and an integer.

2. Call the function and:

   * Use the string, ignore the integer.

   * Then, ignore the string and use the integer.

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func getUserInfo() (string, int) {
    return "Alice", 35
}

func main() {
    // Use string, ignore integer
    name, _ := getUserInfo()
    fmt.Println("Name:", name)

    // Ignore string, use integer
    _, age := getUserInfo()
    fmt.Println("Age:", age)
}
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!  
The important thing is understanding how `_` lets you selectively ignore values while still handling the ones you need.

## Summary

The **blank identifier `_`** is Go's way of saying "ignore this value."

It improves code clarity and avoids unused variable errors.

You can use it to:

* Discard unused return values

* Ignore loop indexes or values

* Handle functions that return multiple values when you only need some of them

Go encourages explicit code, and the blank identifier is part of that clarity.

Next time you see `_`, you'll know: *this value exists, but we're choosing to ignore it.*

