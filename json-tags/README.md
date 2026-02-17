# Go Structs and JSON Tags

**[Watch this lesson on YouTube](https://youtube.com/watch?v=yL_6h0vz7CY)**

## Introduction

Structs define the shape of your data.

When working with HTTP APIs, that data often needs to be converted to JSON.

In this lesson, you will learn how struct field names affect JSON output, and how JSON tags allow you to control the encoded result.

## Uses / Use Cases

* Returning structured data from an API.
* Encoding responses as JSON.
* Matching external API formats.
* Controlling how your data appears in JSON.

## Encoding a Struct to JSON

Go provides JSON support in the `encoding/json` package.

Here is a simple example.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string
    Email string
}

func main() {
    user := User{
        Name:  "Alice",
        Email: "alice@example.com",
    }

    data, _ := json.Marshal(user)

    fmt.Println(string(data))
}
```

Output:

```
{"Name":"Alice","Email":"alice@example.com"}
```

We ignore the error from `json.Marshal` here for simplicity; in real code you would check it.

Notice something important.

The JSON keys match the struct field names exactly.

## Exported Fields Matter

Only exported fields are included in JSON encoding.

This means the field name must start with a capital letter.

If we changed `Name` to `name`, it would not appear in the JSON output.

JSON encoding only works with exported fields.

## Using JSON Tags

Struct tags let you control how fields are encoded.

Here is the same struct with JSON tags.

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

Now the output becomes:

```
{"name":"Alice","email":"alice@example.com"}
```

The tag overrides the default behaviour and sets the JSON key explicitly.

This is very common when building APIs, because JSON typically uses lowercase or snake_case field names.

## Omitting Empty Fields

You can also tell Go to omit empty values using `omitempty`.

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

If `Email` is empty, it will not appear in the JSON output.

This is useful when certain fields are optional.

## Challenge

Write a program that:

1. Defines a struct called `Product`.
2. Includes fields for `ID`, `Name`, and `Price`.
3. Adds JSON tags so the output keys are `id`, `name`, and `price`.
4. Encodes a product to JSON and prints it.

Verify that the JSON keys match your tags exactly.

## Solution

Great job if you attempted this challenge! Here is one possible solution:

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    product := Product{
        ID:    1,
        Name:  "Notebook",
        Price: 9.99,
    }

    data, _ := json.Marshal(product)

    fmt.Println(string(data))
}
```

The output should look like:

```
{"id":1,"name":"Notebook","price":9.99}
```

## Summary

You have learned how structs are encoded to JSON in Go.

* `encoding/json` converts structs to JSON.
* JSON keys match exported struct field names by default.
* Only exported fields are included.
* JSON tags allow you to customise key names.
* `omitempty` removes empty fields from the output.

Understanding this behaviour is essential when building APIs and working with external services.
