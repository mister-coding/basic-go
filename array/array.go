package array

import "fmt"

func Array() {
	var array1 = [3]int{1, 2, 3}
	array2 := [4]int{1, 2, 3, 4}
	array3 := [5]interface{}{1, 2, 3, "Empat", 5}

	fmt.Println("Array----------------")
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array3[3])
	fmt.Println(len(array2))
	fmt.Println("Array----------------")
}
