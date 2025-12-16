# Formatting Your Code with gofmt

**[Watch this lesson on YouTube](https://youtube.com/watch?v=O3GcXK38Ut4)**

## Introduction

Readable code is easier to understand, easier to maintain, and easier to review. Go takes this seriously by providing a built-in tool called `gofmt` that formats your code automatically. Instead of worrying about spacing, indentation, or style choices, you can let `gofmt` handle everything for you.

The good news is that `gofmt` comes with Go, so you don't need to install anything extra. It's ready to use as soon as you have Go installed on your system.

The result is that all Go code looks consistent, no matter who writes it. This reduces friction when working in teams and makes it easier to spot real problems in your code.

In this lesson, you will learn what `gofmt` does, how to run it, and how it helps keep your code clean and idiomatic.

## Uses / Use Cases

* Automatically fixing indentation and spacing.
* Keeping code style consistent across files and projects.
* Making diffs cleaner and easier to read.
* Avoiding style debates by relying on Go's standard formatting rules.
* Producing code that other Go developers can read comfortably.

## Example: Formatting a File

Here is a small program written with messy spacing:

```go
package main

import "fmt"

func main(){
fmt.Println("Hello")
fmt.Println(    "World"    )
}
```

If you want to see what `gofmt` would do without changing your file, you can run it without the `-w` flag:

```
gofmt main.go
```

This prints the formatted version to your terminal:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
    fmt.Println("World")
}
```

If you save this in a file named `main.go` and run:

```
gofmt -w main.go
```

**Note:** Make sure you've saved the file first. `gofmt` reads from the file on disk, not from your editor's unsaved buffer.

The file will be rewritten with proper formatting:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
    fmt.Println("World")
}
```

Let's break this down:

* `gofmt` reads your Go file.
* It rewrites the file using standard formatting rules.
* The `-w` flag means "write the result back to the file".

Without `-w`, `gofmt` prints the formatted version to the terminal instead of replacing the file.

## gofmt with Multiple Files

You can format every Go file in a folder by running:

```
gofmt -w .
```

This is useful when you add new files or refactor existing ones.

## gofmt Through go fmt

Go also provides a wrapper command called:

```
go fmt
```

Running `go fmt` inside a module formats all Go files in the current package. Under the hood, `go fmt` runs `gofmt -l -w` on your package, which formats files and lists any files that were changed.

Key differences:

* `gofmt` gives you more control and can run on individual files or paths.
* `go fmt` is more convenient for formatting an entire package or module.

Either approach will give you clean, idiomatic Go code.

## Editor Integration

Many editors, including GoLand, VS Code, and Vim, can automatically run `gofmt` when you save a file. This means you might not need to run `gofmt` manually in your day-to-day work.

However, understanding how `gofmt` works from the command line is still valuable. It helps when:
* Working in environments without editor integration
* Formatting files in scripts or CI/CD pipelines
* Understanding what your editor is doing behind the scenes

## Expected Output

For the initial example, running `gofmt -w main.go` updates the code in place. After running it, the program prints:

```
Hello
World
```

There is no change in behaviour, only in appearance.

## Challenge

Create a Go file named `messy.go` with intentionally inconsistent formatting. Include at least a function and a print statement.

Then run:

```
gofmt -w messy.go
```

Open the file again and observe how the tool has cleaned up indentation, spacing, and layout.

**Extra credit:** If you want to go further, run:

```
go fmt
```

inside your module to format every file.

## Solution

Here is one possible messy file:

```go
package main

import "fmt"

func  greet( name string ){
fmt.Println("Hello,",name)
}
```

After running `gofmt -w messy.go`, it becomes:

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello,", name)
}
```

Great job if you attempted this challenge! If you also tried `go fmt`, you've seen how easy it is to format entire packages at once.

## Summary

You have now learned how `gofmt` keeps your code clean and consistent. Here is what you covered:

* `gofmt` formats Go code automatically.
* The `-w` flag writes the formatted result back to the file.
* You can format a single file, a folder, or an entire module.
* `go fmt` is a convenient wrapper for formatting the current package.
* Consistent formatting makes your code easier to read and maintain.

By using `gofmt` regularly, you ensure your code stays clear and idiomatic as your projects grow.

