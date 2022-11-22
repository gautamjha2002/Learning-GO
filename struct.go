package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u1 user
	u1.ID = 1
	u1.FirstName = "Gautam"
	u1.LastName = "Jha"
	fmt.Println(u1)
	fmt.Println(u1.FirstName)

	u2 := user{ID: 2, FirstName: "Hardik", LastName: "Aggrwal"}
	fmt.Println(u2)

	u3 := user{
		ID:        3,
		FirstName: "Arthur",
		LastName:  "Dent",
	}

	fmt.Println(u3)
}
