package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type PrefixedLogger struct {
	Prefix string
}

func (l PrefixedLogger) Log(message string) {
	fmt.Println(l.Prefix, message)
}

func runLog(l Logger, m string) {
	l.Log(m)
}

func main() {
	plain := ConsoleLogger{}
	info := PrefixedLogger{Prefix: "[INFO]"}

	runLog(plain, "Starting application")
	runLog(info, "User logged in")
}
