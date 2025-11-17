# Zero Values in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In Go, every variable has a value, even if you don't assign one.

This default is called the **zero value**.

It's Go's way of avoiding *undefined* or *null* values at runtime.

Instead of throwing an error or crashing, uninitialized variables just hold predictable, safe defaults.

## Uses / Use Cases

* Knowing what values your variables hold before initialization.

* Avoiding bugs related to uninitialized memory.

* Understanding what happens when you declare a variable without assigning it.

## Example

```go

package main

import "fmt"

func main() {
    var n int
    fmt.Println(n) // 0

    var s string
    fmt.Println(s) // ""

    var b bool
    fmt.Println(b) // false

    var f float64
    fmt.Println(f) // 0.0
}

```

### Explanation

* **int** defaults to `0`

* **string** defaults to `""` (empty string)

* **bool** defaults to `false`

* **float64** defaults to `0.0`

Go guarantees these values are assigned automatically when you declare variables using `var`.

## More Examples

### Slices, Maps, and Pointers

```go

var nums []int

var settings map[string]string

var ptr *int

fmt.Println(nums == nil)     // true

fmt.Println(settings == nil) // true

fmt.Println(ptr == nil)      // true

```

These types get a **nil** zero value.

### Structs

```go
type Book struct {
    Title string
    Pages int
    IsRead bool
}

var b Book

fmt.Println(b.Title)  // ""
fmt.Println(b.Pages)  // 0
fmt.Println(b.IsRead) // false
```

Struct fields follow the same rule — each field gets the zero value of its type.

## Challenge

Declare variables of the following types without assigning a value:

* `int`

* `string`

* `bool`

* `[]float64`

* A struct of your choice

Print each one to confirm their zero values.

## Solution

Great job if you attempted this challenge! Here’s one possible solution:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    var count int
    var message string
    var active bool
    var prices []float64
    var u User

    fmt.Println(count)   // 0
    fmt.Println(message) // ""
    fmt.Println(active)  // false
    fmt.Println(prices)  // nil
    fmt.Println(u)       // {"" 0}
}
```

If your output looks different but still shows the zero values for each type, that’s perfectly fine. The goal is to get comfortable with the defaults Go assigns.

## Summary

* Go assigns **zero values** to all uninitialized variables.

* This helps you write safe code without needing to check for "undefined" values.

* Each type has a known zero value:

  * `int` → `0`

  * `string` → `""`

  * `bool` → `false`

  * `float64` → `0.0`

  * Pointers, slices, maps → `nil`

  * Struct fields → zero values of their types

## Coming Up

Next, we’ll explore **nil** — how Go represents “no value” for pointers, slices, maps, and interfaces.

---

