# Using Pointers as Function Parameters

**[Watch this lesson on YouTube](https://youtube.com/watch?v=ro5E-DAnVj4)**

## Introduction

In Go, everything you pass to a function is passed by value. This means the function receives a copy of the data, not the original. Most of the time this is fine, but sometimes you want the function to modify the original value. To do that, you need to pass a pointer.

A pointer gives the function access to the original data in memory, so any changes made inside the function will be visible outside the function too. Think of it as giving the function the address of the data, not a copy of the data.

In this lesson, you will learn the difference between passing a value and passing a pointer, and when each approach makes sense.

## Uses / Use Cases

* Allowing a function to update a variable.

* Modifying fields on a struct.

* Avoiding unnecessary copying of large values.

* Making function behaviour clear and predictable.

## Example: Passing a Value (Copy)

Here is a function that tries to change a number, but the original number does not change.

```go
package main

import "fmt"

func addTen(x int) {
    x = x + 10
}

func main() {
    n := 5
    addTen(n)
    fmt.Println("After addTen:", n)
}
```

Output:

```
After addTen: 5
```

Why?

* `n` is passed by value.

* `addTen` receives a copy.

* Modifying the copy does not affect the original.

* This is Go's default behaviour: functions receive copies.

## Example: Passing a Pointer

Now let's pass a pointer to the same kind of function.

```go
package main

import "fmt"

func addTenPtr(x *int) {
    *x = *x + 10
}

func main() {
    n := 5
    addTenPtr(&n)
    fmt.Println("After addTenPtr:", n)
}
```

Output:

```
After addTenPtr: 15
```

This time:

* `addTenPtr` receives a pointer to the original `n`.

* The `&` operator gets the address, and `*` accesses the value at that address.

* `*x` means "the value stored at this pointer".

* Updating `*x` updates `n` directly.

## Example: Updating a Struct

Pointers become even more useful when working with structs.

```go
package main

import "fmt"

type User struct {
    Name  string
    Score int
}

func increaseScoreCopy(u User) {
    u.Score++
}

func increaseScorePtr(u *User) {
    u.Score++
}

func main() {
    userA := User{Name: "Alice", Score: 10}
    userB := User{Name: "Bob", Score: 10}

    increaseScoreCopy(userA)
    increaseScorePtr(&userB)

    fmt.Println("Copy:", userA.Score)
    fmt.Println("Pointer:", userB.Score)
}
```

Expected output:

```
Copy: 10
Pointer: 11
```

Why the difference?

* `increaseScoreCopy` receives a copy of the struct. Changes do not affect the original.

* `increaseScorePtr` receives a pointer. It modifies the original struct in place.

This is the key idea behind pointer parameters. They let functions work with the real data.

## Challenge

Create a struct type called `Book` with these fields:

* `Title` (string)

* `Pages` (int)

Write a function called `addPages` that takes a pointer to a `Book` as its first parameter and an integer amount as its second parameter. The function should add the amount to the book's page count.

In `main()`, create a book, call `addPages`, and print the updated page count to confirm the original struct changed.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

type Book struct {
    Title string
    Pages int
}

func addPages(b *Book, amount int) {
    b.Pages += amount
}

func main() {
    novel := Book{Title: "The Explorer", Pages: 200}
    addPages(&novel, 50)
    fmt.Println("Updated pages:", novel.Pages)
}
```

Expected output:

```
Updated pages: 250
```

## Summary

You have now learned how passing pointers to functions lets those functions modify the original values. Here is what you covered:

* Passing values sends a copy, so the original cannot be changed.

* Passing a pointer gives the function access to the original data.

* Structs behave the same way. A copy does not change the original, but a pointer does.

* Use pointers when you want a function to update data or avoid unnecessary copying.

* Use values when you want safety and simplicity, and no modification is needed.

This understanding prepares you perfectly for the next lesson, where you will apply the same idea to methods through pointer receivers.

