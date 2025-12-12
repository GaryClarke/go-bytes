# While Loops with for in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=Z9ity16C2ZA)**

## Introduction

You have already learned how to use `for` loops with counters and how to range over slices and maps. In Go, `for` is even more flexible. Go does not have a `while` keyword. Instead, `for` handles all looping patterns, including while-style loops.

A while-style loop runs as long as a condition is true. Go expresses this with a `for` loop that contains only a condition. This keeps the language simple and consistent while still allowing all the patterns you are used to.

In this lesson, you will learn how to write while-style loops, how to use infinite loops, and how `break` and `continue` control loop flow.

## Uses / Use Cases

- Repeating code until something changes.
- Waiting for input to be valid.
- Processing data until a condition is met.
- Creating loops with no fixed number of iterations.
- Building long-running processes that check for exit conditions.

## Example: While-Style Loop with a Condition

Here is a loop that keeps halving a number until it is less than one.

```go
package main

import "fmt"

func main() {
    n := 40

    for n >= 1 {
        fmt.Println("n is:", n)
        n = n / 2
    }
}
```

Let's break this down:

- The condition `n >= 1` controls when the loop runs.
- As long as the condition is true, the loop continues.
- Once the condition becomes false, the loop ends.
- This is the Go equivalent of a while loop in other languages.

### Expected Output

```
n is: 40
n is: 20
n is: 10
n is: 5
n is: 2
n is: 1
```

## Example: Infinite Loop with break

You can create an infinite loop by writing `for { }`. Use `break` to exit when needed.

```go
package main

import "fmt"

func main() {
    count := 0

    for {
        if count == 3 {
            break
        }
        fmt.Println("Count:", count)
        count++
    }
}
```

Explanation:

- `for { }` means the loop has no condition and will run forever unless you break out.
- The loop stops when `count` reaches 3.

### Expected Output

```
Count: 0
Count: 1
Count: 2
```

## Example: Using continue to Skip Iterations

The `continue` statement skips the rest of the current iteration and moves to the next one.

```go
package main

import "fmt"

func main() {
    i := 1

    for i <= 6 {
        if i%2 != 0 {
            i++
            continue
        }
        fmt.Println("Even number:", i)
        i++
    }
}
```

Explanation:

- If `i` is odd, `continue` skips the print.
- Only even numbers reach the `fmt.Println`.
- Note: `continue` works in all loop types, not just while-style loops.

### Expected Output

```
Even number: 2
Even number: 4
Even number: 6
```

## Challenge

Write a program that starts with a number and keeps doubling it until it exceeds 100.

1. Start with a number (for example, `n := 1`).
2. Use a while-style loop: `for n <= 100 { ... }`
3. Inside the loop, print the current value of `n`.
4. Double the number: `n = n * 2`
5. The loop should stop when `n` exceeds 100.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func main() {
    n := 1

    for n <= 100 {
        fmt.Println("n is:", n)
        n = n * 2
    }
}
```

Expected output:

```
n is: 1
n is: 2
n is: 4
n is: 8
n is: 16
n is: 32
n is: 64
```

## Summary

You have now learned how Go uses `for` to express every loop pattern. Here is what you covered:

- Go does not have a separate `while` keyword.
- A while-style loop uses `for` followed by a condition.
- `for { }` creates an infinite loop that you exit with `break`.
- `break` stops a loop immediately.
- `continue` skips the rest of the current iteration.
- Choose the loop style that matches the behaviour you need.

This gives you all the looping patterns you need to write flexible and expressive Go code.

