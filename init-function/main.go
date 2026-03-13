package main

import "fmt"

var fileTypes = map[string]string{}

func registerType(name string, description string) {
	fileTypes[name] = description
}

func init() {
	registerType("csv", "Comma-separated values")
}

func main() {
	fmt.Println(fileTypes)
}
