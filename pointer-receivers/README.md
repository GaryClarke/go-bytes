# Understanding Pointer Receivers in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=KcrqXzIbvq0)**

## Introduction

You have already learned how to create methods with value receivers, where the receiver is a copy of the value you call the method on. Go also allows you to use pointer receivers. A pointer receiver gives your method access to the original value, not a copy.

Pointer receivers are essential when you want a method to modify the state of a type, or when copying a value would be too expensive. Understanding when to use value receivers and when to use pointer receivers ensures methods modify state when intended and avoids unnecessary copying.

In this lesson, you will learn what pointer receivers are, when to use them, and how they behave when you call methods on values and pointers.

## Uses / Use Cases

* Allowing methods to modify the receiver's fields.

* Avoiding unnecessary copying of large structs.

* Ensuring consistent method sets for a type.

* Writing types that behave like objects with updatable state.

## Example

Here is a simple type with a method that updates its own value using a pointer receiver.

```go
package main

import "fmt"

type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{Value: 5}
    c.Increment()
    fmt.Println("Counter:", c.Value)
}
```

Let's break this down:

* **Pointer receiver**

  The receiver is written as `(c *Counter)`.

  This means the method receives a pointer to the original `Counter`, not a copy.

* **Modifying the original value**

  `c.Value++` updates the actual struct stored in `main`, because `c` is a pointer.

* **Calling the method**

  Even though the method has a pointer receiver, you can call it on a value.

  Go automatically takes the address for you.

Pointer receivers let methods change the data they belong to.

## Expected Output

```
Counter: 6
```

## Example: Value Receiver vs Pointer Receiver

Let's compare the two kinds of receivers to see how they behave differently.

```go
package main

import "fmt"

type Number struct {
    n int
}

func (x Number) DoubleValue() {
    x.n = x.n * 2
}

func (x *Number) DoublePointer() {
    x.n = x.n * 2
}

func main() {
    a := Number{n: 10}
    a.DoubleValue()
    fmt.Println("After DoubleValue:", a.n)

    b := Number{n: 10}
    b.DoublePointer()
    fmt.Println("After DoublePointer:", b.n)
}
```

Expected output:

```
After DoubleValue: 10
After DoublePointer: 20
```

Why?

* `DoubleValue` receives a **copy**, so the original does not change.

* `DoublePointer` receives a **pointer**, so the method works on the original.

## Challenge

Create a type called `Account` with two fields:

* `Owner` (string)

* `Balance` (float64)

Write two methods:

1. `Deposit(amount float64)`

   * Uses a pointer receiver

   * Adds the amount to the balance

2. `BalanceInfo() string`

   * Must use a value receiver (not a pointer receiver)

   * Returns a string describing the account owner and balance

In `main()`, create an account, call `Deposit` a few times, then print the result of `BalanceInfo()`.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

type Account struct {
    Owner   string
    Balance float64
}

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

func (a Account) BalanceInfo() string {
    return fmt.Sprintf("%s has a balance of %.2f", a.Owner, a.Balance)
}

func main() {
    acc := Account{Owner: "Alice", Balance: 100}
    acc.Deposit(50)
    acc.Deposit(25)
    fmt.Println(acc.BalanceInfo())
}
```

## Summary

You have now learned how pointer receivers work and when to use them. Here is what you covered:

* Pointer receivers allow methods to modify the original value.

* Value receivers operate on a copy and cannot change the original.

* Go automatically handles calling pointer receiver methods on values by taking the address for you.

* Pointer receivers avoid unnecessary copies and help manage mutable state.

* Choosing between value and pointer receivers depends on whether your method needs to modify the value or avoid copying.

Pointer receivers are a core part of writing clean, efficient Go types that behave the way you expect.

---

If you want, we can move straight into the next Go Byte.

