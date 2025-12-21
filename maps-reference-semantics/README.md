# Maps and Reference Semantics in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=3vUSnQ6xYnU)**

## Introduction

You have already seen that some types in Go behave differently when passed to functions. For example, integers and structs are copied by value, while pointers allow functions to modify the original.

Maps behave differently again. Even when you pass a map to a function as a value, the underlying data is not copied. Instead, Go passes a reference to the map's internal structure. This means that functions can modify the contents of a map without needing pointers.

In this lesson, you will learn how maps behave when passed around your program, why they act like references, and how to use this behaviour safely and effectively.

## Uses / Use Cases

- Sharing data across functions without copying.
- Updating a map from helper functions.
- Building counters, frequencies, and lookups.
- Passing configuration or dynamic state between parts of your program.

## Example: Maps Behave Like References

Here is a simple program where a function updates a map.

```go
package main

import "fmt"

func addScore(scores map[string]int, name string, amount int) {
    scores[name] = scores[name] + amount
}

func main() {
    playerScores := map[string]int{
        "Alice": 10,
    }

    addScore(playerScores, "Alice", 5)

    fmt.Println("Alice score:", playerScores["Alice"])
}
```

Expected output:

```
Alice score: 15
```

Let's break this down:

- `playerScores` is a map from string to int.
- When you pass `playerScores` to `addScore`, Go does not copy the entire map.
- Instead, both `playerScores` and `scores` refer to the same underlying data.
- Updating the map inside the function updates the original.

This is very different from how structs or basic values behave.

## Example: Assigning Maps Creates Shared References

Assigning one map to another does not create a copy. Both variables point to the same underlying data.

```go
package main

import "fmt"

func main() {
    a := map[string]int{"x": 1}
    b := a

    b["x"] = 999

    fmt.Println("a:", a["x"])
    fmt.Println("b:", b["x"])
}
```

Expected output:

```
a: 999
b: 999
```

Explanation:

- `a` and `b` refer to the same map.
- Updating through one variable updates the shared data.
- This is reference behaviour, even though the assignment looks like a value copy.

## Example: You Can Still Replace the Map Itself

Although map contents are shared, you can reassign the map variable itself inside a function. This does not affect the original variable.

```go
package main

import "fmt"

func reset(m map[string]int) {
    // Reassigns the parameter, but not the caller's variable
    m = map[string]int{}
}

func main() {
    items := map[string]int{"apples": 3}

    reset(items)

    fmt.Println("Items:", items) // Still has original content
}
```

Expected output:

```
Items: map[apples:3]
```

Explanation:

- Reassigning `m` does not change what `items` points to.
- The reference itself is passed by value, but the data behind it is shared.

## Challenge

Create a program that keeps track of word counts in a map. Your program should:

1. Create a map from string to int.
2. Write a function `incrementCount(m map[string]int, word string)` that increases the count for that word.
3. Call the function several times from `main`.
4. Print the final map to show that updates inside the function changed the original.

Run it using:

```
go run main.go
```

## Solution

Great job if you attempted this challenge! Here's one way to solve it:

```go
package main

import "fmt"

func incrementCount(m map[string]int, word string) {
    m[word] = m[word] + 1
}

func main() {
    counts := map[string]int{}

    incrementCount(counts, "apple")
    incrementCount(counts, "apple")
    incrementCount(counts, "banana")

    fmt.Println(counts)
}
```

Expected output:

```
map[apple:2 banana:1]
```

## Summary

You have now learned how maps behave when passed to functions in Go. Here is what you covered:

- Maps use reference semantics. Passing a map does not copy the underlying data.
- Updating a map inside a function updates the original map.
- Assigning one map variable to another creates shared references, not copies.
- Reassigning a map variable inside a function does not affect the caller.
- This makes maps powerful and flexible, but you must be aware of how shared data works.

This behaviour makes maps efficient for passing around dynamic state and building data structures that multiple parts of your program can update. However, remember that because maps are shared, changes made in one part of your program will be visible everywhere the map is used. This is useful when you want shared state, but be careful when you need independent copies.

