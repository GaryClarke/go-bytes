# Map Lookup with the Comma Ok Idiom in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=yG7hIhhy8XM)**

## Introduction

Maps are one of the most useful data structures in Go, but they come with an important detail. When you look up a key in a map, Go always returns a value, even if the key does not exist.

This can be confusing at first, because you cannot tell whether the value came from the map or is just the zero value for that type. To solve this, Go provides a common pattern called the *comma ok idiom*.

In this lesson, you will learn how map lookups work, why the comma ok idiom exists, and how to use it to safely check whether a key is present in a map.

## Uses / Use Cases

* Checking whether a key exists before using its value.
* Avoiding bugs caused by zero values.
* Safely reading from maps populated dynamically.
* Writing clearer and more defensive code.

## Map Lookups Without Checking

Let's start with a simple map lookup.

```go
package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 10,
        "Bob":   7,
    }

    score := scores["Sarah"]
    fmt.Println("Score:", score)
}
```

Expected output:

```
Score: 0
```

Explanation:

* The key `"Sarah"` does not exist in the map.
* Go returns the zero value for `int`, which is `0`.
* There is no error or warning.

At this point, you cannot tell whether Sarah actually has a score of 0 or whether the key was missing.

## Using the Comma Ok Idiom

To distinguish between a missing key and a real value, use the comma ok idiom.

```go
package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 10,
        "Bob":   7,
    }

    score, ok := scores["Sarah"]
    if !ok {
        fmt.Println("Sarah not found")
        return
    }

    fmt.Println("Score:", score)
}
```

Explanation:

* The lookup returns two values.
* `score` is the value from the map, or the zero value.
* `ok` is a boolean that tells you whether the key exists.
* If `ok` is false, the key was not present.

This is the safest and most common way to read from maps.

## Example: Key Exists vs Key Missing

Here is a comparison that shows both cases.

```go
package main

import "fmt"

func main() {
    users := map[string]string{
        "alice": "admin",
        "bob":   "editor",
    }

    role, ok := users["alice"]
    fmt.Println("alice exists:", ok, "role:", role)

    role, ok = users["sarah"]
    fmt.Println("sarah exists:", ok, "role:", role)
}
```

Expected output:

```
alice exists: true role: admin
sarah exists: false role:
```

Explanation:

* When the key exists, `ok` is true and the value is meaningful.
* When the key does not exist, `ok` is false and the value is the zero value for the type.

The comma ok idiom is important because zero values are valid values in maps. You cannot rely on the value alone to detect missing keys. Explicit checks make your intent clear to anyone reading the code. This pattern is used heavily in real Go programs.

## Challenge

Write a program that:

1. Creates a map of usernames to email addresses.
2. Looks up a username provided in the code.
3. Uses the comma ok idiom to check if the user exists.
4. Prints the email if found, or a helpful message if not.

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
    emails := map[string]string{
        "alice": "alice@example.com",
        "bob":   "bob@example.com",
    }

    username := "sarah"

    email, ok := emails[username]
    if !ok {
        fmt.Println("User not found:", username)
        return
    }

    fmt.Println("Email:", email)
}
```

If you didn't get it exactly right the first time, that's completely normal. You're building new skills!
The important thing is understanding how the comma ok idiom lets you safely check for key existence in maps.

## Summary

You have now learned how to safely access maps using the comma ok idiom. Here is what you covered:

* Looking up a missing key returns the zero value.
* This can be misleading without an existence check.
* The comma ok idiom lets you check whether a key exists.
* `value, ok := map[key]` is the standard Go pattern.
* Using this idiom makes your code safer and clearer.

This is a small pattern, but it appears everywhere in Go code, so becoming comfortable with it early will pay off quickly.

