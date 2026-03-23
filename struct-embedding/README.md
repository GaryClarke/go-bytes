# Struct Embedding in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VIDEO_ID)**

## Introduction

Structs allow you to group related data together.

Sometimes, you want one struct to reuse the fields or behaviour of another. In Go, this is done through embedding.

In this lesson, you will learn how struct embedding works, what promoted fields are, and how embedding differs from using a named field.

## Uses / Use Cases

* Reusing fields across multiple structs.
* Composing larger types from smaller ones.
* Sharing behaviour without inheritance.
* Simplifying access to nested data.

## What is Struct Embedding?

Struct embedding allows you to include one struct inside another without giving it a field name.

Here is a basic example.

```go
package main

import "fmt"

type Address struct {
    City  string
    State string
}

type User struct {
    Name string
    Address
}

func main() {
    user := User{
        Name: "Alice",
        Address: Address{
            City:  "London",
            State: "UK",
        },
    }

    fmt.Println(user.Name)
    fmt.Println(user.City)
}
```

Notice something important:

```go
user.City
```

Even though `City` belongs to `Address`, it can be accessed directly from `User`.

This is called **field promotion**.

## Promoted Fields

When a struct is embedded, its exported fields become available on the outer struct.

So instead of:

```go
user.Address.City
```

you can write:

```go
user.City
```

This makes the code simpler while still keeping the data structured.

## Embedding vs Named Fields

You could also define the struct like this:

```go
type User struct {
    Name    string
    Address Address
}
```

This is not embedding. It is a named field.

Now you must access the data like this:

```go
user.Address.City
```

The key difference:

* **Embedding** promotes fields to the outer struct.
* **Named fields** require explicit access through the field name.

## Example with Methods

Embedding also promotes methods.

```go
type Address struct {
    City string
}

func (a Address) Full() string {
    return a.City
}

type User struct {
    Name string
    Address
}
```

You can call:

```go
user.Full()
```

Even though the method is defined on `Address`.

## When to Use Embedding

Use embedding when:

* The inner struct is a core part of the outer struct.
* You want simpler access to fields or methods.
* The relationship feels like composition rather than containment.

Use a named field when:

* You want clearer structure and separation.
* The inner struct is not central to the outer type.
* Explicit access improves readability.

## Challenge

Write a program that:

1. Defines a `Profile` struct with a `Username` field.
2. Defines a `User` struct that embeds `Profile`.
3. Creates a `User` and prints the username using the promoted field.

Verify that you can access the field directly from the outer struct.

## Solution

Great job if you attempted this challenge! Here is one possible solution.

```go
package main

import "fmt"

type Profile struct {
    Username string
}

type User struct {
    Profile
}

func main() {
    user := User{
        Profile: Profile{
            Username: "alice123",
        },
    }

    fmt.Println(user.Username)
}
```

## Summary

You have learned how struct embedding works in Go.

* Embedding includes one struct inside another without a field name.
* Embedded fields are promoted to the outer struct.
* Methods are also promoted.
* Embedding simplifies access to nested data.
* Named fields require explicit access.
* Use embedding for composition when it improves clarity.

Struct embedding is a key part of how Go models relationships between types.
