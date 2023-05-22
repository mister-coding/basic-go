package pointer

import "fmt"

func Pointer() {
	var a = 8
	var b *int = &a

	fmt.Println("a :", a)
	fmt.Println("a :", &a)
	fmt.Println()
	fmt.Println("b :", *b)
	fmt.Println("b :", b)

	a = 10
	fmt.Println("------------------")
	fmt.Println("a :", a)
	fmt.Println("a :", &a)
	fmt.Println()
	fmt.Println("b :", *b)
	fmt.Println("b :", b)
}
