package main

import "fmt"

type ValidationError struct {
	ErrorField string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field: %s", e.ErrorField)
}

func validateName(name string) error {
	if name == "" {
		return ValidationError{ErrorField: "name"}
	}
	return nil
}

type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

func findUser(id int) error {
	return NotFoundError{Resource: "user"}
}

func main() {
	err := findUser(7)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User found")
}
