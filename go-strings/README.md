# Working with Strings in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Strings are one of the most common types in any programming language, and Go treats them as immutable sequences of bytes.

In this lesson, you'll learn:

* How to declare and initialize strings
* How to concatenate strings
* The difference between double quotes and backticks
* How strings behave under the hood

## Uses / Use Cases

* Displaying messages to users
* Storing text data like names or sentences
* Building dynamic output using string concatenation
* Defining multi-line messages or formatted blocks

## Example

```go
package main

import "fmt"

func main() {
    name := "Alice"
    greeting := "Hello, " + name + "!"
    fmt.Println(greeting)

    multiLine := `This is a message
that spans multiple lines
and preserves formatting.`
    fmt.Println(multiLine)
}
```

### Explanation

* `name := "Alice"`
  
  A standard string declared using **double quotes**.

* `"Hello, " + name + "!"`
  
  Strings can be **concatenated** using the `+` operator. Go does not use template literals or string interpolation.

* Backticks (`` ` ``) define **raw string literals**. They preserve every character exactly as written—including newlines or escape sequences—so they are ideal for multi-line messages or content like JSON, SQL, and logs.

## Expected Output

```
Hello, Alice!
This is a message
that spans multiple lines
and preserves formatting.
```

## Challenge

Write a program that:

1. Declares a string variable for your **first name**.
2. Declares a string variable for your **favorite programming language**.
3. Prints a message like: `My name is Alice and I love Go.`  
   (Use string concatenation — not `fmt.Sprintf` yet.)

## Solution

Great job if you attempted this challenge! Here’s one possible solution:

```go
package main

import "fmt"

func main() {
    name := "Alice"
    language := "Go"
    fmt.Println("My name is " + name + " and I love " + language + ".")
}
```

If your solution was slightly different but still produced the right output, that’s perfectly fine. The key is practicing string concatenation and understanding how string variables fit together.

## Summary

In this lesson, you learned:

* Go strings are immutable and defined with double quotes (`"`).
* You can concatenate strings using the `+` operator.
* Backticks (`` ` ``) allow for raw, multi-line string literals.
* Strings in Go are a sequence of bytes, which will matter later when we explore encoding and `rune` types.

---
