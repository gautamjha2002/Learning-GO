package main

import "fmt"

const (
	first = iota + 6
	second
	third
)

func main() {
	fmt.Println(first, second, third)
}
