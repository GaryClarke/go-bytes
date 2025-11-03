# Working with Maps in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=SvQ5ee6SEwU)**

## Introduction

A **map** in Go is a built-in data type that stores key-value pairs.

Each entry in a map has a key and a value. You use the key to retrieve the corresponding value, making maps useful for organizing and quickly looking up data.

In this lesson, you'll create maps, read values, update them, and learn how to check if a key exists.

## Uses / Use Cases

* Store related data using keys
* Model real-world data like user profiles or settings
* Look up values quickly by key
* Create lightweight, dynamic structures without defining a full `struct`

## Example

Here's a simple program that defines a map of colors:

```go
package main

import "fmt"

func main() {
    colors := map[string]string{
        "r": "red",
        "g": "green",
        "b": "blue",
    }

    fmt.Println(colors["g"]) // green
}
```

Let's break it down:

* `map[string]string` means the keys and values are both strings
* The keys are `"r"`, `"g"`, and `"b"`
* Accessing `colors["g"]` gives us `"green"`

### Adding and Updating Values

```go
colors["y"] = "yellow"  // add
colors["g"] = "gray"    // update
```

Maps are flexible. You can add or update values at any time using bracket notation.

### Deleting a Key

```go
delete(colors, "r")
```

This removes the key `"r"` and its value from the map.

### Checking If a Key Exists

When you read from a map, Go lets you check whether the key was present:

```go
value, exists := colors["x"]
fmt.Println("Value:", value)
fmt.Println("Exists:", exists)
```

The second value returned (`exists`) will be `true` if the key exists in the map, or `false` if it doesn't.

This pattern (`value, ok := map[key]`) is common in Go. We'll cover how to use the `exists` value with conditionals in a later lesson.

## Expected Output

```
green
Value: 
Exists: false
```

When the key doesn't exist, `value` will be an empty string (the zero value for strings), and `exists` will be `false`.

## Challenge

Write a program that:

1. Creates a map of country codes to country names (e.g., "uk": "United Kingdom")
2. Adds one new country
3. Updates one existing country name
4. Deletes one country
5. Prints out the final map using `fmt.Println()`

Try using `map[string]string` and call the variable `countries`.

## Solution

Great job if you attempted this challenge! Here's one way to do it:

```go
package main

import "fmt"

func main() {
    countries := map[string]string{
        "uk": "United Kingdom",
        "fr": "France",
        "de": "Germany",
    }

    countries["us"] = "United States"     // add
    countries["fr"] = "French Republic"   // update
    delete(countries, "de")               // delete

    fmt.Println(countries)
}
```

Your output will look something like:

```
map[fr:French Republic uk:United Kingdom us:United States]
```

Go prints maps using the `map[key:value]` format.

If your solution is different, that's perfectly fine. The goal is to practice creating, updating, and deleting map entries.

## Summary

In this lesson, you learned:

* A map stores key-value pairs in Go
* Use `map[keyType]valueType` to declare a map
* Read values with `map[key]`
* Add or update values with `map[key] = value`
* Use `delete(map, key)` to remove a key
* Check if a key exists with the `value, ok := map[key]` pattern

Maps are a flexible way to structure data without needing a full struct.

They're also very fast for lookups, which makes them perfect for config, settings, or small datasets.

