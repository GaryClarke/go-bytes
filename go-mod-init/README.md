# Initializing a Go Module with `go mod init`

**[Watch this lesson on YouTube](https://youtube.com/watch?v=VoFCSjlyYFI)**

## Introduction
Go modules are how Go organizes and manages your project's code. A module defines the root of your project and tells Go where to find and store dependencies. Each module has a file named `go.mod` that keeps track of this information.

In this lesson, you'll learn what a Go module is and how to create one using the `go mod init` command. You'll see exactly what happens when you run the command and what each line in the generated file means.

## Uses / Use Cases
- Setting up a new Go project
- Defining a clear project root so Go can manage dependencies
- Ensuring consistent builds across environments
- Preparing your code for sharing or importing in other projects

## Example
Create a new folder for your project and move into it:

```bash
mkdir go-mod-init
cd go-mod-init
```

Now initialize a module by running:

```bash
go mod init go-mod-init
```

This creates a new file named `go.mod` inside your project folder.  
Open it, and you'll see something like this:

```go
// This defines the name of our module — it's also used as the import path if someone wants to use your code.
module go-mod-init

// This tells Go which version of the language this module was created with.
// It doesn't restrict you — you can still build this with any Go version >= 1.25.3.
go 1.25.3
```

> **Note:** The Go version shown (1.25.3) will match your system's Go version. Your version may differ.

Let's look at what's happening here:

* **`module go-mod-init`**
  This line sets the name of your module. Go uses it as the base path when importing your code into other projects. If your code lives on GitHub, you could instead use a repository path like:

  ```go
  module github.com/yourusername/go-mod-init
  ```
  
  In real projects, module names often match repository paths so other developers can import your code. For Go Bytes lessons, we use simple names like `go-mod-init` to keep things focused on the concepts.

* **`go 1.25.3`**
  This tells Go which version of the language was active when the module was created. It doesn't lock the version, but it helps Go understand which language features are safe to use.

At this point, your project is now a Go module.
Any Go files you create inside this folder will belong to the same module, and Go will automatically track any new dependencies you import.

## Summary

You learned how to initialize a Go module with `go mod init`.
A module tells Go where your project starts and helps it manage dependencies and versions.

* Run `go mod init <module-name>` to create the module file.
* The `module` line defines your project's name or path.
* The `go` line notes the Go version used when creating the module.

With your first module set up, your project is ready for Go source files and dependencies.
