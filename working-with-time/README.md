# Working with Time in Go

**[Watch this lesson on YouTube](https://youtu.be/89-KuSpFGXo)**

## Introduction

Many programs need to work with time. You might want to record when something happened, measure how long an operation took, or display the current date and time to a user.

Go provides the built-in `time` package for working with dates, times, and durations. It is powerful, but you can get a lot done with just a few core functions.

In this lesson, you will learn how to get the current time, inspect its parts, and format it into readable strings.

## Uses / Use Cases

* Recording timestamps for logs and events.
* Displaying the current date and time.
* Measuring when something happened.
* Formatting time for output or files.
* Adding simple time awareness to your programs.

## Getting the Current Time with time.Now

The most common starting point is `time.Now()`.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)
}
```

Explanation:

* `time.Now()` returns the current local time.
* The value returned is of type `time.Time`.
* Printing it directly gives a default, readable format.

Your output will look similar to this:

```
2026-01-02 14:30:15.123456 +0000 GMT
```

The exact value will depend on your system time and location.

## Accessing Parts of a Time

A `time.Time` value contains a lot of information. You can extract individual parts using methods.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    fmt.Println("Year:", now.Year())
    fmt.Println("Month:", now.Month())
    fmt.Println("Day:", now.Day())
    fmt.Println("Hour:", now.Hour())
    fmt.Println("Minute:", now.Minute())
    fmt.Println("Second:", now.Second())
}
```

Explanation:

* These methods return integers or values representing parts of the date and time.
* This is useful when you need specific components rather than the full timestamp.

## Formatting Time Values

Often, you will want to format a time value in a specific way. Go uses a unique formatting system based on a reference time.

The reference time is:

```
Mon Jan 2 15:04:05 MST 2006
```

Each part of this reference date represents a component of time:

| Reference Component | Represents | Example Output |
|---------------------|------------|----------------|
| `2006` | Year (4 digits) | `2026` |
| `06` | Year (2 digits) | `26` |
| `01` | Month (numeric) | `01` |
| `Jan` | Month (abbreviated) | `Jan` |
| `January` | Month (full name) | `January` |
| `02` | Day | `15` |
| `15` | Hour (24-hour format) | `14` |
| `03` | Hour (12-hour format) | `02` |
| `04` | Minute | `30` |
| `05` | Second | `45` |
| `PM` | AM/PM | `PM` |

When you write a layout string, you use these exact components from the reference date. Go then replaces each part with the corresponding value from your actual time.

You format times by writing a layout string that uses this exact reference.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    formatted := now.Format("2006-01-02 15:04:05")
    fmt.Println(formatted)
}
```

Expected output:

```
2026-01-02 14:30:15
```

Explanation:

* The numbers in the layout are not placeholders.
* They represent the reference date and time.
* Go matches each part of the layout to the corresponding value in `now`.

This takes a little getting used to, but it becomes very predictable once you know the pattern.

## Challenge

Write a program that:

1. Gets the current time using `time.Now()`.
2. Prints the full time value.
3. Prints the date in `YYYY-MM-DD` format.
4. Prints the time in `HH:MM:SS` format.

Run it using:

```
go run main.go
```

## Common Formatting Examples

Here are a few common formats you may find useful:

```go
now := time.Now()

fmt.Println(now.Format("2006-01-02"))       // Date only
fmt.Println(now.Format("15:04:05"))         // Time only
fmt.Println(now.Format("02 Jan 2006"))      // Human readable date
fmt.Println(now.Format("2006/01/02 15:04")) // Custom format
```

Each call returns a string. The original `time.Time` value is unchanged.

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    fmt.Println("Full:", now)
    fmt.Println("Date:", now.Format("2006-01-02"))
    fmt.Println("Time:", now.Format("15:04:05"))
}
```

Expected output:

```
Full: 2026-01-02 14:30:15.123456 +0000 GMT
Date: 2026-01-02
Time: 14:30:15
```

The exact values will match your current system time when you run the program.

## Aside: Useful Time Constants

The `time` package provides several constants that represent common durations. These are useful when you need to work with time intervals or durations.

```go
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour
```

Each constant represents the duration it names. For example, `time.Second` is one second, `time.Minute` is one minute, and `time.Hour` is one hour.

These constants are particularly useful when working with durations, time arithmetic, or functions that accept time intervals. You will see them used in future lessons when we cover time-based operations like adding time to a timestamp or measuring elapsed time.

## Summary

You have now learned the basics of working with time in Go. Here is what you covered:

* The `time` package is used for working with dates and times.
* `time.Now()` returns the current local time.
* A `time.Time` value contains many useful parts.
* You can extract individual components like year, month, and hour.
* Time formatting uses a fixed reference date.
* Formatting returns strings and does not modify the original value.

These basics will support future lessons on durations, time comparisons, and scheduling.


