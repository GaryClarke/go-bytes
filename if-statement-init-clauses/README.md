# If Statement Init Clauses in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=f0LQUXwon8M)**

## Introduction

In Go, an `if` statement can do more than just check a condition. It can also include an initialisation statement. This lets you declare and initialise a variable directly inside the `if`, and then use it immediately in the condition.

This pattern is very common in Go, especially when working with errors, function return values, and short lived variables. It helps keep code concise and limits variable scope to where it is actually needed.

In this lesson, you will learn how `if` init clauses work, how their scope behaves, and when they are most useful.

## Uses / Use Cases

* Checking errors returned from functions.
* Validating results from function calls.
* Avoiding temporary variables that live longer than needed.
* Writing clearer, more compact conditional logic.

## Syntax Overview

The general form looks like this:

```go
if initStatement; condition {
    // code
}
```

The init statement runs first.
The condition is evaluated next.
If the condition is true, the block runs.

The variable declared in the init statement exists only inside the `if` and `else` blocks.

## Example: Basic If Init Clause

Here is a simple example that shows a variable being initialised and then used immediately.

```go
package main

import "fmt"

func main() {
    if length := len("hello"); length > 3 {
        fmt.Println("Length:", length)
        fmt.Println("String is long enough")
    }
}
```

Explanation:

* `length := len("hello")` runs first and initialises the variable.
* The condition `length > 3` uses that initialised value.
* The same variable is also available inside the `if` block.
* `length` does not exist outside the `if`.

This makes the lifetime of the variable clear and intentional.

## Example: The Common Error Handling Pattern

This is the most common real world use of an if init clause.

```go
package main

import (
    "errors"
    "fmt"
)

func loadConfig() error {
    return errors.New("missing config file")
}

func main() {
    if err := loadConfig(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Config loaded successfully")
}
```

Explanation:

* `err` is declared inside the `if`.
* If an error is returned, it is handled immediately.
* `err` is not accessible outside the `if`, which avoids accidental reuse.

This pattern appears everywhere in Go code.

## Example: If Init Clause with else

The variable declared in the init clause is also available in the `else` block.

```go
package main

import "fmt"

func main() {
    if n := 10; n%2 == 0 {
        fmt.Println("Even number:", n)
    } else {
        fmt.Println("Odd number:", n)
    }
}
```

Here, `n` is available in both branches, but nowhere else.

## Variable Scope Matters

One important detail is scope. The variable declared in the init clause:

* Exists inside the `if` block.
* Exists inside the `else` block, if present.
* Does not exist after the `if` statement ends.

Trying to use it outside will cause a compile error.

```go
if x := 5; x > 3 {
    fmt.Println(x)
}

fmt.Println(x) // not allowed
```

## Challenge

Write a program that:

1. Calls a function `getUsername()` that returns a string and an error.
2. Uses an if init clause to call the function.
3. Prints an error message if an error is returned.
4. Otherwise, prints the username.

You can simulate the function returning either a name or an error.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import (
    "errors"
    "fmt"
)

func getUsername() (string, error) {
    return "", errors.New("user not found")
}

func main() {
    if name, err := getUsername(); err != nil {
        fmt.Println("Error:", err)
        return
    } else {
        fmt.Println("Username:", name)
    }
}
```

Expected output:

```
Error: user not found
```

If you did not get it exactly right the first time, that is completely normal. If init clauses can be tricky when you first encounter them. The important thing is understanding how the init statement runs first, and how the variable scope is limited to the if and else blocks.

## Summary

You have now learned how to use init clauses with `if` statements in Go. Here is what you covered:

* An `if` statement can include an initialisation statement.
* The init statement runs before the condition is checked.
* Variables declared in the init clause are fully initialised and usable immediately.
* Their scope is limited to the `if` and `else` blocks.
* This pattern is commonly used for error handling and short lived values.

This small feature appears constantly in real Go code and is worth becoming comfortable with early.

