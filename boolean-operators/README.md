# Boolean Operators in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=aF4mp2GhLcc)**

## Introduction

Now that you know how to use `if` and `else`, let's take things further by combining multiple conditions using **Boolean operators**.

Boolean operators let you write logic like:

* "Check if this *and* that are true"
* "Allow access if this *or* that is true"
* "Run this only if the condition is *not* true"

Go gives us three main Boolean operators:

* `&&` (AND)
* `||` (OR)
* `!` (NOT)

Let's see how they work.

## Uses / Use Cases

* Combine multiple conditions in a single `if` statement
* Invert a condition using `!`
* Create more flexible, expressive logic in your programs
* Support branching logic based on multiple flags or states

## Example

```go
package main

import "fmt"

func main() {
    age := 20
    hasID := true

    if age >= 18 && hasID {
        fmt.Println("Access granted")
    } else {
        fmt.Println("Access denied")
    }
}
```

Here's how this works:

* `age >= 18 && hasID` checks that **both** conditions are true.
* If either is false, access is denied.

Now let's try an OR (`||`):

```go
package main

import "fmt"

func main() {
    isAdmin := false
    hasKey := true

    if isAdmin || hasKey {
        fmt.Println("Access granted")
    } else {
        fmt.Println("Access denied")
    }
}
```

Using the NOT operator (`!`):

```go
package main

import "fmt"

func main() {
    isBanned := false

    if !isBanned {
        fmt.Println("You may enter")
    }
}
```

* `!isBanned` means "if not banned"
* This flips the Boolean value

## Grouping Conditions

You can group conditions with parentheses for clarity:

```go
if (age > 18 && hasID) || isAdmin {
    fmt.Println("Access granted")
}
```

This checks:

* The user is over 18 **and** has ID
  **or**
* The user is an admin

Parentheses help control how Go evaluates the logic.

## Expected Output

Depending on the example you run, you'll see:

```
Access granted
```

or

```
Access denied
```

Try changing values and watch how the logic behaves.

## Challenge

Create a program with three variables:

* `loggedIn` (true or false)
* `isStaff` (true or false)
* `hasPermission` (true or false)

Write a condition that prints `"Action allowed"` only if:

* The user is logged in
* AND they are staff **or** have permission

Then test different combinations of values.

## Solution

Great job if you attempted this challenge! Here's one possible solution:

```go
package main

import "fmt"

func main() {
    loggedIn := true
    isStaff := false
    hasPermission := true

    if loggedIn && (isStaff || hasPermission) {
        fmt.Println("Action allowed")
    } else {
        fmt.Println("Action denied")
    }
}
```

If your logic differs slightly, that's perfectly fine. The important part is understanding how && and || work together and how parentheses control the order of evaluation.

## Summary

You learned:

* `&&` checks if **both** conditions are true
* `||` checks if **at least one** is true
* `!` flips a Boolean value (true → false, false → true)
* Parentheses let you group conditions and control logic

These operators make your `if` statements much more powerful. Boolean operators combine with if statements to create powerful, flexible conditions that control your program's behaviour.

