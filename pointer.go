package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "Gautam"
	fmt.Println(*firstName)

	lastName := "Jha"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Sharma"
	fmt.Println(ptr, *ptr)
}
