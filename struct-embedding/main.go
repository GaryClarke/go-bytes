package main

import "fmt"

type Profile struct {
	Username string
}

type User struct {
	Profile
}

func main() {
	user := User{
		Profile: Profile{
			Username: "alice123",
		},
	}

	fmt.Println(user.Username)
}
