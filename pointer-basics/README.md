# Understanding Pointers in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=vEvQptSGYfU)**

## Introduction

In Go, a **pointer** stores the **memory address** of a value. That value might live in a named variable (`age := 30; ptr := &age`) or the compiler can create it for you when you take the address of a composite literal (e.g. `app := &App{DB: db}`). Either way, a pointer lets you reference and modify data indirectly, which is useful when you want to share or update values without copying them.

If you've used pointers in other languages like C or C++, Go's approach will feel familiar, but simpler.

## Uses / Use Cases

* Share data between functions without copying
* Modify values in-place
* Work with `nil` to indicate missing or optional data
* Understand Go’s underlying value vs reference semantics

## Example

Let’s start with a basic example:

```go
package main

import "fmt"

func main() {
    age := 30

    ptr := &age

    fmt.Println("Value of age:", age)
    fmt.Println("Pointer to age:", ptr)
    fmt.Println("Value at pointer:", *ptr)

    *ptr = 35
    fmt.Println("Updated age:", age)
}
```

## Expected Output

```
Value of age: 30
Pointer to age: 0xc0000120a8
Value at pointer: 30
Updated age: 35
```

(The exact pointer address will differ each time you run the program.)

### What’s happening here?

* `age := 30` creates a variable.
* `ptr := &age` assigns a **pointer to `age`**.
* `*ptr` gives us access to the **value stored at that address**.
* When we assign `*ptr = 35`, we are **updating the original variable** through the pointer.

## Explanation

| Symbol | Meaning                                                  |
| ------ | -------------------------------------------------------- |
| `&`    | Address-of operator (gets the memory address)            |
| `*`    | Dereference operator (accesses the value at the address) |

Pointers are always tied to a **specific type**. For example, `*int` is a pointer to an `int`, and `*string` is a pointer to a `string`.

Go also uses `nil` to represent a pointer that does not point to any value. Always check for `nil` before dereferencing pointers you didn’t create yourself.

## Challenge

Create a program that:

1. Declares a `score` variable set to `100`.
2. Creates a pointer to that variable.
3. Updates the value to `150` using the pointer.
4. Prints the updated value.

## Solution

Great job if you attempted this challenge! Here’s one possible solution:

```go
package main

import "fmt"

func main() {
    score := 100

    ptr := &score
    *ptr = 150

    fmt.Println("Updated score:", score)
}
```

If your solution works differently but still updates the value through the pointer, that’s perfectly fine. The goal is to practice reading and writing through pointer references.

## Summary

* A pointer stores the address of a value.
* Use `&` to get the address of an existing value or of a composite literal.
* Use `*` to access or modify the value at that address.
* Changing the value through a pointer updates the original variable.

Pointers help you write more efficient and expressive code, especially when passing large structures or working with optional values.

---
