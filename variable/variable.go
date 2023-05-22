package variable

import "fmt"

func Variable() {
	var x int = 1
	var y = 10
	var a, b, c, d = 1, 2, 3, "empat"
	aaa := 123
	abc, def := 123, 345

	var ccc bool = false
	ddd := true

	fmt.Println("Variable-------------------------")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a, b, c, d)
	fmt.Println(abc, def)
	fmt.Println(aaa)
	fmt.Println(ccc)
	fmt.Println(ddd)
	fmt.Println("Variable-------------------------")
}
