# Grouping Related Data with Go Structs

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

In Go, a struct (short for structure) is a way to group related values together under a single type.

If you're working with multiple variables that describe one thing, like a user's name and age, you can combine them into a struct. This makes your code easier to read, organize, and maintain.

In this lesson, you'll learn how to declare a struct, create a value of that struct type, and access individual fields.

## Uses / Use Cases

* Grouping related data into a single value
* Creating custom types to represent real-world objects
* Keeping your code organized and readable

## Example

Let's say you want to store information about a person, their name and age.

```go
package main

import "fmt"

type Person struct {
    name string
    age  int
}

func main() {
    p := Person{
        name: "Alice",
        age:  35,
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}
```

Let's break it down:

* `type Person struct { ... }`
  This defines a new struct type called `Person` with two fields: `name` (a `string`) and `age` (an `int`).

* `p := Person{ name: "Alice", age: 35 }`
  This creates a new `Person` value by setting each field.

* `p.name` and `p.age`
  These access the individual fields of the struct.

To run the program:

1. Save the code to `main.go`
2. Open your terminal in that folder
3. Run:

```
go run main.go
```

## Expected Output

```
Name: Alice
Age: 35
```

## Challenge

Create a struct called `Book` with the following fields:

* `title` (string)
* `author` (string)
* `year` (int)

Then create a book value, set the fields to anything you like, and print them out using `fmt.Println()`.

## Solution

Here's one possible solution:

```go
package main

import "fmt"

type Book struct {
    title  string
    author string
    year   int
}

func main() {
    b := Book{
        title:  "Go in Practice",
        author: "Alice Smith",
        year:   2025,
    }

    fmt.Println("Title:", b.title)
    fmt.Println("Author:", b.author)
    fmt.Println("Year:", b.year)
}
```

## Summary

You've learned:

* How to define a custom struct using the `type` keyword
* How to group related data into a single value
* How to assign values to a struct using field names
* How to access struct fields with dot notation

Structs are one of Go's most powerful features for organizing data and building real-world applications. You'll use them everywhere, from web APIs to command-line tools.

