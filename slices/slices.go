package slices

import "fmt"

func Slice() {
	var slice1 = []int{1, 2, 4, 5, 6}
	slice2 := []int{1, 2, 4, 5, 6}

	fmt.Println("Slice----------------")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice2[2])
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println("Slice----------------")
}
