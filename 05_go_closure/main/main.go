package main

import "fmt"

var x int = 123
var inc int = 312

func main() {
	fmt.Println(x)
	foo()

	// block
	y := 42
	fmt.Println(y)
	{
		fmt.Println(y) // Y can be accessed in inner scope or you can initialize new one also.
		z := "Lords of the Ring : The Return of the King"
		fmt.Println(z)
	}
	// fmt.Println(z)  - Not available outside of the block

	fmt.Println(increment())
	fmt.Println(increment())
}

func foo() {
	fmt.Println(x)
}

func increment() int {
	inc++
	return inc
}
