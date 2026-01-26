# Break and Continue in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=dz4GKV2CSc4)**

## Introduction

When working with loops, you often need more control than simply running code a fixed number of times. Sometimes you want to stop a loop early, or skip part of a loop and move on to the next iteration.

Go provides two simple keywords for this: `break` and `continue`.

In this lesson, you will learn how `break` and `continue` work, when to use them, and how they affect the flow of a loop.

## Uses / Use Cases

* Exiting a loop as soon as a condition is met.
* Skipping unwanted values during iteration.
* Avoiding deeply nested `if` statements.
* Writing clearer and more intentional loop logic.
* Handling input or data until a stopping condition occurs.

## Using break to Exit a Loop

The `break` keyword immediately stops the loop and continues execution after the loop.

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

Expected output:

```
Count: 0
Count: 1
Count: 2
```

Explanation:

* This is an infinite loop using `for {}`.
* The loop runs until a condition is met.
* `break` exits the loop when `count` reaches 3.
* This pattern is common when you do not know in advance how many iterations are needed.

## break in While-Style Loops

`break` works the same way in loops that use only a condition.

```go
package main

import "fmt"

func main() {
    count := 0

    for count < 10 {
        if count == 4 {
            break
        }
        fmt.Println(count)
        count++
    }
}
```

Expected output:

```
0
1
2
3
```

Explanation:

* The loop would normally run while `count < 10`.
* `break` stops it early when `count` reaches 4.

## Using continue to Skip an Iteration

The `continue` keyword skips the rest of the current loop iteration and moves on to the next one.

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }
}
```

Expected output:

```
1
2
4
5
```

Explanation:

* When `i` is 3, `continue` is executed.
* The print statement is skipped for that iteration.
* The loop continues with the next value.

`continue` does not stop the loop, it just skips ahead.

## continue in While-Style Loops

`continue` is also useful in condition-based loops.

```go
package main

import "fmt"

func main() {
    i := 0

    for i < 5 {
        i++

        if i == 2 {
            continue
        }

        fmt.Println(i)
    }
}
```

Expected output:

```
1
3
4
5
```

Explanation:

* The loop increments `i` first.
* When `i` equals 2, the print statement is skipped.
* The loop then continues normally.

Be careful to update loop variables before using `continue`, otherwise you may create an infinite loop.

## Choosing Between break and continue

Use `break` when:

* You want to stop looping completely.
* A result has been found.
* Continuing would be unnecessary or incorrect.

Use `continue` when:

* You want to skip certain values.
* Only part of the loop body should be skipped.
* The loop should keep running.

Both keywords help reduce nested `if` statements and make intent clearer.

## Challenge

Write a program that:

1. Loops from 1 to 10.
2. Skips numbers that are divisible by 3.
3. Prints each number you don't skip, and stops when the number reaches 8.

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
    for i := 1; i <= 10; i++ {
        if i%3 == 0 {
            continue
        }

        if i == 8 {
            break
        }

        fmt.Println(i)
    }
}
```

Expected output:

```
1
2
4
5
7
```

If you didn't get this exactly right the first time, that's completely normal. The key is understanding how `break` and `continue` change the flow of the loop.

## Summary

You have now learned how `break` and `continue` work in Go. Here is what you covered:

* `break` exits a loop immediately.
* `continue` skips the rest of the current iteration.
* Both work in all forms of `for` loops.
* They help control flow and reduce nesting.
* Used well, they make loops easier to read and reason about.

Understanding these two keywords will help you write clearer and more precise looping logic in Go programs.
