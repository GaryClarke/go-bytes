# Switch Statements in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=bkSS-VllkOk)**

## Introduction

In Go, `switch` statements offer a clear and simple way to handle multiple conditions.

Instead of writing several `if` and `else if` blocks, you can use a `switch` to compare a value against several options.

Go's `switch` statements are easy to read and don't require `break` statements like in some other languages.

In this lesson, you'll learn how `switch` works and how to use it to simplify conditional logic.

## Uses / Use Cases

* Choosing different actions based on a value
* Replacing long chains of `if` and `else if`
* Writing cleaner and more readable decision logic
* Creating simple menus or response handling based on input

## Example

Here's a Go program that uses a `switch` statement:

```go
package main

import "fmt"

func main() {
    color := "green"

    switch color {
    case "red":
        fmt.Println("Stop")
    case "yellow":
        fmt.Println("Slow down")
    case "green":
        fmt.Println("Go")
    default:
        fmt.Println("Unknown color")
    }
}
```

Here's how this works:

* `switch color`
  Tells Go to compare the value of `color` to each `case`.

* `case "red":`
  If `color` is `"red"`, this block runs.

* `default:`
  This block runs if none of the cases match.

Go automatically stops after the first match.

You don't need a `break` statement like in other languages.

### Multiple values in one case

You can match several values in a single case using commas:

```go
day := "Saturday"
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

This checks if `day` is either `"Saturday"` or `"Sunday"`.

## Expected Output

```
Go
```

For the first example where `color` is `"green"`.

## Challenge

Create a program that checks the value of a `status` variable and prints a message:

* If `status` is `"success"`, print `"Operation completed successfully."`
* If `status` is `"error"`, print `"Something went wrong."`
* If `status` is `"loading"`, print `"Please wait..."`
* For anything else, print `"Unknown status."`

Use a `switch` statement to handle the logic.

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    status := "success"

    switch status {
    case "success":
        fmt.Println("Operation completed successfully.")
    case "error":
        fmt.Println("Something went wrong.")
    case "loading":
        fmt.Println("Please wait...")
    default:
        fmt.Println("Unknown status.")
    }
}
```

If your solution works differently but gets the same result, that's perfectly fine. The important part is understanding how switch compares values to cases and executes the matching block.

## Summary

In this lesson, you learned:

* `switch` is a cleaner alternative to multiple `if` or `else if` blocks
* Go compares the value to each `case` in order
* `default` runs when there's no match
* No `break` is needed in Go — it stops automatically
* You can check multiple values in one `case` using commas

Now you're ready to write cleaner decision logic and build smarter programs.

