package main

import "fmt"

func main() {
	world()
	hello()
	// It will print WorldHello
	// using defer, it'll print correctly
	//
	fmt.Println("After using defer")
	defer world()
	hello()

	// A defer statement pushes a function call onto a list.
	// The list of saved calls is executed after the surrounding function returns.
	// Defer is commonly used to simplify functions that perform various clean-up actions.
}

func world() {
	fmt.Println("World")
}

func hello() {
	fmt.Println("Hello ")
}
