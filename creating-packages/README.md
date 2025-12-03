# Creating Packages in Go

**[Watch this lesson on YouTube](https://youtube.com/watch?v=N3zZDu7pLP4)**

## Introduction

As your Go programs grow, you'll want to organise your code into smaller, focused pieces.

This is where **packages** come in. Packages let you group related functions and logic in their own folders, making your project easier to navigate, reuse, and maintain.

In this lesson, you'll create your own package, learn how to import it, and understand how Go decides what can be accessed from outside the package.

You'll also get familiar with how `go mod` works and how folder structure affects your imports.

## Uses / Use Cases

* Organising code into reusable pieces across your project

* Making your `main.go` file cleaner by moving logic into packages

* Sharing logic between different parts of a program

* Building structure for real-world applications and APIs

## Example

Let's create a simple folder structure like this:

```
myapp/
├── go.mod
├── main.go
└── greetings/
    └── greet.go
```

### Step 1: Set up your module

In the root folder (`myapp`), run:

```
go mod init myapp
```

This creates a `go.mod` file. The name `myapp` becomes the base path for importing your own packages.

---

### Step 2: Create a new package

Inside the `greetings` folder, create a file named `greet.go`:

```go
package greetings

import "fmt"

func SayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

* `package greetings` defines the package name.

* `SayHello` begins with a capital letter, so it's **exported** — meaning it can be used outside this package.

* The function simply prints a friendly message.

---

### Step 3: Import and use your package

In `main.go`:

```go
package main

import (
    "myapp/greetings"
)

func main() {
    greetings.SayHello("Alice")
}
```

Then run:

```
go run .
```

If everything's set up properly, you'll see:

```
Hello, Alice!
```

## Expected Output

```
Hello, Alice!
```

## Challenge

Create a new package called `messages`.

Inside it, define a function named `WelcomeMessage` that returns a custom greeting string.

Follow these steps:

1. Create a folder called `messages`.

2. Inside it, create a file called `text.go`.

3. Define an exported function named `WelcomeMessage` that takes a `name` and returns:

   ```
   "Welcome, <name>!"
   ```

4. Import the `messages` package into `main.go`.

5. Call `messages.WelcomeMessage()` and print the returned string.

Run your program with:

```
go run .
```

## Solution

Great job if you attempted this challenge! Here's the solution:

```go
// messages/text.go
package messages

import "fmt"

func WelcomeMessage(name string) string {
    return fmt.Sprintf("Welcome, %s!", name)
}
```

```go
// main.go
package main

import (
    "fmt"
    "myapp/messages"
)

func main() {
    msg := messages.WelcomeMessage("Alice")
    fmt.Println(msg)
}
```

Expected output:

```
Welcome, Alice!
```

If your code didn't work right away, double-check:

* Your package name matches the folder name (`messages`)

* You're using a capital letter for the exported function

* Your `go.mod` file is in the root of the project and has the right module path

## Summary

In this lesson you learned how to split your project into packages.

Here is what you covered:

* `go mod init` sets up your module name for imports

* Every package lives in its own folder and starts with `package <name>`

* Only exported names (capitalised) can be accessed from other packages

* You can import your own packages using the module name and folder structure

* This pattern is used in all real Go apps to keep code clean and maintainable

Now you're ready to start organising your own Go projects like a pro.

