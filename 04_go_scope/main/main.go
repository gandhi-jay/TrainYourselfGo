package main

import (
	"../variables"
	"fmt"
)

// package level scope
var x int = 42

func main() {
	fmt.Println(x)
	foo()
	max := max(213)
	fmt.Println(max)

	// Getting from different variable.
	fmt.Println(variables.X_from_var)
	fmt.Println(variables.Increment())
	fmt.Println(variables.Increment())
}

func foo() {
	fmt.Println(x)
}

func max(x int) int {
	return 42 + x
}

// don't do this; bad coding practice to shadow variables
