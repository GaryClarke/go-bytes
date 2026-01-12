# Reading User Input in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=sKtLimBXrAo)**

## Introduction

So far, your programs have mostly worked with values defined directly in code. Many real programs need to accept input from the user, such as names, numbers, or commands typed into the terminal.

Go provides simple functions in the `fmt` package for reading input from standard input, usually the terminal. The most common one is `fmt.Scan`.

In this lesson, you will learn how to read user input using `fmt.Scan` and understand how it works.

## Uses / Use Cases

* Reading values entered by a user.
* Creating interactive command line programs.
* Accepting configuration or parameters at runtime.
* Practising basic input and output workflows.

## Example

The `fmt.Scan` function reads space separated values from standard input.

```go
package main

import "fmt"

func main() {
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scan(&name)

    fmt.Println("Hello,", name)
}
```

Explanation:

* `fmt.Scan` reads input until it encounters whitespace.
* You must pass the address of the variable using `&`.
* The value entered by the user is stored in the variable.

## Expected Output

Example interaction:

```
Enter your name: Alice
Hello, Alice
```

If the user types `Alice` and presses Enter, the program prints the greeting with their name.

## Challenge

Write a program that:

1. Asks the user for their first name and age.
2. Reads both values from the terminal using `fmt.Scan`.
3. Prints a message using the entered values.

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
    var name string
    var age int

    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age)

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

If the user enters:

```
Alice 30
```

The output will be:

```
Name: Alice
Age: 30
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!
The important thing is understanding how `fmt.Scan` reads multiple values when you pass multiple variable addresses.

## Summary

You have now learned how to read user input from the terminal in Go. Here is what you covered:

* `fmt.Scan` reads space separated input from the terminal.
* You must pass variable addresses using `&` to store input.
* You can read multiple values by passing multiple variables.
* Reading user input is a key step towards building interactive Go programs and command line tools.


