package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	// slice is always built on top of array
	slice := arr[:] // begining to end of array to slice

	arr[1] = 42   // change displayed in both slice and array
	slice[2] = 27 // change displayed in both slice and array
	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2)

	slice3 := slice[1:]  // go from index 1 to last
	slice4 := slice[:2]  // from start to element 2 not including 2
	slice5 := slice[1:2] // go from 1 to element 2 but not including 2

	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
