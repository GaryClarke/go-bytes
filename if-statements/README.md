# If Statements in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=kRoYqt5GeK8)**

## Introduction

In Go, you use `if` statements to run code **only when a condition is true**.

This is how programs make decisions, by checking whether something is true and acting accordingly.

In this lesson, you'll learn how to:

* Write an `if` statement
* Use comparison operators
* Run code based on a condition

## Uses / Use Cases

* Run different logic depending on a value
* Check numbers, strings, or booleans
* Skip code if a condition is false
* Build up control flow in programs

## Example

Here's a simple Go program that checks an age and prints a message:

```go
package main

import "fmt"

func main() {
    age := 17

    if age >= 18 {
        fmt.Println("You can vote.")
    }
}
```

Let's break that down:

* `age := 17`
  This declares and initializes a variable called `age`.

* `if age >= 18`
  This condition checks: "Is `age` greater than or equal to 18?"

* `fmt.Println("You can vote.")`
  This only runs **if the condition is true**.

In this case, because `17 >= 18` is false, nothing is printed.

Try changing the value of `age` to `21` and running the program again. You should now see:

```
You can vote.
```

## Comparison Operators in Go

Here are the most common comparison operators you can use in an `if` condition:

* `==` - equal to
* `!=` - not equal to
* `<` - less than
* `<=` - less than or equal to
* `>` - greater than
* `>=` - greater than or equal to

Each one returns a `bool`, either `true` or `false`.

## Challenge

Write a program that:

* Stores a number in a variable (e.g. `score := 42`)
* If the score is **greater than or equal to 50**, print `"Pass"`
* Otherwise, print **nothing**

Try it using only an `if` statement, no `else` yet.

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    score := 42

    if score >= 50 {
        fmt.Println("Pass")
    }
}
```

If you change `score` to `75` and run it, you'll see:

```
Pass
```

Otherwise, nothing is printed.

If your score value was different, that's perfectly fine. The important part is understanding how the condition controls when the code runs.

## Summary

In this lesson you learned:

* How to write an `if` statement in Go
* How to check conditions using comparison operators
* That code inside the block only runs if the condition is true
* You can compare numbers, strings, or booleans

With `if` statements, you can make your programs respond to different conditions. Next, you'll learn how to handle both true and false cases with `if/else` statements.

