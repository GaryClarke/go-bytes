# If / Else and Else If in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In Go, `if` statements let you run code only when a condition is true. But what if you want to run *different* code when the condition is false?

That's where `else` and `else if` come in.

* `else` runs if the `if` condition is false.
* `else if` lets you check another condition before falling back to `else`.

These are essential tools for controlling the flow of your program.

## Uses / Use Cases

* Run different blocks of code based on different conditions
* Build branching logic for user input, program state, or values
* Handle multiple decision points cleanly and clearly

## Example

Here's a basic example using `if`, `else if`, and `else`:

```go
package main

import "fmt"

func main() {
    age := 17

    if age >= 18 {
        fmt.Println("You're an adult.")
    } else if age >= 13 {
        fmt.Println("You're a teenager.")
    } else {
        fmt.Println("You're a child.")
    }
}
```

Here's how this works:

* If `age` is 18 or more, it prints "You're an adult."
* If not, but `age` is 13 or more, it prints "You're a teenager."
* Otherwise, it prints "You're a child."

Only one block of code runs — the first condition that matches.

## Expected Output

```
You're a teenager.
```

(If you change `age := 10`, it would print "You're a child.")

## Challenge

Write a program that checks a score and prints a message based on this logic:

* If the score is 90 or above, print "Grade: A"
* Else if the score is 80 or above, print "Grade: B"
* Else if the score is 70 or above, print "Grade: C"
* Else, print "Grade: F"

Use an `int` variable named `score`.

Try it yourself, then run:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    score := 82

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }
}
```

You can test it by changing the value of `score`.

If your logic differs slightly, that's perfectly fine. The important part is understanding how the conditions check in order and only one block executes.

## Summary

You learned how to:

* Use `if` to check a condition
* Use `else` to run code when the `if` is false
* Use `else if` to check additional conditions
* Build branching logic to handle multiple outcomes

These are essential tools that you'll use in almost every Go program.

Next up: boolean operators like `&&`, `||`, and `!` — which let you build even more powerful conditions.

