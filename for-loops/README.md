# Using `for` Loops in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=4X__5f8-VcE)**

## Introduction

Loops let you run a block of code multiple times. In Go, the only loop keyword you need is `for`.

This lesson shows how to write a loop that:

* Repeats a block of code a set number of times
* Iterates over collections like slices

You'll learn the most common `for` loop patterns in Go and how to use them in your programs.

## Uses / Use Cases

* Repeat an action multiple times
* Loop over items in a slice or array
* Perform a task for each element in a collection
* Count up or down using an index

## Example

Here's a Go program that prints numbers from 1 to 5 using a basic `for` loop:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
```

Let's break it down:

* `i := 1` - Start the counter at 1
* `i <= 5` - Keep going while `i` is less than or equal to 5
* `i++` - Add 1 to `i` after each loop

Each time the loop runs, it prints the current value of `i`.

This pattern is great when you know exactly how many times you want to repeat something.

### Looping Over a Slice

You can also use `for` to loop over a slice:

```go
package main

import "fmt"

func main() {
    names := []string{"Alice", "Bob", "Charlie"}

    for _, name := range names {
        fmt.Println(name)
    }
}
```

Here's how this works:

* `range names` goes through each item in the slice
* `_` ignores the index (you can replace it with `i` if you need it)
* `name` is the current value at that position in the slice

This loop runs once for each name in the list.

## Expected Output

From the first example:

```
1
2
3
4
5
```

From the second example:

```
Alice
Bob
Charlie
```

## Challenge

Write a program that prints the following 3 lines using a `for` loop:

```
Line 1
Line 2
Line 3
```

Use `fmt.Println()` and a loop counter.

Bonus: Try starting your loop counter at 1 instead of 0.

## Solution

Here's one possible solution:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        fmt.Println("Line", i)
    }
}
```

Your output should look like:

```
Line 1
Line 2
Line 3
```

If you did it differently, that's okay. There are many valid ways to use a loop. Loops can feel tricky at first, so give yourself credit for trying.

## Summary

In this lesson, you learned:

* Go uses a single `for` keyword for all loops
* You can use `for` with a counter (`for i := 0; i < n; i++`)
* You can use `for ... range` to loop over slices
* `range` gives you both the index and value from a collection

Loops are powerful. Now that you can repeat actions, you're ready to explore even more dynamic Go programs.

