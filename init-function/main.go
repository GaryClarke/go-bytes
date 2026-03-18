package main

import "fmt"

var formats = map[string]string{}
var fileTypes = map[string]string{}

func registerFormat(name string, description string) {
	formats[name] = description
}

func registerTypes(name string, description string) {
	fileTypes[name] = description
}

func init() {
	registerTypes("csv", "Comma-separated values")
}

func main() {
	fmt.Println(fileTypes)
}
