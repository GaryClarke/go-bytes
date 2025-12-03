package main

import (
	"fmt"
	"myapp/greetings"
	"myapp/messages"
)

func main() {
	greetings.SayHello("Gary")

	msg := messages.WelcomeMessage("Gary")
	fmt.Println(msg)
}
