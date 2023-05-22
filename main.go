package main

import (
	"belajargo/basic/array"
	"belajargo/basic/constant"
	"belajargo/basic/operator"
	"belajargo/basic/pointer"
	"belajargo/basic/slices"
	"belajargo/basic/variable"
	"fmt"
)

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	//variable
	variable.Variable()
	fmt.Println()

	//Constant
	constant.Constant()

	fmt.Println()
	//Array
	array.Array()

	// Slices
	fmt.Println()
	slices.Slice()

	// Operator
	fmt.Println()
	operator.Operator()

	// Pointer
	fmt.Println()
	pointer.Pointer()
}
