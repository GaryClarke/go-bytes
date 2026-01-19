# Runes, Bytes, and Strings in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=OKshtT6g0_k)**

## Introduction

Strings in Go can be a little surprising at first, especially when working with characters outside basic ASCII. This is because strings are not sequences of characters. They are sequences of bytes.

To work correctly with text, especially Unicode text, Go introduces the concept of *runes*. Understanding how strings, bytes, and runes relate to each other is essential for avoiding subtle bugs and writing correct string handling code.

In this lesson, you will learn what bytes and runes are, how strings are stored, and why indexing and ranging over strings behave differently.

## Uses / Use Cases

* Working with Unicode text safely.
* Understanding why string length can be surprising.
* Iterating over characters correctly.
* Avoiding bugs when indexing strings.
* Writing correct text processing code.

## Strings Are Sequences of Bytes

In Go, a string is a read-only slice of bytes.

```go
package main

import "fmt"

func main() {
    s := "hello"

    fmt.Println(len(s))
}
```

Expected output:

```
5
```

Explanation:

* Each character in `"hello"` is one byte.
* `len(s)` returns the number of bytes, not characters.

This works as expected for basic ASCII text.

## Bytes in Go

A byte in Go is an alias for `uint8`. You can access individual bytes in a string using indexing.

```go
package main

import "fmt"

func main() {
    s := "hello"

    b := s[0]
    fmt.Println(b)
}
```

Expected output:

```
104
```

Explanation:

* `s[0]` returns the first byte, not the first character.
* `104` is the ASCII value of `'h'`.

You can convert it to a character like this:

```go
fmt.Println(string(b))
```

## Unicode and Runes

Unicode characters may take more than one byte. Go represents Unicode code points using the `rune` type.

A rune is an alias for `int32` and represents a single Unicode code point.

```go
package main

import "fmt"

func main() {
    s := "café"

    fmt.Println(len(s))
}
```

Expected output:

```
5
```

Explanation:

* The string `"café"` contains 4 characters.
* The `é` character takes two bytes in UTF-8.
* `len(s)` counts bytes, not characters.

This is where confusion often starts.

## Indexing a Unicode String

Indexing a string still returns bytes, even for Unicode text.

```go
package main

import "fmt"

func main() {
    s := "café"

    fmt.Println(s[3])
    fmt.Println(string(s[3]))
}
```

Expected output:

```
195

```

Explanation:

* `s[3]` returns a single byte (195).
* That byte is only part of the `é` character.
* Converting it directly to a string produces an invalid character (shown as ``).

Indexing is rarely what you want when working with Unicode strings.

## Ranging Over a String

The correct way to iterate over characters in a string is using `range`.

```go
package main

import "fmt"

func main() {
    s := "café"

    for i, r := range s {
        fmt.Println(i, r, string(r))
    }
}
```

Expected output:

```
0 99 c
1 97 a
2 102 f
3 233 é
```

Explanation:

* `range` decodes UTF-8 automatically.
* `r` is a rune, representing a full Unicode character.
* `i` is the byte index where the rune starts.

This is the safest and most common way to work with characters in strings.

## Challenge

Write a program that:

1. Stores a string containing at least one Unicode character.
2. Prints the length of the string using `len`.
3. Iterates over the string using `range`.
4. Prints each rune as a character.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import "fmt"

func main() {
    text := "Go ❤️"

    fmt.Println("Byte length:", len(text))

    for _, r := range text {
        fmt.Println(string(r))
    }
}
```

Expected output:

```
Byte length: 9
G
o
 
❤️
```

If you didn't get this right on the first attempt, that's completely normal. Unicode handling takes time to feel intuitive. The important part is understanding when you're working with bytes versus runes.

## Aside: Converting Between Strings, Bytes, and Runes

You can convert a string to a slice of bytes or runes explicitly.

```go
bytes := []byte("café")
runes := []rune("café")

fmt.Println(len(bytes))
fmt.Println(len(runes))
```

Expected output:

```
5
4
```

Explanation:

* `[]byte` shows the raw UTF-8 bytes.
* `[]rune` shows actual Unicode characters.
* This makes the difference between bytes and runes very clear.

## Summary

You have now learned how strings, bytes, and runes relate to each other in Go. Here is what you covered:

* Strings are sequences of bytes.
* `len(string)` returns the number of bytes.
* Indexing a string returns bytes, not characters.
* Runes represent Unicode code points.
* Using `range` over a string yields runes safely.
* Converting to `[]rune` helps when you need character counts.

