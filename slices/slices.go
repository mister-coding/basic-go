package slices

import "fmt"

func Slice() {
	var slice1 = []int{1, 2, 4, 5, 6}
	slice2 := []int{1, 2, 4, 5, 6}

	type slice3 struct {
		nama string
		umur int
	}

	biodata := []slice3{
		{
			nama: "Agik Setiawan",
			umur: 20,
		},
		{
			nama: "Agik Setiawan 2",
			umur: 22,
		},
	}

	fmt.Println("Slice----------------")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(biodata)
	fmt.Println(slice2[2])
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println("Slice----------------")
}
