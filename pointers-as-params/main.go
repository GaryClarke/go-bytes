package main

import "fmt"

func addTen(x *int) {
	*x = *x + 10
}

type User struct {
	Name  string
	Score int
}

func increaseScoreCopy(u User) {
	u.Score++
}

func increaseScorePtr(u *User) {
	u.Score++
}

type Book struct {
	Title string
	Pages int
}

func addPages(b *Book, amount int) {
	b.Pages += amount
}

func main() {
	novel := Book{Title: "The Explorer", Pages: 200}
	addPages(&novel, 50)
	fmt.Println("Updated pages:", novel.Pages)
}
