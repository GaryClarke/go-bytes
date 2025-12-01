package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Height: 10, Width: 5}
	fmt.Println(rect.Area())
}
